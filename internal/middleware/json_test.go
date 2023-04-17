package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DanielKirkwood/youchooseserver/internal/middleware"
)

// TestJson checks that the Content-Type value is
// application/json after Json middleware is used.
func TestJson(t *testing.T) {
	xmlType := "application/xml"
	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("[]"))
		if err != nil {
			t.Errorf("Error writing %v", err)
		}
	})

	// explicitly set Content-Type to value that is
	// not application/json
	req := httptest.NewRequest("GET", "/test", nil)
	req.Header.Set("Content-Type", xmlType)
	if ctype := req.Header.Get("Content-Type"); ctype == "application/json" {
		t.Error("content type header set to application/json before middleware applied")
	}

	rr := httptest.NewRecorder()

	handler := middleware.Json(testHandler)
	handler.ServeHTTP(rr, req)

	if ctype := rr.Header().Get("Content-Type"); ctype != "application/json" {
		t.Errorf("content type header does not match: got %v want %v",
			ctype, "application/json")
	}
}
