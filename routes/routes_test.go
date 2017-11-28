package routes

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/franela/goblin"
)

func performRequest(r http.Handler, method, path string, body io.Reader) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, body)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func Test(t *testing.T) {
	r := GetMainRouter("../")
	g := goblin.Goblin(t)

	// Test the status endpoint
	g.Describe("Status Endpoint", func() {
		g.It("Should should return 200", func() {
			w := performRequest(r, "GET", "/status", nil)
			g.Assert(w.Code).Equal(http.StatusOK)
		})
	})
	// Test the hello endpoint
	g.Describe("Hello Endpoint", func() {
		g.It("Should should return 200", func() {
			w := performRequest(r, "GET", "/v1/hello", nil)
			g.Assert(w.Code).Equal(http.StatusOK)
		})
	})
	// Test the hello_ui endpoint
	g.Describe("Hello_UI Endpoint", func() {
		g.It("Should should return 200", func() {
			w := performRequest(r, "GET", "/v1/hello_ui", nil)
			g.Assert(w.Code).Equal(http.StatusOK)
		})
	})
	// Test the /v1/styles/main.css endpoint
	g.Describe("/v1/styles/main.css Endpoint", func() {
		g.It("Should should return 200", func() {
			w := performRequest(r, "GET", "/v1/styles/main.css", nil)
			g.Assert(w.Code).Equal(http.StatusOK)
		})
	})
}
