package cmd

import (
	"context"
	"exchange-rate-calculator/configs"
	"exchange-rate-calculator/internal/build"
	"net/http"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func restCmd(ctx context.Context, conf configs.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "rest",
		Short: "start rest server",
		RunE: func(cmd *cobra.Command, args []string) error {
			builder := build.New(conf)
			ctx, cancel := context.WithCancel(ctx)
			defer cancel()

			go func() {
				builder.WaitShutdown(ctx)
				cancel()
			}()

			server, err := builder.RestAPIServer(ctx)
			if err != nil {
				return errors.Wrap(err, "build rest api server")
			}

			if err = server.Listen(":" + conf.Port); err != nil && !errors.Is(err, http.ErrServerClosed) {
				return errors.Wrap(err, "rest api server serve")
			}

			<-ctx.Done()

			return nil
		},
	}
}
