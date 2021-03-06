//
// The http responder
//
package service

import (
	"net/http"

	"github.com/urfave/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

func NewServer() *negroni.Negroni {
	formatter := render.New(render.Options{
		IndentJSON: true,
	})

	n := negroni.Classic()
	mux := mux.NewRouter()

	initRoutes(mux, formatter)

	n.UseHandler(mux)

	return n
}

func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/test", testHandler(formatter)).Methods("GET")
	mx.HandleFunc("/games/", gameListHandler(formatter)).Methods("GET")
	mx.HandleFunc("/games/{id}", gameHandler(formatter)).Methods("GET")
	mx.HandleFunc("/maps/", mapListHandler(formatter)).Methods("GET")
	mx.HandleFunc("/maps/{id}", mapHandler(formatter)).Methods("GET")
	mx.HandleFunc("/matches/", matchListHandler(formatter)).Methods("GET")
	mx.HandleFunc("/matches/{id}", matchHandler(formatter)).Methods("GET")
	mx.HandleFunc("/users/", userListHandler(formatter)).Methods("GET")
	mx.HandleFunc("/users/{id}", userHandler(formatter)).Methods("GET")
	
}

func testHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct{ Test string }{"This is a test"})
	}
}
