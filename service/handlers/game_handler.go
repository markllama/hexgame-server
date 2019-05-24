package handlers

import (
	"net/http"
	"github.com/unrolled/render"
)

func createGameHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		Gw.Header().Add("Location", "some value")
		formatter.JSON(w, http.StatusCreated, struct{ Test string }{"This is a test"})
	}
}


