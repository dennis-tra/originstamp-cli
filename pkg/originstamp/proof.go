package originstamp

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"path"
)

// GetProof .
func (c *Client) GetProof(ctx context.Context, hash string, currency Currency, proofType ProofType) (*ProofResponse, error) {

	request := struct {
		Currency   int64  `json:"currency"`
		HashString string `json:"hash_string"`
		ProofType  int64  `json:"proof_type"`
	}{
		Currency:   int64(currency),
		HashString: hash,
		ProofType:  int64(proofType),
	}

	payload, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	wrapper, err := c.Request(ctx, http.MethodPost, path.Join("timestamp", "proof", "url"), bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}

	proof := &ProofResponse{}
	err = json.Unmarshal(wrapper.Data, proof)
	if err != nil {
		return nil, err
	}

	return proof, nil
}

type ProofResponse struct {
	DownloadURL   string `json:"download_url"`
	FileName      string `json:"file_name"`
	FileSizeBytes int64  `json:"file_size_bytes"`
}
