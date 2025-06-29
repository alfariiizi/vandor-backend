package cmd

import (
	"github.com/spf13/cobra"
)

var seederCmd = &cobra.Command{
	Use:   "seed",
	Short: "Run ",
	Run: func(cmd *cobra.Command, args []string) {
		// client := repository.NewClient()
		// seeder.GenerateAdmin(client)
	},
}

func init() {
	appCmd.AddCommand(seederCmd)
}
