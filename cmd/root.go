package cmd

import (
	"context"
	"exchange-rate-calculator/configs"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func Run(ctx context.Context, conf configs.Config) error {
	root := &cobra.Command{
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Usage()
		},
	}

	root.AddCommand(
		updateRatesCmd(ctx, conf),
		restCmd(ctx, conf),
	)

	return errors.Wrap(root.ExecuteContext(ctx), "run application")
}
