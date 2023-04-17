package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DanielKirkwood/youchooseserver/internal/middleware"
)

func TestCORS(t *testing.T) {
	type test struct {
		header string
		want   string
	}

	tests := []test{
		{header: "Access-Control-Allow-Origin", want: "*"},
		{header: "Access-Control-Allow-Methods", want: "*"},
		{header: "Access-Control-Allow-Headers", want: "*"},
		{header: "Access-Control-Allow-Credentials", want: "true"},
		{header: "Access-Control-Max-Age", want: "300"},
	}

	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("hello world!"))
		if err != nil {
			t.Errorf("Error writing %v", err)
		}
	})


	req := httptest.NewRequest("GET", "/test", nil)
	rr := httptest.NewRecorder()

	handler := middleware.CORS(testHandler)
	handler.ServeHTTP(rr, req)

	for _, tc := range tests {
		got := rr.Header().Get(tc.header)
		if got != tc.want {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}
