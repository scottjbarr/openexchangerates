package openexchangerates

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
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

// Latest returns the latest exchange rates for the given base currency.
func (c Client) Latest(p *LatestParams) (*LatestResponse, error) {
	// build the url
	url := c.buildLatestURL(p)

	r := LatestResponse{}

	if err := c.get(url, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

// Historical returns the historical rate data.
func (c Client) Historical(date time.Time, p *HistoricalParams) (*HistoricalResponse, error) {
	// build the url
	url := c.buildHistoricalURL(date, p)

	r := HistoricalResponse{}

	if err := c.get(url, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

// buildLatestURL returns a URL for the latest endpoint.
func (c Client) buildLatestURL(p *LatestParams) string {
	pathFormat := "%s/api/latest.json?app_id=%s"

	path := fmt.Sprintf(pathFormat, Host, c.AppID)

	url := path

	params := p.Encode()
	if len(params) > 0 {
		url = fmt.Sprintf("%s&%s", url, params)
	}

	return url
}

// buildHistoricalURL returns a URL for the historical endpoint.
func (c Client) buildHistoricalURL(date time.Time, p *HistoricalParams) string {
	// The base path format
	pathFormat := "%s/api/historical/%s.json?app_id=%s"

	d := date.Format("2006-01-02")

	path := fmt.Sprintf(pathFormat, Host, d, c.AppID)

	url := path

	params := p.Encode()
	if len(params) > 0 {
		url = fmt.Sprintf("%s&%s", url, params)
	}

	return url
}

// get is a helper to DRY up making requests to the API and unpacking the responses.
func (c Client) get(url string, response interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, &response)
}
