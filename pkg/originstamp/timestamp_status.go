package originstamp

import (
	"context"
	"encoding/json"
	"net/http"
	"path"
)

func (c *Client) TimestampStatus(ctx context.Context, hash string) (*TimestampResponse, error) {

	wrapper, err := c.Request(ctx, http.MethodGet, path.Join("timestamp", hash), nil)
	if err != nil {
		return nil, err
	}

	resp := &TimestampResponse{}
	err = json.Unmarshal(wrapper.Data, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
