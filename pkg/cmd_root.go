package pkg

import (
	"fmt"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

// Run initializes the stamp command line app
func Run() {
	app := &cli.App{
		Name:     "stamp",
		HelpName: "stamp",
		Usage:    "create anonymous, tamper-proof timestamps for any digital content",
		Version:  version,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     flagAPIKey,
				Aliases:  []string{"k"},
				Usage:    "The OriginStamp AG API-Key (also applies to all sub commands)",
				EnvVars:  []string{"ORIGINSTAMP_API_KEY"},
				Required: true,
			},
			&cli.StringFlag{
				Name:     flagComment,
				Aliases:  []string{"c"},
				Usage:    "Comment (max. 256 chars) for the timestamp for indexing and searching (public)",
				Required: false,
			},
			&cli.StringFlag{
				Name:  flagFormat,
				Usage: "Go layout of how to format timestamp when printed to the screen. See https://golang.org/pkg/time/#pkg-constants. Defaults to RFC3339.",
				Value: time.RFC3339,
			},
			&cli.StringFlag{
				Name:     flagHash,
				Usage:    "Provide the hash string instead of a file",
				Required: false,
			},
		},
		Action: StampAction,
		Authors: []*cli.Author{
			{
				Name:  "Dennis Trautwein",
				Email: "originstamp-cli@dtrautwein.eu",
			},
		},
		Commands: []*cli.Command{
			UsageCmd(),
			StatusCmd(),
			ProofCmd(),
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
