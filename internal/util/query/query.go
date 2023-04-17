package query

import (
	"net/http"
	"strconv"
)

func Int(r *http.Request, query string) (int, error) {
	val, err := strconv.Atoi(r.URL.Query().Get(query))
	if err != nil {
		return 0, err
	}

	return val, nil
}

func String(r *http.Request, query string) string {
	return r.URL.Query().Get(query)
}
