package pkg

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/briandowns/spinner"
	"github.com/dennis-tra/originstamp-cli/pkg/originstamp"
	"github.com/urfave/cli/v2"
)

func defaultSpinner(suffix string) *spinner.Spinner {
	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond, spinner.WithWriter(os.Stderr))
	s.Suffix = suffix
	return s
}

func parseHash(ctx *cli.Context) (string, error) {
	hashStr := ctx.String(flagHash)
	if hashStr != "" {
		return hashStr, nil
	}
	if ctx.Args().Len() != 1 {
		return "", fmt.Errorf("please provide either a file name or hash string")
	}
	filePath := ctx.Args().First()

	f, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}

	hashBytes := sha256.Sum256(data)
	hashStr = fmt.Sprintf("%x", hashBytes[:])

	return hashStr, nil
}

func printTimestampResponse(ctx *cli.Context, resp *originstamp.TimestampResponse) {
	fmt.Printf("%10s | %12s | %25s | %s\n", "CURRENCY", "STATUS", "TIMESTAMP", "TRANSACTION")
	fmt.Printf("%10s | %12s | %25s | %s\n", "---", "---", "---", "---")
	for _, ts := range resp.Timestamps {
		timestamp := ts.Timestamp().Format(ctx.String(flagFormat))
		if ts.SubmitStatus < originstamp.INCLUDED {
			timestamp = ""
		}
		fmt.Printf("%10s | %12s | %25s | %s\n", ts.CurrencyID, ts.SubmitStatus, timestamp, ts.Transaction)
	}
}
