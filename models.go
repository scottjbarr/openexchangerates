package openexchangerates

import (
	"fmt"
	"net/url"
	"strings"
)

// BaseResponse is a wrapper for error details, and nomenclature such as license information.
type BaseResponse struct {
	ErrorResponse
	Nomenclature
}

// ErrorResponse is returned by the API when the request could not be fulfulled.
type ErrorResponse struct {
	Error       bool   `json:"error,omitempty"`
	Status      int    `json:"status,omitempty"`
	Message     string `json:"message,omitempty"`
	Description string `json:"description,omitempty"`
}

// Nomenclature is returned by the API with any valid API request.
type Nomenclature struct {
	Disclaimer string `json:"disclaimer,omitempty"`
	License    string `json:"license,omitempty"`
}

// LatestParams is passed to the client to specify fields in calls to the "latest" endpoint.
type LatestParams struct {
	Base            string
	Symbols         []string
	PrettyPrint     bool
	ShowAlternative bool
}

// Encode returns the values as an encoded url string.
func (p *LatestParams) Encode() string {
	v := url.Values{}

	if p == nil {
		return ""
	}

	if len(p.Base) > 0 {
		v.Add("base", p.Base)
	}

	if len(p.Symbols) > 0 {
		v.Add("symbols", strings.Join(p.Symbols, ","))
	}

	v.Add("show_alternative", fmt.Sprintf("%v", p.ShowAlternative))

	// defaults to true, set it to be sure because mostly we don't beed pretty print.
	v.Add("prettyprint", fmt.Sprintf("%v", p.PrettyPrint))

	return v.Encode()
}

// LatestResponse is used for the response from the "latest" endpoint.
type LatestResponse struct {
	BaseResponse
	LatestData
}

// LatestData is the rates specific to the "latest" endpoint.
type LatestData struct {
	Timestamp int64              `json:"timestamp"`
	Base      string             `json:"base,omitempty"`
	Rates     map[string]float64 `json:"rates"`
}

// HistoricalParams is passed to the client to specify fields in calls to the "historical" endpoint.
type HistoricalParams struct {
	Base            string
	Symbols         []string
	PrettyPrint     bool
	ShowAlternative bool
}

// Encode returns the values as an encoded url string.
func (p *HistoricalParams) Encode() string {
	v := url.Values{}

	if p == nil {
		return ""
	}

	if len(p.Base) > 0 {
		v.Add("base", p.Base)
	}

	if len(p.Symbols) > 0 {
		v.Add("symbols", strings.Join(p.Symbols, ","))
	}

	v.Add("show_alternative", fmt.Sprintf("%v", p.ShowAlternative))

	// defaults to true, set it to be sure because mostly we don't beed pretty print.
	v.Add("prettyprint", fmt.Sprintf("%v", p.PrettyPrint))

	return v.Encode()
}

// HistoricalResponse is used for the response from the "historical" endpoint.
type HistoricalResponse struct {
	BaseResponse
	HistoricalData
}

// HistoricalData is the rates specific data from the "historical" endpoint.
type HistoricalData struct {
	Timestamp int64              `json:"timestamp,omitempty"`
	Base      string             `json:"base,omitempty"`
	Rates     map[string]float64 `json:"rates,omitempty"`
}
