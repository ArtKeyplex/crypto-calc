package usecase_test

import (
	"context"
	"exchange-rate-calculator/internal/adapter/api"
	"exchange-rate-calculator/internal/entity"
	"exchange-rate-calculator/internal/usecase"
	"exchange-rate-calculator/internal/usecase/mocks"
	"testing"

	customErrors "exchange-rate-calculator/internal/errors"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestConvertCurrency_Success(t *testing.T) {
	ctx := context.TODO()
	repo := mocks.NewCurrencyRepository(t)
	client := mocks.NewFastForexClient(t)
	uc := usecase.NewCurrencyUsecase(repo, client)

	fromCurrency := &entity.Currency{
		Code:      "USD",
		Available: true,
		Type:      "fiat",
		Rate:      decimal.NewFromFloat(1.0),
	}
	toCurrency := &entity.Currency{
		Code:      "ETH",
		Available: true,
		Type:      "crypto",
		Rate:      decimal.NewFromFloat(0.00026),
	}

	repo.On("GetCurrencyByCode", ctx, "USD").Return(fromCurrency, nil)
	repo.On("GetCurrencyByCode", ctx, "ETH").Return(toCurrency, nil)

	result, err := uc.ConvertCurrency(ctx, "USD", "ETH", 100)
	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.InDelta(t, 0.026, result.Result, 0.001)
}

func TestConvertCurrency_CurrencyNotAvailable(t *testing.T) {
	ctx := context.TODO()
	repo := mocks.NewCurrencyRepository(t)
	client := mocks.NewFastForexClient(t)
	uc := usecase.NewCurrencyUsecase(repo, client)

	fromCurrency := &entity.Currency{
		Code:      "USD",
		Available: true,
		Type:      "fiat",
		Rate:      decimal.NewFromFloat(1.0),
	}
	toCurrency := &entity.Currency{
		Code:      "ETH",
		Available: false, // Not available
		Type:      "crypto",
		Rate:      decimal.NewFromFloat(0.00026),
	}

	repo.On("GetCurrencyByCode", ctx, "USD").Return(fromCurrency, nil)
	repo.On("GetCurrencyByCode", ctx, "ETH").Return(toCurrency, nil)

	result, err := uc.ConvertCurrency(ctx, "USD", "ETH", 100)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, customErrors.ErrCurrencyNotAvailable, err)
}

func TestConvertCurrency_SameCurrencyType(t *testing.T) {
	ctx := context.TODO()
	repo := mocks.NewCurrencyRepository(t)
	client := mocks.NewFastForexClient(t)
	uc := usecase.NewCurrencyUsecase(repo, client)

	fromCurrency := &entity.Currency{
		Code:      "USD",
		Available: true,
		Type:      "fiat",
		Rate:      decimal.NewFromFloat(1.0),
	}
	toCurrency := &entity.Currency{
		Code:      "EUR",
		Available: true,
		Type:      "fiat",
		Rate:      decimal.NewFromFloat(0.85),
	}

	repo.On("GetCurrencyByCode", ctx, "USD").Return(fromCurrency, nil)
	repo.On("GetCurrencyByCode", ctx, "EUR").Return(toCurrency, nil)

	result, err := uc.ConvertCurrency(ctx, "USD", "EUR", 100)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, customErrors.ErrCurrencyCanNotBeCalculated, err)
}

func TestUpdateRates(t *testing.T) {
	ctx := context.TODO()
	repo := mocks.NewCurrencyRepository(t)
	client := mocks.NewFastForexClient(t)
	uc := usecase.NewCurrencyUsecase(repo, client)

	currencies := []*entity.Currency{
		{
			Code:      "USD",
			Available: true,
		},
		{
			Code:      "EUR",
			Available: true,
		},
	}

	ratesResponse := &api.FetchOneResponse{
		Result: map[string]decimal.Decimal{
			"USD": decimal.NewFromFloat(1.0),
			"EUR": decimal.NewFromFloat(0.85),
		},
	}

	repo.On("GetAllCurrencies", ctx).Return(currencies, nil)
	client.On("GetRates", ctx, "USD", "USD").Return(ratesResponse, nil)
	client.On("GetRates", ctx, "USD", "EUR").Return(ratesResponse, nil)
	repo.On("UpdateRates", ctx, mock.MatchedBy(func(r []*entity.Currency) bool { return len(r) == 2 })).Return(nil)

	err := uc.UpdateRates(ctx)
	require.NoError(t, err)
	repo.AssertCalled(t, "UpdateRates", ctx, mock.Anything)
}
