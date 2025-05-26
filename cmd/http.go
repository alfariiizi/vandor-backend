package cmd

import (
	"github.com/alfariiizi/go-service/internal/delivery/http/api"
	http "github.com/alfariiizi/go-service/internal/delivery/http/server"
	"github.com/alfariiizi/go-service/internal/infrastructure/database"
	"github.com/alfariiizi/go-service/internal/repository"
	"github.com/alfariiizi/go-service/internal/service"
	"github.com/go-chi/chi/v5"
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
	database.CreateSQLCDB,
	repository.InitRepositories,
	service.InitServices,
	chi.NewMux,
	// echo.New,
	// route.NewHttpApi,
	api.NewHttpApi,
	http.NewHttpServer,
)

var HttpServerStart = fx.Invoke(
	func(server *http.HttpServer) {},
)

func HttpServerRun() {
	app := fx.New(
		HttpServerProvider,
		HttpServerStart,
	)
	app.Run()
}

func init() {
	rootCmd.AddCommand(httpCmd)
}
