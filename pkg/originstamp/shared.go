package originstamp

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"path"
)

type ResponseWrapper struct {
	Data         json.RawMessage `json:"data"`
	ErrorCode    int64           `json:"error_code"`
	ErrorMessage string          `json:"error_message"`
}

func (c *Client) Request(ctx context.Context, method string, endpoint string, body io.Reader) (*ResponseWrapper, error) {

	u := baseURL()
	u.Path = path.Join(u.Path, endpoint)

	req, err := http.NewRequestWithContext(ctx, method, u.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", c.apiKey)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "OriginStampCLI/"+version)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			fmt.Println("Error closing response body", err)
		}
	}()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	wrapper := &ResponseWrapper{}
	err = json.Unmarshal(data, wrapper)
	if err != nil {
		return nil, err
	}

	if wrapper.ErrorCode != 0 {
		return nil, Error{
			Code:    int(wrapper.ErrorCode),
			Message: wrapper.ErrorMessage,
		}
	}

	return wrapper, nil
}
