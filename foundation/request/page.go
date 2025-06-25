package request

import (
	"net/http"
	"strconv"

	"github.com/lucasHSantiago/go-shop-ms/foundation/cerr"
)

// Page represents the requested page and rows per page.
type Page struct {
	Number      int
	RowsPerPage int
}

// ParsePage parse the request for the page and rows query string. The
// defaults are provided as well.
func ParsePage(r *http.Request) (Page, error) {
	values := r.URL.Query()

	number := 1
	if page := values.Get("page"); page != "" {
		var err error
		number, err = strconv.Atoi(page)
		if err != nil {
			return Page{}, cerr.NewFieldsError("page", err)
		}
	}

	rowsPerPage := 10
	if rows := values.Get("rows"); rows != "" {
		var err error
		rowsPerPage, err = strconv.Atoi(rows)
		if err != nil {
			return Page{}, cerr.NewFieldsError("rows", err)
		}
	}

	return Page{
		Number:      number,
		RowsPerPage: rowsPerPage,
	}, nil
}
