package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// var AppCmd = &cobra.Command{
// 	Use:   "go-app",
// 	Short: "Go Application CLI",
// }

var appCmd = &cobra.Command{
	Use:   "go-app",
	Short: "Go Application CLI",
}

func Execute() {
	if err := appCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
