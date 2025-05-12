package cmd

import (
	"github.com/alfariiizi/go-service/internal/delivery/http"
	"github.com/alfariiizi/go-service/internal/delivery/route"
	"github.com/alfariiizi/go-service/internal/infrastructure/database"

	// "github.com/alfariiizi/go-service/internal/repository/repoadapter"
	// serviceadapter "github.com/alfariiizi/go-service/internal/service/adapter"
	"github.com/labstack/echo/v4"
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
	// repoadapter.NewUserRepository,
	// serviceadapter.NewUserService,
	echo.New,
	route.NewHttpApi,
	http.NewHttpServer,
)

var HttpServerStart = fx.Invoke(
	func(server http.HttpServer) {},
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
