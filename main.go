package main

import (
	"api/database"
	"api/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var getM []models.Movie
	database.DB.Preload("Movie").Find(&getM)
	database.DB.Model(&models.Movie{})
	json.NewEncoder(w).Encode(getM)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var movies []models.Movie
	// for index, item := range movies {
	// 	if item.ID == params["id"] {
	// 		movies = append(movies[:index], movies[index+1:]...)
	// 		break
	// 	}
	// }
	movie := models.Movie{
		ID: params["id"],
	}
	database.DB.Delete(&movie)

	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	// for _, item := range movies {
	// 	if item.ID == params["id"] {
	// 		json.NewEncoder(w).Encode(item)
	// 		return
	// 	}
	// }
	var getM models.Movie
	database.DB.Where("id=?", params["id"]).Preload("Movie").First(&getM)
	json.NewEncoder(w).Encode(getM)

}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie models.Movie

	_ = json.NewDecoder(r.Body).Decode(&movie)

	// movie.ID = strconv.Itoa(rand.Intn(100000000))

	database.DB.Create(&movie)

	// movies = append(movies, movie)

	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	// for index, item := range movies {
	// 	if item.ID == params["id"] {
	// 		movies = append(movies[:index], movies[index+1:]...)
	// 		var movie models.Movie
	// 		_ = json.NewDecoder(r.Body).Decode(&movie)
	// 		movie.ID = strconv.Itoa(rand.Intn(100000000))
	// 		movies = append(movies, movie)
	// 		json.NewEncoder(w).Encode(movie)
	// 	}
	// }

	movie := models.Movie{
		ID: params["id"],
	}

	_ = json.NewDecoder(r.Body).Decode(&movie)

	database.DB.Model(&movie).Updates(movie)

	json.NewEncoder(w).Encode(movie)
}

func main() {

	database.Connect()

	r := mux.NewRouter()

	// movies = append(movies, Movie{ID: "1", Isbn: "438227", Title: "Movie One", Director: &Director{FirstName: "John", LastName: "Doe"}})
	// movies = append(movies, Movie{ID: "2", Isbn: "45455", Title: "Movie Two", Director: &Director{FirstName: "Steve", LastName: "Smith"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("Port")

	fmt.Printf("Starting at port 3000\n")
	log.Fatal(http.ListenAndServe(":"+port, r))
}
