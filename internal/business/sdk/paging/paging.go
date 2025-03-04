package paging

import (
	"fmt"
	"strconv"
)

const (
	defaultPageNumber  = 1
	defaultRowsPerPage = 10
	maxRowsPerPage     = 100
)

// Page represents the requested page and rows per page.
type Page struct {
	number int
	rows   int
}

var DefaultPage = Page{
	number: defaultPageNumber,
	rows:   defaultRowsPerPage,
}

// Parse parses the strings and validates the values are in reason.
func Parse(page string, rowsPerPage string) (Page, error) {
	number := defaultPageNumber
	if page != "" {
		var err error
		number, err = strconv.Atoi(page)
		if err != nil {
			return Page{}, fmt.Errorf("page conversion: %w", err)
		}
	}

	rows := defaultRowsPerPage
	if rowsPerPage != "" {
		var err error
		rows, err = strconv.Atoi(rowsPerPage)
		if err != nil {
			return Page{}, fmt.Errorf("rows conversion: %w", err)
		}
	}

	if number <= 0 {
		return Page{}, fmt.Errorf("page value too small, must be larger than 0")
	}

	if rows <= 0 {
		return Page{}, fmt.Errorf("rows value too small, must be larger than 0")
	}

	if rows > maxRowsPerPage {
		return Page{}, fmt.Errorf("rows value too large, must be less than %d", maxRowsPerPage)
	}

	p := Page{
		number: number,
		rows:   rows,
	}

	return p, nil
}

// MustParse creates a paging value for testing.
func MustParse(page string, rowsPerPage string) Page {
	pg, err := Parse(page, rowsPerPage)
	if err != nil {
		panic(err)
	}

	return pg
}

// String implements the stringer interface.
func (p Page) String() string {
	return fmt.Sprintf("page: %d rows: %d", p.number, p.rows)
}

// Number returns the page number.
func (p Page) Number() int {
	return p.number
}

// RowsPerPage returns the rows per page.
func (p Page) RowsPerPage() int {
	return p.rows
}

// Offset returns the offset for the query.
func (p Page) Offset() int {
	return (p.number - 1) * p.rows
}

func (p Page) IsZero() bool {
	return p == Page{}
}
