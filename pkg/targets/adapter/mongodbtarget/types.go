package mongodbtarget

import "encoding/json"

// InsertPayload defines the expected data structure found at the "io.triggermesh.mongodb.insert" payload.
type InsertPayload struct {
	Database   string          `json:"database"`
	Collection string          `json:"collection"`
	Key        string          `json:"key"`
	Document   json.RawMessage `json:"document"`
}

// QueryPayload defines the expected data found at the "io.triggermesh.mongodb.query" payload.
type QueryPayload struct {
	Database   string `json:"database"`
	Collection string `json:"collection"`
	Key        string `json:"key"`
	Value      string `json:"value"`
}

// UpdatePayload defines the expected data found at the "io.triggermesh.mongodb.update" payload.
type UpdatePayload struct {
	Database    string `json:"database"`
	Collection  string `json:"collection"`
	SearchKey   string `json:"searchKey"`
	SearchValue string `json:"searchValue"`
	UpdateKey   string `json:"updateKey"`
	UpdateValue string `json:"updateValue"`
}

// QueryResponse defines the expected data structure received from a query to MongoDB.
type QueryResponse struct {
	Collection map[string]string `json:"collection"`
}
