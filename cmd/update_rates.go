package cmd

import (
	"context"
	"exchange-rate-calculator/configs"
	"exchange-rate-calculator/internal/build"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func updateRatesCmd(ctx context.Context, conf configs.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "update-rates",
		Short: "start updating rates",
		RunE: func(cmd *cobra.Command, args []string) error {
			builder := build.New(conf)
			ctx, cancel := context.WithCancel(ctx)
			defer cancel()

			go func() {
				builder.WaitShutdown(ctx)
				cancel()
			}()

			if err := builder.UpdateRatesCron(ctx); err != nil {
				return errors.Wrap(err, "build update rates cron")
			}

			return nil
		},
	}
}
