package postgresql

import (
	"context"
	"errors"

	"exchange-rate-calculator/internal/entity"
	customError "exchange-rate-calculator/internal/errors"

	"gorm.io/gorm"
)

type CurrencyRepositoryPostgresql struct {
	db *gorm.DB
}

func NewCurrencyRepository(db *gorm.DB) *CurrencyRepositoryPostgresql {
	return &CurrencyRepositoryPostgresql{db: db}
}

func (r *CurrencyRepositoryPostgresql) GetCurrencyByCode(ctx context.Context, code string) (*entity.Currency, error) {
	var currency entity.Currency
	result := r.db.WithContext(ctx).Where("code = ?", code).First(&currency)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, customError.ErrCurrencyNotFound
		}

		return nil, result.Error
	}

	return &currency, nil
}

func (r *CurrencyRepositoryPostgresql) GetAllCurrencies(ctx context.Context) ([]*entity.Currency, error) {
	var currencies []*entity.Currency
	result := r.db.WithContext(ctx).Find(&currencies)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, customError.ErrCurrencyNotFound
		}

		return nil, result.Error
	}

	return currencies, nil
}

func (r *CurrencyRepositoryPostgresql) UpdateRates(ctx context.Context, currencies []*entity.Currency) error {
	tx := r.db.Begin().WithContext(ctx)
	if tx.Error != nil {
		return tx.Error
	}

	for _, currency := range currencies {
		if err := tx.Model(&entity.Currency{}).Where("code = ?", currency.Code).Update("rate", currency.Rate).Error; err != nil {
			tx.Rollback()

			return err
		}
	}

	return tx.Commit().Error
}
