package query

import (
	"github.com/machilan1/plpr2/internal/business/sdk/paging"
)

// Result is the data model used when returning a query result.
type Result[T any] struct {
	Items    []T `json:"items"`
	Total    int `json:"total"`
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
}

// NewResult constructs a result value to return query results.
func NewResult[T any](items []T, total int, page paging.Page) Result[T] {
	return Result[T]{
		Items:    items,
		Total:    total,
		Page:     page.Number(),
		PageSize: page.RowsPerPage(),
	}
}
