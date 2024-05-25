package usecase

import (
	"context"
	"fmt"

	"exchange-rate-calculator/internal/adapter/api"
	"exchange-rate-calculator/internal/entity"
	customErrors "exchange-rate-calculator/internal/errors"

	"github.com/rs/zerolog"
	"github.com/shopspring/decimal"
)

//go:generate mockery --name=CurrencyRepository
type CurrencyRepository interface {
	GetCurrencyByCode(ctx context.Context, code string) (*entity.Currency, error)
	GetAllCurrencies(ctx context.Context) ([]*entity.Currency, error)
	UpdateRates(ctx context.Context, currencies []*entity.Currency) error
}

type CurrencyUsecase interface {
	ConvertCurrency(ctx context.Context, from, to string, amount float64) (*entity.ConversionResult, error)
	UpdateRates(ctx context.Context) error
}

//go:generate mockery --name=FastForexClient
type FastForexClient interface {
	GetRates(ctx context.Context, baseCurrency string, symbol string) (*api.FetchOneResponse, error)
}

type CurrencyUsecaseImpl struct {
	repo            CurrencyRepository
	fastForexClient FastForexClient
}

func NewCurrencyUsecase(repo CurrencyRepository, fastForexClient FastForexClient) *CurrencyUsecaseImpl {
	return &CurrencyUsecaseImpl{repo: repo, fastForexClient: fastForexClient}
}

func (u *CurrencyUsecaseImpl) ConvertCurrency(ctx context.Context,
	from, to string, amount float64,
) (*entity.ConversionResult, error) {
	fromCurrency, err := u.repo.GetCurrencyByCode(ctx, from)
	if err != nil {
		return nil, fmt.Errorf("failed to get currency: %w", err)
	}

	toCurrency, err := u.repo.GetCurrencyByCode(ctx, to)
	if err != nil {
		return nil, fmt.Errorf("failed to get currency: %w", err)
	}

	if !fromCurrency.Available || !toCurrency.Available {
		return nil, customErrors.ErrCurrencyNotAvailable
	}

	if fromCurrency.Type == toCurrency.Type {
		return nil, customErrors.ErrCurrencyCanNotBeCalculated
	}

	amountDec := decimal.NewFromFloat(amount)

	fromRateDec := fromCurrency.Rate
	toRateDec := toCurrency.Rate

	resultDec := amountDec.Mul(toRateDec).Div(fromRateDec)

	// в зависимости от требований можно округлить результирующее значение до нужного количества нулей
	result, _ := resultDec.Float64()

	return &entity.ConversionResult{
		From:   from,
		To:     to,
		Amount: amount,
		Result: result,
	}, nil
}

func (u *CurrencyUsecaseImpl) UpdateRates(ctx context.Context) error {
	currencies, err := u.repo.GetAllCurrencies(ctx)
	if err != nil {
		return fmt.Errorf("failed to get all currencies: %w", err)
	}

	const baseCurrency = "USD"

	for _, currency := range currencies {
		ratesResponse, err := u.fastForexClient.GetRates(ctx, baseCurrency, currency.Code)
		if err != nil {
			currency.Available = false
			zerolog.Ctx(ctx).Error().Err(err).Msgf("failed to update rate for currency %s", currency.Code)

			continue
		}

		if rate, ok := ratesResponse.Result[currency.Code]; ok {
			currency.Rate = rate
			currency.Available = true
		} else {
			currency.Available = false
			zerolog.Ctx(ctx).Error().Msgf("rate for currency %s not found in response", currency.Code)
		}
	}

	if err := u.repo.UpdateRates(ctx, currencies); err != nil {
		return fmt.Errorf("failed to update rates in repository: %w", err)
	}

	return nil
}
