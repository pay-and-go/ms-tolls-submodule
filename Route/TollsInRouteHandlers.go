package Route

import (
	"github.com/gorilla/mux"
	"goproj/BusinessLogic"
)

func addTollsInRouteHandlers(router *mux.Router) {
	// ***** 'TollsInRoutes' Routes *****
	router.HandleFunc("/addTollsInRoute", BusinessLogic.CreateTollsInRouteEndpoint).Methods("POST")
	router.HandleFunc("/getAllTollsInRoute", BusinessLogic.GetAllTollsInRoutesEndpoint).Methods("GET")
	router.HandleFunc("/getTollInARoute/{id}", BusinessLogic.GetTollsInARouteEndpoint).Methods("GET")
	router.HandleFunc("/deleteTollsInRoute/{id}", BusinessLogic.DeleteTollsInRouteEndpoint).Methods("DELETE")
	router.HandleFunc("/updateTollsInRoute/{id}", BusinessLogic.UpdateTollsInRouteEndpoint).Methods("PUT")
}
