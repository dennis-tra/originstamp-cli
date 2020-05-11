package pkg

import (
	"context"
	"fmt"
	"github.com/dennis-tra/originstamp-cli/pkg/originstamp"
	"github.com/urfave/cli/v2"
)

func StatusCmd() *cli.Command {
	return &cli.Command{
		Name:        "status",
		Description: "Retrieve timestamp information for a certain file",
		Usage:       "Retrieve timestamp information for a certain file",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     FLAG_HASH,
				Usage:    "Provide the hash string instead of file",
				Required: false,
			},
		},
		Action: StatusAction,
	}
}

func StatusAction(ctx *cli.Context) error {

	apiKey := ctx.String(FLAG_API_KEY)

	hashStr, err := parseHash(ctx)
	if err != nil {
		fmt.Println(err)
		fmt.Println()
		cli.ShowAppHelpAndExit(ctx, 1)
	}

	c := originstamp.NewClient(apiKey)

	s := defaultSpinner(" Getting status...")
	s.Start()
	resp, err := c.TimestampStatus(context.Background(), hashStr)
	s.Stop()
	if err != nil {
		return err
	}

	printTimestampResponse(ctx, resp)

	return nil
}
