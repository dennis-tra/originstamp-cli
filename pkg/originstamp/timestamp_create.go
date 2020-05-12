package originstamp

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"path"
	"time"
)

// Add a notification/notification list to your submission. Our system will
// notify the specified target with the timestamp information.
// Depending on the notification type, specify the target for the notification
// (e.g. mail address or webhook URL). The webhook URL will receive a POST request
// after timestamp was successfully created.
type Notification struct {
	Currency         Currency         `json:"currency"`
	NotificationType NotificationType `json:"notification_type"`
	Target           string           `json:"target"`
}

type CreateOptions struct {
	// You can add a short comment (max. 256 characters) which can be used for indexing and searching (public).
	Comment       string
	Notifications []Notification
	URL           string
}

// TimestampCreate is used to submit your hash. If your api key is valid, your hash is added to
// batch and is scheduled for timestamping. If the hash already exists, the created flag in the
// response is set to false and the notification(s) of the current request will be totally ignored.
// You are also able to submit additional information, such as comment or notification credentials.
// Once a hash is successfully created for a certain crypto-currency, we can notify your desired
// target with the timestamp information (POST Request). The webhook is triggered as soon as the
// tamper-proof timestamp with the selected crypto currency has been created. Additionally, it is
// possible to include a preprint URL in the hash submission. Before the generation of the timestamp
// hash you can create a random UUID Version 4 and include https://originstamp.com/u/UUID where UUID
// is your UUID e.g. in a document you want to timestamp. In the preprint URL field you include your
// UUID and then it is possible to verify the timestamp within the document (or whatever).
func (c *Client) TimestampCreate(ctx context.Context, hash string, options *CreateOptions) (*TimestampResponse, error) {

	type createRequest struct {
		Comment       string         `json:"comment,omitempty"`
		Hash          string         `json:"hash"`
		Notifications []Notification `json:"notifications,omitempty"`
		URL           string         `json:"url,omitempty"`
	}

	cReq := &createRequest{
		Hash: hash,
	}

	if options != nil {
		cReq.Comment = options.Comment
		cReq.URL = options.URL
		cReq.Notifications = options.Notifications
	}

	payload, err := json.Marshal(cReq)
	if err != nil {
		return nil, err
	}

	endpoint := path.Join("timestamp", "create")
	wrapper, err := c.Request(ctx, http.MethodPost, endpoint, bytes.NewReader(payload))
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

type TimestampResponse struct {

	// The comment which was added in the submission of the hash.
	Comment string `json:"comment"`

	// Field is set to true if it is a novel hash.If the flag is false,
	//the hash was already submitted before.
	Created bool `json:"created"`

	// The time when your hash was submitted to OriginStamp.
	// The date is returned in the following format: [ms] since
	// 1.1.1970 (unix epoch), timezone: UTC. This is not
	// considered as a true timestamp
	DateCreated int64 `json:"date_created"`

	// The submitted hash in hex representation.
	HashString string `json:"hash_string"`

	// Contains all the timestamp data of your hash until now.
	Timestamps []Timestamp `json:"timestamps"`
}

type Timestamp struct {
	CurrencyID Currency `json:"currency_id"`

	// The private key represents the top hash in the Merkle Tree
	// (see https://en.wikipedia.org/wiki/Merkle_tree ) or the hash of all hashes in the transaction.
	PrivateKey string `json:"private_key"`

	// The submit status of the hash
	SubmitStatus SubmitStatus `json:"submit_status"`

	// The date is returned in the following format: [ms] since 1.1.1970 (unix epoch), timezone: UTC. This is a true timestamp.
	TimestampUnix int64 `json:"timestamp"`
	// If available: the transaction hash of the timestamp.
	Transaction string `json:"transaction"`
}

func (t *Timestamp) Timestamp() time.Time {
	return time.Unix(t.TimestampUnix*int64(time.Millisecond/time.Second), t.TimestampUnix*int64(time.Millisecond/time.Nanosecond))
}
