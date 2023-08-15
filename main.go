package main

import (
	"fmt"
	"log"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
)

type Movie struct {
	ID string `json:"Id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(movies)
}

func main(){
	r := mux.NewRouter()

	movies=append(movies,Movie{ID:"1",Isbn:"3344",Title:"Movie one",Director:&Director{Firstname:"John",Lastname:"Doe"}})
	movies=append(movies,Movie{ID:"2",Isbn:"3348",Title:"Movie Two",Director:&Director{Firstname:"Steve",Lastname:"Smith"}})

	r.HandleFunc("/movies",getMovies).Methods("GET")
	// r.HandleFunc("/movie/{id}",getMovie).Methods("GET")
	// r.HandleFunc("/movie",createMovie).Methods("POST")
	// r.HandleFunc("/movie/{id}",updateMovie).Methods("PUT")
	// r.HandleFunc("/movie/{id}",deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000",r))
}