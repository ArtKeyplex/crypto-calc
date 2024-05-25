package postgresql_test

import (
	"context"
	"exchange-rate-calculator/internal/entity"
	"exchange-rate-calculator/internal/repository/postgresql"
	"testing"

	customError "exchange-rate-calculator/internal/errors"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	t.Helper()
	// для тестов бд, по-хорошему, надо поднимать такую же бд как и в основном приложении
	dsn := "file::memory:?cache=shared"
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect to in-memory database: %v", err)
	}

	err = db.AutoMigrate(&entity.Currency{})
	if err != nil {
		t.Fatalf("failed to migrate database: %v", err)
	}

	return db
}

func TestCurrencyRepositoryPostgresql_GetCurrencyByCode(t *testing.T) {
	db := setupTestDB(t)
	repo := postgresql.NewCurrencyRepository(db)
	ctx := context.TODO()

	currency := entity.Currency{Code: "USD", Available: true, Type: "fiat", Rate: decimal.NewFromFloat(1.0)}
	db.Create(&currency)

	t.Run("Currency Found", func(t *testing.T) {
		result, err := repo.GetCurrencyByCode(ctx, "USD")
		require.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, "USD", result.Code)
	})

	t.Run("Currency Not Found", func(t *testing.T) {
		result, err := repo.GetCurrencyByCode(ctx, "EUR")
		require.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, customError.ErrCurrencyNotFound, err)
	})
}

func TestCurrencyRepositoryPostgresql_GetAllCurrencies(t *testing.T) {
	db := setupTestDB(t)
	repo := postgresql.NewCurrencyRepository(db)
	ctx := context.TODO()

	currencies := []entity.Currency{
		{Code: "USD", Available: true, Type: "fiat", Rate: decimal.NewFromFloat(1.0)},
		{Code: "EUR", Available: true, Type: "fiat", Rate: decimal.NewFromFloat(0.85)},
	}
	for _, c := range currencies {
		db.Create(&c)
	}

	t.Run("Get All Currencies", func(t *testing.T) {
		result, err := repo.GetAllCurrencies(ctx)
		require.NoError(t, err)
		assert.Len(t, result, 2)
	})
}

func TestCurrencyRepositoryPostgresql_UpdateRates(t *testing.T) {
	db := setupTestDB(t)
	repo := postgresql.NewCurrencyRepository(db)
	ctx := context.TODO()

	currencies := []entity.Currency{
		{Code: "USD", Available: true, Type: "fiat", Rate: decimal.NewFromFloat(1.0)},
		{Code: "EUR", Available: true, Type: "fiat", Rate: decimal.NewFromFloat(0.85)},
	}
	for _, c := range currencies {
		db.Create(&c)
	}

	t.Run("Update Rates", func(t *testing.T) {
		var currenciesToUpdate []*entity.Currency
		db.Find(&currenciesToUpdate)

		for _, currency := range currenciesToUpdate {
			switch currency.Code {
			case "USD":
				currency.Rate = decimal.NewFromFloat(1.1)
			case "EUR":
				currency.Rate = decimal.NewFromFloat(0.9)
			}
		}

		err := repo.UpdateRates(ctx, currenciesToUpdate)
		require.NoError(t, err)

		var updatedCurrencies []entity.Currency
		db.Find(&updatedCurrencies)

		assert.Len(t, updatedCurrencies, 2)

		expectedRates := map[string]decimal.Decimal{
			"USD": decimal.NewFromFloat(1.1),
			"EUR": decimal.NewFromFloat(0.9),
		}

		for _, currency := range updatedCurrencies {
			assert.Equal(t, expectedRates[currency.Code], currency.Rate)
		}
	})
}
