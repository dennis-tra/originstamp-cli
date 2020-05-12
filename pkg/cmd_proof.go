package pkg

import (
	"context"
	"fmt"

	"github.com/dennis-tra/originstamp-cli/pkg/originstamp"
	"github.com/urfave/cli/v2"
)

// ProofCmd defines the `proof` subcommand
func ProofCmd() *cli.Command {
	return &cli.Command{
		Name:        "proof",
		Description: "Retrieve the timestamp proof for a certain file or hash",
		Usage:       "Retrieve the timestamp proof for a certain file or hash",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     flagHash,
				Usage:    "Provide the hash string instead of a file",
				Required: false,
			},
			&cli.GenericFlag{
				Name:     flagCurrency,
				Aliases:  []string{"c"},
				Usage:    "The currency for which you want to download the proof",
				Required: false,
				Value: &EnumValue{
					Enum:    originstamp.CurrencyStrValues(),
					Default: originstamp.BITCOIN.String(),
				},
				DefaultText: "",
				HasBeenSet:  false,
			},
			&cli.GenericFlag{
				Name:     flagProofType,
				Aliases:  []string{"p"},
				Usage:    "The proof type that you want to download",
				Required: false,
				Value: &EnumValue{
					Enum:    originstamp.ProofTypeStrValues(),
					Default: originstamp.SEED.String(),
				},
			},
		},
		Action: ProofAction,
	}
}

// ProofAction contains the logic for the `proof` subcommand.
func ProofAction(ctx *cli.Context) error {

	apiKey := ctx.String(flagAPIKey)

	currency, err := originstamp.CurrencyFromString(ctx.String(flagCurrency))
	if err != nil {
		return err
	}

	proofType, err := originstamp.ProofTypeFromString(ctx.String(flagProofType))
	if err != nil {
		return err
	}

	hashStr, err := parseHash(ctx)
	if err != nil {
		fmt.Println(err)
		fmt.Println()
		cli.ShowAppHelpAndExit(ctx, 1)
	}

	c := originstamp.NewClient(apiKey)

	s := defaultSpinner(" Getting proof url...")
	s.Start()
	resp, err := c.GetProof(context.Background(), hashStr, currency, proofType)
	s.Stop()
	if err != nil {
		return err
	}

	if resp.DownloadURL == "" {
		return cli.Exit("Download URL not found", 1)
	}
	fmt.Println(resp.DownloadURL)

	return nil
}
