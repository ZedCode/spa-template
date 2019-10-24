package spaserver

import (
	"mime"
	"net/http"
	"os"
	"path/filepath"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/unrolled/render"
)

var (
	formatter = render.New(render.Options{
		IndentJSON: true,
	})
	// This is what we will pass into the index route.
	indexContent = IndexContent{
		IndexTitle: "EXAMPLE CONTENT",
	}
	c = cors.New(cors.Options{
		//Debug: true, // Uncomment to debug
	})
)

// NewServer is consumed by the main application
func NewServer() *negroni.Negroni {
	n := negroni.Classic()
	n.Use(c)
	mx := mux.NewRouter()
	initRoutes(mx, formatter)
	n.UseHandler(mx)
	return n
}

// This sets all of the routes we'll end up using.
func initRoutes(mx *mux.Router, formatter *render.Render) {
	// Endpoints that serve HTML will not have api in their route
	mx.HandleFunc("/", serveIndex(formatter)).Methods("GET")
	// Under /static we are shipping a /js and /css
	mx.HandleFunc("/static/{path}/{filename}", staticServer)
	// Endpoints that are API responses will all be under the api route
	mx.HandleFunc("/api/v1/example", serveExampleEndpoint(formatter)).Methods("GET")
	mx.HandleFunc("/api/v1/examplepost", serveExamplePost(formatter)).Methods("POST")
}

func staticServer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	serverPath, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	file := filepath.Join(serverPath, "static", vars["path"], vars["filename"])
	w.Header().Set("Content-Type", mime.TypeByExtension(filepath.Ext(file)))
	http.ServeFile(w, r, file)
}
