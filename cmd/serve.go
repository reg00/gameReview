package cmd

import (
	"log"

	"github.com/Reg00/gameReview/internal/infrastructure/config"
	"github.com/Reg00/gameReview/internal/infrastructure/di"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run game review server",
	Long:  `Run game review server`,
	RunE: func(cmd *cobra.Command, args []string) error {
		server, err := di.InitWebServer()
		if err != nil {
			return err
		}
		server.Run()
		return nil
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}

func StartServer() error {
	cfg, err := config.LoadConfig()
	log.Println(cfg)
	if err != nil {
		return errors.Wrap(err, "error while loading configuration")
	}

	server, err := di.InitWebServer()
	if err != nil {
		return err
	}
	server.Run()
	return nil
}
