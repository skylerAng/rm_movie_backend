package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Movie struct {
	// Map native go structure
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title        string             `json:"title,omitempty" bson:"title,omitempty"`
	Description  string             `json:"description,omitempty" bson:"description,omitempty"`
	Filename     string             `json:"filename,omitempty" bson:"filename,omitempty"`
	Originallink string             `json:"originallink,omitempty" bson:"originallink,omitempty"`
}

var client *mongo.Client

// Create post endpoint
func CreateMovieEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")

	var movie Movie
	_ = json.NewDecoder(request.Body).Decode(&movie)

	collection := client.Database("movie_db").Collection("movie")

	// add timeout
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, _ := collection.InsertOne(ctx, movie)
	json.NewEncoder(response).Encode(result)

}

func GetMoviesEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	// slice of movie
	var movies []Movie
	collection := client.Database("movie_db").Collection("movie")
	// define context
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// return everything in the collection
	cursor, err := collection.Find(ctx, bson.M{})
	// error handling
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		// Return JSON from header
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var movie Movie
		cursor.Decode(&movie)
		// Decode then append to movie slice
		// Cannot assign as it will be formatted to mongo cursor language
		// Loop and store each item to be returned
		movies = append(movies, movie)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(response).Encode(movies)
}

func GetMovieEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	//
	params := mux.Vars(request)
	// Passing in an ID and convert it into a usable Mongo DB Object ID
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var movie Movie
	// Set up context and opening of context
	collection := client.Database("movie_db").Collection("movie")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := collection.FindOne(ctx, Movie{ID: id}).Decode(&movie)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(response).Encode(movie)
}

func GetImage() {

}

func Startup() {
	fmt.Println("Starting the application...")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	var err error
	client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	router := mux.NewRouter()

	// link up method that can access
	router.HandleFunc("/addMovie", CreateMovieEndpoint).Methods("POST")
	router.HandleFunc("/getMovies", GetMoviesEndpoint).Methods("POST")
	router.HandleFunc("/movie/{id}", GetMovieEndpoint).Methods("POST")
	http.ListenAndServe(":12345", router)
}
