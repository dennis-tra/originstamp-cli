package pkg

import (
	"fmt"
	"github.com/urfave/cli/v2" // imports as package "cli"
	"os"
	"time"
)

func Run() {
	app := &cli.App{
		Name:     "stamp",
		HelpName: "stamp",
		Usage:    "create anonymous, tamper-proof timestamps for any digital content",
		Version:  VERSION,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     FLAG_API_KEY,
				Aliases:  []string{"k"},
				Usage:    "The OriginStamp AG API-Key (also applies to all sub commands)",
				EnvVars:  []string{"ORIGINSTAMP_API_KEY"},
				Required: true,
			},
			&cli.StringFlag{
				Name:     FLAG_COMMENT,
				Aliases:  []string{"c"},
				Usage:    "Comment (max. 256 chars) for the timestamp for indexing and searching (public)",
				Required: false,
			},
			&cli.StringFlag{
				Name:  FLAG_FORMAT,
				Usage: "Go layout of how to format timestamp when printed to the screen. See https://golang.org/pkg/time/#pkg-constants. Defaults to RFC3339.",
				Value: time.RFC3339,
			},
			&cli.StringFlag{
				Name:     FLAG_HASH,
				Usage:    "Provide the hash string instead of a file",
				Required: false,
			},
		},
		Action: StampAction,
		Authors: []*cli.Author{
			{
				Name:  "Dennis Trautwein",
				Email: "dennis.trautwein@originstamp.com",
			},
		},
		Copyright: "OriginStamp AG 2020",
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
