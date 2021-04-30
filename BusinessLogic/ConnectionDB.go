package BusinessLogic

import (
	"context"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

var client *mongo.Client

func CreateConnection(ctx context.Context) {
	loadEnv()
	user := os.Getenv("USER")
	pass := os.Getenv("PASSWORD")
	client, _ = mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb+srv://"+user+":"+pass+"@cluster0.blmmt.mongodb.net/TollSubmodule?retryWrites=true&w=majority",
	))
}

func loadEnv(){
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
	}
}