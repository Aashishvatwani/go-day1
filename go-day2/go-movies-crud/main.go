package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type Movie struct {
	ID       string    `json:"id"`
	ISBN     string    `json:"isbn"`
	Name     string    `json:"name"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	for i, movie := range movies {
		if movie.ID == id {
			json.NewEncoder(w).Encode(movies[i])
			return
		}
	}
	json.NewEncoder(w).Encode(&Movie{})
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(1000000)) // Mock ID - not safe
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie) // return the created movie, not all movies
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["id"]

	for index, movie := range movies {
		if movie.ID == id {
			movies = append(movies[:index], movies[index+1:]...)
			json.NewEncoder(w).Encode(movies)
			return
		}
	}

	json.NewEncoder(w).Encode(movies)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["id"]
	for i, k := range movies {
		if k.ID == id {
			movies = append(movies[:i], movies[i+1:]...) // remove the movie
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			// return the updated movie, not all moviese
			movie.ID = id
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

func main() {
	r := mux.NewRouter()
	movies = append(movies, Movie{ID: "1", ISBN: "1234567890", Name: "Inception", Director: &Director{FirstName: "Christopher", LastName: "Nolan"}})
	movies = append(movies, Movie{ID: "2", ISBN: "0987654321", Name: "The Matrix", Director: &Director{FirstName: "Lana", LastName: "Wachowski"}})
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
