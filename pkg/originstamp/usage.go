package originstamp

import (
	"context"
	"encoding/json"
	"net/http"
	"path"
)

type UsageData struct {
	ConsumedCredits  float64 `json:"consumed_credits"`
	CreditsPerMonth  float64 `json:"credits_per_month"`
	RemainingCredits float64 `json:"remaining_credits"`
}

// GetUsage fetches the current api usage
func (c *Client) GetUsage(ctx context.Context) (*UsageData, error) {

	wrapper, err := c.Request(ctx, http.MethodGet, path.Join("api_key", "usage"), nil)
	if err != nil {
		return nil, err
	}

	usage := &UsageData{}
	err = json.Unmarshal(wrapper.Data, usage)
	if err != nil {
		return nil, err
	}

	return usage, nil
}
