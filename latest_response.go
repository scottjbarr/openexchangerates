package openexchangerates

// LatestResponse is used for the response from the /latest.json endpoint.
type LatestResponse struct {
	Timestamp int64              `json:"timestamp"`
	Base      string             `json:"base"`
	Rates     map[string]float64 `json:"rates"`
}
