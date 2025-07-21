package cmd

import (
	"github.com/alfariiizi/vandor/internal/infrastructure/db"
	"github.com/alfariiizi/vandor/internal/seeder"
	"github.com/spf13/cobra"
)

var seederCmd = &cobra.Command{
	Use:   "seed",
	Short: "Run Seeder",
	Run: func(cmd *cobra.Command, args []string) {
		client := db.NewDB()
		seeder.GenerateAdmin(client)
	},
}

func init() {
	appCmd.AddCommand(seederCmd)
}
