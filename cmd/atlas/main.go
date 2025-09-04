package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/alfariiizi/vandor/internal/config"
	"github.com/spf13/cobra"
)

var (
	migrationsDir = "file://database/migrate/migrations"
	entSchemaURL  = "ent://database/schema"
	devURL        = "docker://postgres/17/test?search_path=public"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "atlasgo",
		Short: "Atlas migration",
	}

	// Register subcommands
	rootCmd.AddCommand(newDiffCmd())
	rootCmd.AddCommand(newApplyCmd())
	rootCmd.AddCommand(newStatusCmd())
	rootCmd.AddCommand(newLintCmd())

	// Load config
	_ = config.GetConfig()

	// viper.SetConfigName("config")
	// viper.SetConfigType("yaml")
	// viper.AddConfigPath(".")
	// if err := viper.ReadInConfig(); err != nil {
	// 	fmt.Fprintf(os.Stderr, "warning: cannot read config file: %v\n", err)
	// }

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("command execution failed: %v", err)
		os.Exit(1)
	}
}

// --- SUBCOMMANDS ---

func newDiffCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "diff [name]",
		Short: "Create a new migration from schema",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			migrationName := args[0]
			atlasCmd := exec.Command("atlas", "migrate", "diff", migrationName,
				"--dir", migrationsDir,
				"--to", entSchemaURL,
				"--dev-url", devURL,
			)
			atlasCmd.Env = os.Environ()
			out, err := atlasCmd.CombinedOutput()
			if err != nil {
				return fmt.Errorf("atlas failed: %w\nOutput:\n%s", err, string(out))
			}
			fmt.Print(string(out))
			return nil
		},
	}
}

func newApplyCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "apply",
		Short: "Apply migrations to the database",
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg := config.GetConfig()
			dbURL := cfg.DB.URL
			atlasCmd := exec.Command("atlas", "migrate", "apply",
				"--dir", migrationsDir,
				"--url", dbURL,
			)
			atlasCmd.Stdout = os.Stdout
			atlasCmd.Stderr = os.Stderr
			atlasCmd.Env = append(os.Environ(),
				"DB_URL="+dbURL,
			)
			return atlasCmd.Run()
		},
	}
}

func newStatusCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "status",
		Short: "Check migration status",
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg := config.GetConfig()
			dbURL := cfg.DB.URL
			atlasCmd := exec.Command("atlas", "migrate", "status",
				"--dir", migrationsDir,
				"--url", dbURL,
			)
			atlasCmd.Stdout = os.Stdout
			atlasCmd.Stderr = os.Stderr
			atlasCmd.Env = append(os.Environ(),
				"DB_URL="+dbURL,
			)
			return atlasCmd.Run()
		},
	}
}

func newLintCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "lint",
		Short: "Lint migration directory",
		RunE: func(cmd *cobra.Command, args []string) error {
			atlasCmd := exec.Command("atlas", "migrate", "lint",
				"--dev-url", devURL,
				"--dir", migrationsDir,
			)
			atlasCmd.Stdout = os.Stdout
			atlasCmd.Stderr = os.Stderr
			atlasCmd.Env = os.Environ()
			return atlasCmd.Run()
		},
	}
}
