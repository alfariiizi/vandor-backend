package cmd

import (
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

var serviceCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run all services",
	Run: func(cmd *cobra.Command, args []string) {
		ServiceRun()
	},
}

func ServiceRun() {
	app := fx.New(
		HttpServerProvider,
		HttpServerStart,
	)
	app.Run()
}

func init() {
	rootCmd.AddCommand(serviceCmd)
}
