package service

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/markllama/hexgame-server/hexmap"
)


func gameListHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Location", "https://mysite.com/games/")
		formatter.JSON(w, http.StatusOK, []hexmap.Map {} )
	}	
}

func gameHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		mapId := vars["id"]
		
		w.Header().Set("Location", req.URL.String())
		formatter.JSON(w, http.StatusOK, hexmap.Map { Name: mapId } )
	}	
}
