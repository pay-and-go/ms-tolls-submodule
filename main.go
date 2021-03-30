package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"time"
)

var client *mongo.Client

type Coordinate struct{
    lat float64			`json:"lat" bson:"lat"`
    lng float64			`json:"lng" bson:"lng"`
}

type Toll struct {
	TollId           int        `json:"tollId" bson:"tollId"`
	Administrator    string     `json:"administrator" bson:"administrator"`
	Coordinate `json:"coordinates" bson:"coordinates"`
	CranePhoneNumber string     `json:"crane_phone_number" bson:"crane_phone_number"`
	Name             string     `json:"name" bson:"name"`
	Price            float64    `json:"price" bson:"price"`
	Sector           string     `json:"sector" bson:"sector"`
	Territory        string     `json:"territory" bson:"territory"`
	TollPhoneNumber  string     `json:"toll_phone_number" bson:"toll_phone_number"`
}

func CreateTollEndpoint(response http.ResponseWriter, request *http.Request) {
    response.Header().Set("content-type", "application/json")
    var toll Toll
    _ = json.NewDecoder(request.Body).Decode(&toll)
    
	fmt.Printf("%s\n", toll)

    collection := client.Database("TollSubmodule").Collection("Tolls")
    ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
    result, _ := collection.InsertOne(ctx, toll)
    json.NewEncoder(response).Encode(result)
}
func GetAllTollsEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var tolls []Toll
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
		var toll Toll
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
	
}

func main() {
	fmt.Println("Starting the application...")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, _ = mongo.Connect(ctx, options.Client().ApplyURI(
         "mongodb+srv://administrator:administratorPayAndGo123@cluster0.blmmt.mongodb.net/TollSubmodule?retryWrites=true&w=majority",
      ))
    //if err != nil { log.Fatal(err) }

	router := mux.NewRouter()
	router.HandleFunc("/addToll", CreateTollEndpoint).Methods("POST")
	router.HandleFunc("/getAllTolls", GetAllTollsEndpoint).Methods("GET")
	router.HandleFunc("/getToll/{id}", GetTollEndpoint).Methods("GET")
	http.ListenAndServe(":12345", router)
}