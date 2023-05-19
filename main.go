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
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func main() {

	r := mux.NewRouter()

	movies = append(
		movies,
		Movie{
			ID:    "1",
			Isbn:  "6466",
			Title: "XYXYX",
			Director: &Director{
				Firstname: "John",
				Lastname:  "Doe",
			}},
		Movie{
			ID:    "2",
			Isbn:  "776767",
			Title: "ZZZZZ",
			Director: &Director{
				Firstname: "KAA",
				Lastname:  "IOU",
			}},
	)

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server on PORT 8080")

	log.Fatal(http.ListenAndServe(":8080", r))
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]
	var index int = -1
	for i, movie := range movies {
		if id == movie.ID {
			index = i
		}
	}

	if index != -1 {
		movies = append(movies[:index], movies[index+1:]...)
		json.NewEncoder(w).Encode(movies)
	}
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]
	for _, movie := range movies {
		if id == movie.ID {
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var m Movie
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	m.ID = strconv.Itoa((rand.Intn(10000000)))
	movies = append(movies, m)
	json.NewEncoder(w).Encode(movies)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]
	for index, movie := range movies {
		if id == movie.ID {
			movies = append(movies[:index], movies[index+1:]...)
			var mo Movie
			_ = json.NewDecoder(r.Body).Decode(&mo)
			mo.ID = id
			movies = append(movies, mo)
			json.NewEncoder(w).Encode(movies)
			return
		}
	}

}
