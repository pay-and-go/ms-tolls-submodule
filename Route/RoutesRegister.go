package Route

import (
	"github.com/gorilla/mux"
	"net/http"
)

func RegisterRoutes() {
	router := mux.NewRouter()
	addTollHandlers(router)
	addTollsInRouteHandlers(router)
	http.ListenAndServe(":80", router)
}

