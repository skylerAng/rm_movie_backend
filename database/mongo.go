package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/joho/godotenv"
)

func ConnectDB() *mongo.Client {
	ctx := context.Background()
	// Load env file
	err := godotenv.Load("env/api.env")
	if err != nil {
		log.Fatal("Error Loading Env File")
	}

	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	db_name := os.Getenv("DB_NAME")

	uri := fmt.Sprintf(
		"mongodb://%s:%s@%s/%s",
		username, password, host, db_name,
	)

	log.Println(uri)

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("Connection Failed to Database")
		panic(err)
	}

	if err := client.Connect(ctx); err != nil {
		panic(err)
	}

	if err := client.Ping(context.Background(), readpref.Primary()); err != nil {
		log.Fatal("Connection Failed to Database")
		panic(err)
	}

	color.Green("Connected to Database at:\n" + uri)

	return client
}

// Client instance
var DB *mongo.Client = ConnectDB()

// getting database collections
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database(os.Getenv("DB_NAME")).Collection(collectionName)
	return collection
}
