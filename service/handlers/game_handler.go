package handlers

import (
	"net/http"
	"github.com/unrolled/render"
)

func createGameHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct{ Test string }{"This is a test"})
	}
}
