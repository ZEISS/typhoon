package datadogtarget

// MetricData represents the body of a Datadog metric (time-series).
// Ref. https://docs.datadoghq.com/api/latest/metrics/#submit-metrics
//
// The data of CloudEvents of type "com.zeiss.datadog.metric" must satisfy
// this schema.
type MetricData struct {
	Series []Series `json:"series"`
}

// Series is an individual metric series.
type Series struct {
	Host     *string    `json:"host,omitempty"`
	Interval *int       `json:"interval,omitempty"`
	Metric   string     `json:"metric"`
	Points   [][]string `json:"points"`
	Tags     []string   `json:"tags,omitempty"`
	Type     *string    `json:"type,omitempty"`
}

// EventData represents the body of a Datadog event
// Ref. https://docs.datadoghq.com/api/latest/events/#post-an-event
//
// The data of CloudEvents of type "com.zeiss.datadog.event.post" must satisfy
// this schema.
type EventData struct {
	AlertType      *string  `json:"alert_type,omitempty"`
	DateHappened   *int     `json:"date_happened,omitempty"`
	DeviceName     *string  `json:"device_name,omitempty"`
	Host           *string  `json:"host,omitempty"`
	ID             *int     `json:"id,omitempty"`
	Priority       *string  `json:"priority,omitempty"`
	RelatedEventID *int     `json:"related_event_id,omitempty"`
	SourceTypeName *string  `json:"source_type_name,omitempty"`
	Status         *string  `json:"status,omitempty"`
	Tags           []string `json:"tags,omitempty"`
	Text           string   `json:"text"`
	Title          string   `json:"title"`
	URL            *string  `json:"url,omitempty"`
}

// LogData represents the body of a Datadog log entry
// Ref. https://docs.datadoghq.com/api/latest/logs/#send-logs
//
// The data of CloudEvents of type "com.zeiss.datadog.logs.send" must satisfy
// this schema.
type LogData struct {
	DdSource string `json:"ddsource"`
	DdTags   string `json:"ddtags"`
	Hostname string `json:"hostname"`
	Message  string `json:"message"`
	Service  string `json:"service"`
}
