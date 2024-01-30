package models

// PaginatedResult ...
type PaginatedResult struct {
	Limit  float32 `json:"limit,omitempty"`
	Offset float32 `json:"offset,omitempty"`
	Total  float32 `json:"total,omitempty"`
}
