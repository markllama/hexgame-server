package service

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/markllama/hexgame-server/hexmap"
)


func userListHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, []hexmap.Map {} )
	}	
}

func userHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		mapId := vars["id"]
		
		formatter.JSON(w, http.StatusOK, hexmap.Map { Name: mapId } )
	}	
}
