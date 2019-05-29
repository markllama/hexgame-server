package service

import (
	"net/http"
	"github.com/unrolled/render"
	"github.com/markllama/hexgame-server/hexmap"
)


func mapListHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, []hexmap.Map {} )
	}	
}
