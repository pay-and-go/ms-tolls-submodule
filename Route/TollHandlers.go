package Route

import (
	"goproj/BusinessLogic"
	"github.com/gorilla/mux"
)

func addTollHandlers(router *mux.Router) {
	// ***** 'Tolls' Routes *****
	router.HandleFunc("/tolls-ms/addToll", BusinessLogic.CreateTollEndpoint).Methods("POST")
	router.HandleFunc("/tolls-ms/getAllTolls", BusinessLogic.GetAllTollsEndpoint).Methods("GET")
	router.HandleFunc("/tolls-ms/getToll/{id}", BusinessLogic.GetTollEndpoint).Methods("GET")
	router.HandleFunc("/tolls-ms/deleteToll/{id}", BusinessLogic.DeleteTollEndpoint).Methods("DELETE")
	router.HandleFunc("/tolls-ms/updateToll/{id}", BusinessLogic.UpdateTollEndpoint).Methods("PUT")
}
