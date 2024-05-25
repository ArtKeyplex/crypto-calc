package build

import (
	"context"

	"github.com/pkg/errors"
	"github.com/robfig/cron/v3"
	"github.com/rs/zerolog"
)

func (b *Builder) UpdateRatesCron(ctx context.Context) error {
	cronJob := cron.New()
	currencyUsecase, err := b.CurrencyUsecase()
	if err != nil {
		return errors.Wrap(err, "initialize currency usecase")
	}

	_, err = cronJob.AddFunc("@every 1m", func() {
		err := currencyUsecase.UpdateRates(ctx)
		if err != nil {
			zerolog.Ctx(ctx).Error().Err(err).Msg("Failed to update rates")
		}

		zerolog.Ctx(ctx).Info().Msg("Rates updated")
	})
	if err != nil {
		return errors.Wrap(err, "add cron job")
	}

	cronJob.Start()

	<-ctx.Done()

	defer cronJob.Stop()

	return nil
}
