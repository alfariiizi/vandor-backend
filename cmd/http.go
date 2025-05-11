package cmd

import (
	"github.com/alfariiizi/go-service/internal/core/repository/repoadapter"
	serviceadapter "github.com/alfariiizi/go-service/internal/core/service/adapter"
	"github.com/alfariiizi/go-service/internal/delivery/adapter/http"
	portHttp "github.com/alfariiizi/go-service/internal/delivery/port/http"
	"github.com/alfariiizi/go-service/internal/infrastructure/database"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "Run HTTP server",
	Run: func(cmd *cobra.Command, args []string) {
		HttpServerRun()
	},
}

var HttpServerProvider = fx.Provide(
	database.NewGormDB,
	repoadapter.NewUserRepository,
	serviceadapter.NewUserService,
	http.NewHttpServer,
)

var HttpServerStart = fx.Invoke(
	func(server portHttp.HttpServer) {},
)

func HttpServerRun() {
	app := fx.New(
		HttpServerProvider,
		fx.Invoke(func(server portHttp.HttpServer) {}),
	)
	app.Run()
}

func init() {
	rootCmd.AddCommand(httpCmd)
}
