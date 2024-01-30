package models

import (
	openapi "github.com/zeiss/typhoon/api"
)

// PaginatedListTeams ...
type PaginatedListTeams struct {
	Limit   *float32        `json:"limit,omitempty"`
	Offset  *float32        `json:"offset,omitempty"`
	Results *[]openapi.Team `json:"results,omitempty"`
	Total   *float32        `json:"total,omitempty"`
}
