package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//need to define out data types. We want to work with movies. They'll need an id, Isbn, title & a director. Let's make the director a separate struct and point to it from movies.
type Movie struct {
	ID       string    `json: "id"`
	Isbn     string    `json: "isbn"`
	Title    string    `json: "title"`
	Director *Director `json: "director"`
}

type Director struct {
	Firstname string `json: "first name"`
	Lastname  string `json: "lastname"`
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", CreateMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Print("Starting server at port 8000 \n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
