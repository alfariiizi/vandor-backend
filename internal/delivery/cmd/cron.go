package cmd

import (
	"github.com/alfariiizi/go-service/internal/core"
	"github.com/alfariiizi/go-service/internal/cron"
	"github.com/alfariiizi/go-service/internal/infrastructure"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

var cronCmd = &cobra.Command{
	Use:   "cron",
	Short: "Run All Cron Jobs",
	Run: func(cmd *cobra.Command, args []string) {
		app := fx.New(
			infrastructure.Module,
			core.Module,
			// Cron
			cron.Module,
		)
		app.Run()
	},
}

func init() {
	appCmd.AddCommand(cronCmd)
}
