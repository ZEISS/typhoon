package models

import (
	"math"

	"gorm.io/gorm"
)

// Pagination is a struct that contains the pagination information.
type Pagination[R any] struct {
	// Limit is the number of items to return.
	Limit int `json:"limit" xml:"limit" form:"limit"`
	// Offset is the number of items to skip.
	Offset int `json:"offset" xml:"offset" form:"offset"`
	// Search is the search term to filter the results.
	Search string `json:"search" xml:"search" form:"search"`
	// Sort is the sorting order.
	Sort string `json:"sort,omitempty" xml:"sort" form:"sort"`
	// TotalRows is the total number of rows.
	TotalRows int `json:"total_rows"`
	// TotalPages is the total number of pages.
	TotalPages int `json:"total_pages"`
	// Rows is the items to return.
	Rows []R `json:"rows"`
}

// NewPagination returns a new pagination.
func NewPagination[R any]() Pagination[R] {
	return Pagination[R]{}
}

// GetLimit returns the limit.
func (p *Pagination[R]) GetLimit() int {
	if p.Limit == 0 {
		p.Limit = 10
	}

	return p.Limit
}

// GetOffset returns the page.
func (p *Pagination[R]) GetOffset() int {
	if p.Offset < 0 {
		p.Offset = 0
	}

	return p.Offset
}

// GetSort returns the sort.
func (p *Pagination[R]) GetSort() string {
	if p.Sort == "" {
		p.Sort = "desc"
	}

	return p.Sort
}

// Paginate returns a function that paginates the results.
func Paginate[R any](value interface{}, pagination *Pagination[R], db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var totalRows int64
	db.Model(value).Count(&totalRows).Where("deleted_at IS NULL")

	pagination.TotalRows = int(totalRows)
	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.Limit)))
	pagination.TotalPages = totalPages

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit())
	}
}
