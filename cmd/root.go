package cmd

import (
	"log"
	"os"

	"github.com/Reg00/gameReview/internal/infrastructure/config"
	"github.com/pkg/errors"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   config.ProjectName,
	Short: config.ProjectName + " service",
	Long:  config.ProjectName + " is an service, where you can review games.",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return loadConfig()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(serverCmd)
}

func loadConfig() error {
	cfg, err := config.LoadConfig()
	log.Println(cfg)
	if err != nil {
		return errors.Wrap(err, "error while loading configuration")
	}
	return err
}
