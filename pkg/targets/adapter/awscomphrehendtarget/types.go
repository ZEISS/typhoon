package awscomphrehendtarget

// Response is the data structure returned by the event target.
type Response struct {
	Positive float64 `json:"positive"`
	Negative float64 `json:"negative"`
	Mixed    float64 `json:"mixed"`
	Result   string  `json:"result"`
}
