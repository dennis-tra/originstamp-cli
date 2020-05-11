package originstamp

import (
	"net/http"
	"net/url"
)

type Client struct {
	httpClient *http.Client
	apiKey     string
}

func NewClient(apiKey string) *Client {
	return &Client{
		httpClient: &http.Client{},
		apiKey:     apiKey,
	}
}

func baseURL() *url.URL {
	return &url.URL{
		Scheme: "https",
		Host:   "api.originstamp.com",
		Path:   "v3",
	}
}
