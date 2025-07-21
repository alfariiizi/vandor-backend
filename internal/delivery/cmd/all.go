package cmd

import (
	"github.com/alfariiizi/vandor/internal/core"
	"github.com/alfariiizi/vandor/internal/cron"
	"github.com/alfariiizi/vandor/internal/delivery/http"
	"github.com/alfariiizi/vandor/internal/infrastructure"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

var allServiceCmd = &cobra.Command{
	Use:   "all",
	Short: "Run all services. Perfect for development and testing.",
	Run: func(cmd *cobra.Command, args []string) {
		app := fx.New(
			infrastructure.Module,
			core.Module,
			// HTTP Server
			http.Module,
			// Cron Jobs
			cron.Module,
		)
		app.Run()
	},
}

func init() {
	appCmd.AddCommand(allServiceCmd)
}
