package pkg

import (
	"context"
	"fmt"
	"github.com/dennis-tra/originstamp-cli/pkg/originstamp"
	"github.com/urfave/cli/v2"
)

func ProofCmd() *cli.Command {
	return &cli.Command{
		Name:        "proof",
		Description: "Retrieve the timestamp proof for a certain file or hash",
		Usage:       "Retrieve the timestamp proof for a certain file or hash",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     FLAG_HASH,
				Usage:    "Provide the hash string instead of file",
				Required: false,
			},
			&cli.GenericFlag{
				Name:     FLAG_CURRENCY,
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
				Name:     FLAG_PROOF_TYPE,
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

func ProofAction(ctx *cli.Context) error {

	apiKey := ctx.String(FLAG_API_KEY)

	currency, err := originstamp.CurrencyFromString(ctx.String(FLAG_CURRENCY))
	if err != nil {
		return err
	}

	proofType, err := originstamp.ProofTypeFromString(ctx.String(FLAG_PROOF_TYPE))
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
