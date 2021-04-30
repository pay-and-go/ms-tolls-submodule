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
	"strconv"
	"time"
)

// ***** Tolls CRUD *****
func CreateTollEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var toll Model.Toll
	_ = json.NewDecoder(request.Body).Decode(&toll)
	collection := client.Database("TollSubmodule").Collection("Tolls")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, _ := collection.InsertOne(ctx, toll)
	json.NewEncoder(response).Encode(result)
}
func GetAllTollsEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var tolls []Model.Toll
	collection := client.Database("TollSubmodule").Collection("Tolls")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var toll Model.Toll
		cursor.Decode(&toll)
		tolls = append(tolls, toll)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(tolls)
}
func GetTollEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := strconv.Atoi(params["id"])
	var toll Model.Toll
	collection := client.Database("TollSubmodule").Collection("Tolls")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err := collection.FindOne(ctx, bson.M{"tollId": id}).Decode(&toll)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(toll)
}
func DeleteTollEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	collection := client.Database("TollSubmodule").Collection("Tolls")
	id, _ := strconv.Atoi(params["id"])
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, err := collection.DeleteOne(ctx, bson.M{"tollId": id})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(result)
}
func UpdateTollEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := strconv.Atoi(params["id"])
	var toll Model.Toll
	_ = json.NewDecoder(request.Body).Decode(&toll)
	collection := client.Database("TollSubmodule").Collection("Tolls")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	//result, _ := collection.InsertOne(ctx, toll)
	filter := bson.M{"tollId": bson.M{"$eq": id}}

	update := bson.M{
		"$set": bson.M{
			"administrator": toll.Administrator,
			"coor_lat": toll.CoorLat,
			"coor_lng": toll.CoorLng,
			"crane_phone_number": toll.CranePhoneNumber,
			"name": toll.Name,
			"price": toll.Price,
			"sector": toll.Sector,
			"territory": toll.Territory,
			"toll_phone_number": toll.TollPhoneNumber,
		},
	}
	result, err := collection.UpdateOne(ctx, filter, update)

	if err != nil {
		fmt.Println("UpdateOne() result ERROR:", err)
		os.Exit(1)
	}

	json.NewEncoder(response).Encode(result)
}
// ***** End Tolls CRUD *****

