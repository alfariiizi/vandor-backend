package cmd

import (
	"github.com/alfariiizi/vandor/internal/core"
	"github.com/alfariiizi/vandor/internal/delivery/worker"
	"github.com/alfariiizi/vandor/internal/infrastructure"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

var workerCmd = &cobra.Command{
	Use:   "worker",
	Short: "Run worker server",
	Run: func(cmd *cobra.Command, args []string) {
		app := fx.New(
			infrastructure.Module,
			core.Module,
			// worker Server
			worker.Module,
		)
		app.Run()
	},
}

func init() {
	appCmd.AddCommand(workerCmd)
}
