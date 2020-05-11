package pkg

import (
	"context"
	"fmt"
	"github.com/dennis-tra/originstamp-cli/pkg/originstamp"
	"github.com/urfave/cli/v2"
)

func StampAction(ctx *cli.Context) error {

	apiKey := ctx.String(FLAG_API_KEY)

	c := originstamp.NewClient(apiKey)

	hashStr, err := parseHash(ctx)
	if err != nil {
		fmt.Println(err)
		fmt.Println()
		cli.ShowAppHelpAndExit(ctx, 1)
	}

	opts := &originstamp.CreateOptions{
		Comment: ctx.String(FLAG_COMMENT),
	}

	s := defaultSpinner(" Time stamping...")
	s.Start()
	resp, err := c.TimestampCreate(context.Background(), hashStr, opts)
	s.Stop()
	if err != nil {
		return err
	}

	if resp.Created {
		fmt.Println("Successfully initiated timestamp creation!")
	} else {
		fmt.Println("Hash already submitted!")
	}

	printTimestampResponse(ctx, resp)

	return nil
}
