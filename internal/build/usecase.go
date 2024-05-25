package build

import (
	"context"

	"exchange-rate-calculator/internal/adapter/api"
	"exchange-rate-calculator/internal/repository/postgresql"
	"exchange-rate-calculator/internal/usecase"

	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func (b *Builder) CurrencyUsecase() (*usecase.CurrencyUsecaseImpl, error) {
	db, err := b.PostgresClient()
	if err != nil {
		return nil, err
	}

	db.Debug()

	currencyRepository := postgresql.NewCurrencyRepository(db)
	fastForexClient := api.NewFastForexClient(b.config.FastForexAPI, b.config.FastForexURL)

	return usecase.NewCurrencyUsecase(currencyRepository, fastForexClient), nil
}

func (b *Builder) PostgresClient() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(b.config.DSN), &gorm.Config{})
	if err != nil {
		return nil, errors.Wrap(err, "open db connection")
	}

	conn, err := db.DB()
	if err != nil {
		return nil, errors.Wrap(err, "Cannot get pgsql connection")
	}

	b.shutdown.add(func(_ context.Context) error {
		if err = conn.Close(); err != nil {
			return errors.Wrap(err, "close db connection")
		}

		return nil
	})

	return db, nil
}
