package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"mdb/model"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://debankanmitra:pass@cluster0.ducoyoi.mongodb.net/?retryWrites=true&w=majority"
const dbname string = "Netflix"
const Collectionname = "Watchlist"

// MOst Important
var collection *mongo.Collection

// connect with mongodb
func init() { // this func runs for the very first time

	//client option , responsible for getting a connection for client
	clientOption := options.Client().ApplyURI(connectionString) // responsible for building a connection , client option says this is the uri i wil be using

	// connect to mongodb , fires up a connection request
	client, _ := mongo.Connect(context.TODO(), clientOption)
	fmt.Println("mongodb connection success")

	// creating collection instance
	collection = client.Database(dbname).Collection(Collectionname) // collection is inside the database
	fmt.Println("collection instance is ready")

}

// MONGODB helpers - file
// insert 1 record
func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("inserted 1 movie in db with id ", inserted.InsertedID)
}

// update 1 record
func updaterecord(movieID string) {
	id, _ := primitive.ObjectIDFromHex(movieID)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}
	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Documents updated: %v\n", result.ModifiedCount)
}

func deleterecord(movieID string) {
	id, _ := primitive.ObjectIDFromHex(movieID)
	filter := bson.M{"_id": id}
	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Documents deleted: %v\n", result.DeletedCount)
}

func getAllmovies() []primitive.M {
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var movies []primitive.M
	for cursor.Next(context.TODO()) {
		var movie bson.M
		if err = cursor.Decode(&movie); err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	defer cursor.Close(context.TODO())
	return movies
}

// Actual controllers --file
func GetMymovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/x-www-form-urlencode")
	allmovies := getAllmovies()

	// throwing json values
	json.NewEncoder(w).Encode(allmovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie model.Netflix
	json.NewDecoder(r.Body).Decode(&movie)

	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func MarkasWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	params := mux.Vars(r)
	updaterecord(params["id"])
}

func Deletemovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	params := mux.Vars(r)
	deleterecord(params["id"])
}
