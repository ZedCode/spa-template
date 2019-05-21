package spaserver

import (
	"net/http"

	"github.com/unrolled/render"
)

// This is the default endpoint, It returns the UI of the SPA
func serveIndex(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		// We return everything we need to the template here.
		formatter.HTML(w, http.StatusOK, "index", indexContent)
	}
}

// This is an example API endpoint that returns a JSON response
func serveExampleEndpoint(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		r := ExampleContent{
			ExampleTitle:       "This is an example title being served over the API.",
			ExampleDescription: "This is an example description being served over the API.",
		}
		formatter.JSON(w, http.StatusOK, r)
	}
}
