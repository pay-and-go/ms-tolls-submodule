package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

var client *mongo.Client

// ***** Structures *****
type Toll struct {
	TollId           int     `json:"tollId" bson:"tollId"`
	Administrator    string  `json:"administrator" bson:"administrator"`
	CoorLat          float64 `json:"coor_lat" bson:"coor_lat"`
	CoorLng          float64 `json:"coor_lng" bson:"coor_lng"`
	CranePhoneNumber string  `json:"crane_phone_number" bson:"crane_phone_number"`
	Name             string  `json:"name" bson:"name"`
	Price            float64 `json:"price" bson:"price"`
	Sector           string  `json:"sector" bson:"sector"`
	Territory        string  `json:"territory" bson:"territory"`
	TollPhoneNumber  string  `json:"toll_phone_number" bson:"toll_phone_number"`
}
type TollsInRoute struct {
	Route int     `json:"route" bson:"route"`
	Tolls []int   `json:"tolls" bson:"tolls"`
}
// ***** End Structures *****

// ***** Tolls CRUD *****
func CreateTollEndpoint(response http.ResponseWriter, request *http.Request) {
    response.Header().Set("content-type", "application/json")
    var toll Toll
    _ = json.NewDecoder(request.Body).Decode(&toll)
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
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := strconv.Atoi(params["id"])
	var toll Toll
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
	var toll Toll
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

// ***** TollsInRoutes CRUD *****
func CreateTollsInRouteEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var tollsInRoute TollsInRoute
	_ = json.NewDecoder(request.Body).Decode(&tollsInRoute)
	collection := client.Database("TollSubmodule").Collection("TollsInRoutes")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, _ := collection.InsertOne(ctx, tollsInRoute)
	json.NewEncoder(response).Encode(result)
}
func GetAllTollsInRoutesEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var tollsInRoutes []TollsInRoute
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
		var tollsInRoute TollsInRoute
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
	id, _ := strconv.Atoi(params["id"])
	var tollsInRoute TollsInRoute
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
	id, _ := strconv.Atoi(params["id"])
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
	id, _ := strconv.Atoi(params["id"])
	var tollsInRoute TollsInRoute
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

func main() {
	loadEnv()
	fmt.Println("Starting the application...")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	user := os.Getenv("USER")
	pass := os.Getenv("PASSWORD")
	client, _ = mongo.Connect(ctx, options.Client().ApplyURI(
        "mongodb+srv://"+user+":"+pass+"@cluster0.blmmt.mongodb.net/TollSubmodule?retryWrites=true&w=majority",
     ))

	//client, _ = mongo.Connect(ctx, options.Client().ApplyURI(
	//	"mongodb+srv://administrator:administratorPayAndGo123@cluster0.blmmt.mongodb.net/TollSubmodule?retryWrites=true&w=majority",
	// ))
    //if err != nil { log.Fatal(err) }

	router := mux.NewRouter()
	// ***** 'Tolls' Routes *****
	router.HandleFunc("/addToll", CreateTollEndpoint).Methods("POST")
	router.HandleFunc("/getAllTolls", GetAllTollsEndpoint).Methods("GET")
	router.HandleFunc("/getToll/{id}", GetTollEndpoint).Methods("GET")
	router.HandleFunc("/deleteToll/{id}", DeleteTollEndpoint).Methods("DELETE")
	router.HandleFunc("/updateToll/{id}", UpdateTollEndpoint).Methods("PUT")

	// ***** 'TollsInRoutes' Routes *****
	router.HandleFunc("/addTollsInRoute", CreateTollsInRouteEndpoint).Methods("POST")
	router.HandleFunc("/getAllTollsInRoute", GetAllTollsInRoutesEndpoint).Methods("GET")
	router.HandleFunc("/getTollInARoute/{id}", GetTollsInARouteEndpoint).Methods("GET")
	router.HandleFunc("/deleteTollsInRoute/{id}", DeleteTollsInRouteEndpoint).Methods("DELETE")
	router.HandleFunc("/updateTollsInRoute/{id}", UpdateTollsInRouteEndpoint).Methods("PUT")

	http.ListenAndServe(":80", router)
}

func loadEnv(){
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
	}
}