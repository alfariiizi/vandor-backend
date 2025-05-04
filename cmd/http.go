package cmd

import (
	"github.com/alfariiizi/go-echo-fx-template/config"
	"github.com/alfariiizi/go-echo-fx-template/internal/core/repository/repoadapter"
	serviceadapter "github.com/alfariiizi/go-echo-fx-template/internal/core/service/adapter"
	"github.com/alfariiizi/go-echo-fx-template/internal/delivery/adapter/http"
	portHttp "github.com/alfariiizi/go-echo-fx-template/internal/delivery/port/http"
	"github.com/alfariiizi/go-echo-fx-template/internal/infrastructure/database"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "Run HTTP server",
	Run: func(cmd *cobra.Command, args []string) {
		Run()
	},
}

func Run() {
	app := fx.New(
		fx.Provide(config.NewConfig),

		// database
		fx.Provide(database.NewGormDB),

		// repositories
		fx.Provide(repoadapter.NewUserRepository),

		// services
		fx.Provide(serviceadapter.NewUserService),

		// http server
		fx.Provide(http.NewHttpServer),

		fx.Invoke(func(server portHttp.HttpServer) {}),
	)
	app.Run()
}

func init() {
	rootCmd.AddCommand(httpCmd)
}
