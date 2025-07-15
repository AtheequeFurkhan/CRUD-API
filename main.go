package crudapi

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct{
	ID string`json:"id"`
	Isbn string`json:"isbn"`
	Title string`json:"title"`
	Director *Director `json:"director"`

}
type Director struct{
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
}

var movies []Movie

func getMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(movies)

}
 
func deleteMovie(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	for index, item:= range movies{
		if item.ID == params["id"]{
			movies = append(movies[:index], movies[index+1:]...)
			break
	}
	
}
}



func createMovie(){

}

func updateMovie(){

}

func getMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r) 
	for _, item:= range movies{
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
		}
	} 
}

func main(){
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "4311", Title: "F1", Director: &Director{FirstName: "John", LastName: "Cage"}})
	movies = append(movies, Movie{ID: "2", Isbn: "8722", Title: "Mad Mux", Director: &Director{FirstName: "Peter", LastName: "Kevin"}})
	movies = append(movies, Movie{ID: "3", Isbn: "7613", Title: "End Game", Director: &Director{FirstName: "Kevin", LastName: "Fiege"}})
	movies = append(movies, Movie{ID: "4", Isbn: "7282", Title: "Thor Ragnorak", Director: &Director{FirstName: "Alex", LastName: "Hales"}})
	movies = append(movies, Movie{ID: "5", Isbn: "6529", Title: "The Revenant", Director: &Director{FirstName: "Alejandro", LastName: "González"}})
	movies = append(movies, Movie{ID: "6", Isbn: "9894", Title: "Interstellar", Director: &Director{FirstName: "Christopher", LastName: "Nolan"}})


	r.HandleFunc("/movies",getMovies).Methods("GET")
	r.HandleFunc("movies/{id}",getMovie).Methods("GET")
	r.HandleFunc("/movies",createMovie).Methods("POST")
	r.HandleFunc("movies/{id}",updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}",deleteMovie).Methods("DELETE")

	fmt.Printf("Starting Server at 8080!")
	log.Fatal(http.ListenAndServe(":8080",r))
}

