package BusinessLogic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"goproj/Model"
	"net/http"
	"os"
	"time"
)

// ***** TollsInRoutes CRUD *****
func CreateTollsInRouteEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var tollsInRoute Model.TollsInRoute
	_ = json.NewDecoder(request.Body).Decode(&tollsInRoute)
	collection := client.Database("TollSubmodule").Collection("TollsInRoutes")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, _ := collection.InsertOne(ctx, tollsInRoute)
	json.NewEncoder(response).Encode(result)
}
func GetAllTollsInRoutesEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var tollsInRoutes []Model.TollsInRoute
	collection := client.Database("TollSubmodule").Collection("TollsInRoutes")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var tollsInRoute Model.TollsInRoute
		cursor.Decode(&tollsInRoute)
		tollsInRoutes = append(tollsInRoutes, tollsInRoute)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(tollsInRoutes)
}
func GetTollsInARouteEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	//id, _ := strconv.Atoi(params["id"])
	id, _ := params["id"]
	var tollsInRoute Model.TollsInRoute
	collection := client.Database("TollSubmodule").Collection("TollsInRoutes")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err := collection.FindOne(ctx, bson.M{"route": id}).Decode(&tollsInRoute)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(tollsInRoute)
}
func DeleteTollsInRouteEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	collection := client.Database("TollSubmodule").Collection("TollsInRoutes")
	//id, _ := strconv.Atoi(params["id"])
	id, _ := params["id"]
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, err := collection.DeleteOne(ctx, bson.M{"route": id})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(result)
}
func UpdateTollsInRouteEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	//id, _ := strconv.Atoi(params["id"])
	id, _ := params["id"]
	var tollsInRoute Model.TollsInRoute
	_ = json.NewDecoder(request.Body).Decode(&tollsInRoute)
	collection := client.Database("TollSubmodule").Collection("TollsInRoutes")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	//result, _ := collection.InsertOne(ctx, toll)
	filter := bson.M{"route": bson.M{"$eq": id}}

	update := bson.M{
		"$set": bson.M{
			"tolls": tollsInRoute.Tolls,
		},
	}
	result, err := collection.UpdateOne(ctx, filter, update)

	if err != nil {
		fmt.Println("UpdateOne() result ERROR:", err)
		os.Exit(1)
	}

	json.NewEncoder(response).Encode(result)
}
// ***** End TollsInRoutes CRUD *****
