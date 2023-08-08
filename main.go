package main

import (
	"delta-import-deal-batch/cmd"
	"delta-import-deal-batch/config"
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"math/rand"
	"os"
	"time"
)

const maxTraversalLinks = 32 * (1 << 20)

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

var Commit string
var Version string

func main() {
	// get the config
	cfg := config.InitConfig()

	cfg.Common.Commit = Commit
	cfg.Common.Version = Version

	// get all the commands
	var commands []*cli.Command

	// clicd .
	commands = append(commands, cmd.ImportDealerCmd(&cfg)...)

	app := &cli.App{
		Commands:    commands,
		Name:        "delta batch import dealer",
		Description: "A file/directory car generator for Delta",
		Version:     fmt.Sprintf("%s+git.%s\n", cfg.Common.Version, cfg.Common.Commit),
		Usage:       "delta-import [command] [arguments]",
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
