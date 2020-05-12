package pkg

import (
	"context"
	"fmt"

	"github.com/dennis-tra/originstamp-cli/pkg/originstamp"
	"github.com/urfave/cli/v2"
)

// UsageCmd defines the `proof` subcommand
func UsageCmd() *cli.Command {
	return &cli.Command{
		Name:        "usage",
		Description: "With this command you can view the current api usage.",
		Usage:       "Retrieve information about the current api usage.",
		Action:      UsageAction,
	}
}

// UsageAction contains the logic for the `usage` subcommand.
func UsageAction(ctx *cli.Context) error {

	apiKey := ctx.String(flagAPIKey)

	c := originstamp.NewClient(apiKey)

	s := defaultSpinner(" Getting usage data...")
	s.Start()
	resp, err := c.GetUsage(context.Background())
	s.Stop()
	if err != nil {
		return err
	}

	fmt.Printf("Consumed credits for the current month:\t\t%.1f\n", resp.ConsumedCredits)
	fmt.Printf("Remaining credits for the current month:\t%.1f\n", resp.RemainingCredits)
	fmt.Printf("Total number of credits per month:\t\t%.1f\n", resp.CreditsPerMonth)
	fmt.Println("---")
	fmt.Printf("You have consumed %.1f%% of your available credits\n", 100*resp.ConsumedCredits/resp.CreditsPerMonth)

	return nil
}
