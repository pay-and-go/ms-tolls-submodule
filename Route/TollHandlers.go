package Route

import (
	"goproj/BusinessLogic"
	"github.com/gorilla/mux"
)

func addTollHandlers(router *mux.Router) {
	// ***** 'Tolls' Routes *****
	router.HandleFunc("/addToll", BusinessLogic.CreateTollEndpoint).Methods("POST")
	router.HandleFunc("/getAllTolls", BusinessLogic.GetAllTollsEndpoint).Methods("GET")
	router.HandleFunc("/getToll/{id}", BusinessLogic.GetTollEndpoint).Methods("GET")
	router.HandleFunc("/deleteToll/{id}", BusinessLogic.DeleteTollEndpoint).Methods("DELETE")
	router.HandleFunc("/updateToll/{id}", BusinessLogic.UpdateTollEndpoint).Methods("PUT")
}
