package Route

import (
	"github.com/gorilla/mux"
	"goproj/BusinessLogic"
)

func addTollsInRouteHandlers(router *mux.Router) {
	// ***** 'TollsInRoutes' Routes *****
	router.HandleFunc("/tolls-ms/addTollsInRoute", BusinessLogic.CreateTollsInRouteEndpoint).Methods("POST")
	router.HandleFunc("/tolls-ms/getAllTollsInRoute", BusinessLogic.GetAllTollsInRoutesEndpoint).Methods("GET")
	router.HandleFunc("/tolls-ms/getTollInARoute/{id}", BusinessLogic.GetTollsInARouteEndpoint).Methods("GET")
	router.HandleFunc("/tolls-ms/deleteTollsInRoute/{id}", BusinessLogic.DeleteTollsInRouteEndpoint).Methods("DELETE")
	router.HandleFunc("/tolls-ms/updateTollsInRoute/{id}", BusinessLogic.UpdateTollsInRouteEndpoint).Methods("PUT")
}
