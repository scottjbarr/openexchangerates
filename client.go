package openexchangerates

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	// Host is the the host for the APIn
	Host = "https://openexchangerates.org"
)

// Client is used to communicate with the API
type Client struct {
	AppID string
}

// New returns a new Client with the given Application ID.
func New(appID string) Client {
	return Client{
		AppID: appID,
	}
}

// Latest return the latest exchange rates for the given base currency.
func (c Client) Latest(base string) (LatestResponse, error) {
	url := fmt.Sprintf("%s/api/latest.json?base=%s&app_id=%s", Host, base, c.AppID)
	resp, err := http.Get(url)
	if err != nil {
		return LatestResponse{}, err
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return LatestResponse{}, err
	}

	lr := LatestResponse{}
	if err := json.Unmarshal(b, &lr); err != nil {
		return LatestResponse{}, err
	}

	return lr, nil
}
