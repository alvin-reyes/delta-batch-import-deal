// Package cmd The `CarChunkRunnerCmd` function generates car files from a given file or directory and splits the process into batches
// for parallel processing.
package cmd

import (
	"delta-import-deal-batch/config"
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
)

func ImportDealerCmd(config *config.DeltaBatchImportDealerConfig) []*cli.Command {
	var carCommands []*cli.Command
	carChunkerCmd := &cli.Command{
		Name:  "dealer",
		Usage: "Import deals from a given file or directory and split the process into batches for parallel processing.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "run-config",
				Usage: "path to run config file",
			},
			&cli.IntFlag{
				Name:  "interval",
				Value: 5,
				Usage: "interval in seconds between each batch",
			},
		},
		Action: func(c *cli.Context) error {
			configFile := c.String("run-config")

			data, err := os.ReadFile(configFile)
			if err != nil {
				fmt.Println("Error reading config file:", err)
				return err
			}
			fmt.Println("config file:", configFile)
			fmt.Println("config data:", string(data))
			return nil
		},
	}
	carCommands = append(carCommands, carChunkerCmd)

	return carCommands
}
