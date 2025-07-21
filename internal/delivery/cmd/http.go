package cmd

import (
	"github.com/alfariiizi/vandor/internal/core"
	"github.com/alfariiizi/vandor/internal/delivery/http"
	"github.com/alfariiizi/vandor/internal/infrastructure"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "Run HTTP server",
	Run: func(cmd *cobra.Command, args []string) {
		app := fx.New(
			infrastructure.Module,
			core.Module,
			// HTTP Server
			http.Module,
		)
		app.Run()
	},
}

func init() {
	appCmd.AddCommand(httpCmd)
}
