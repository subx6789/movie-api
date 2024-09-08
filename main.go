package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

// Director struct holds director's first and last name.
type Director struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

// Movie struct represents a movie with fields like Id, ISBN, title, overview, and director.
type Movie struct {
	Id       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Overview string    `json:"overview"`
	Director *Director `json:"director"`
}

var movies []Movie // Global variable to hold the list of movies

// init function loads environment variables before main() is executed
func init() {
	// Loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("Warning: No .env file found")
	}
}

func main() {
	// Set some sample movies to the movies slice
	// Movie 1: Adding movie details including title, overview, and director information
	movies = append(movies, Movie{
		Id:       generateId(),
		Isbn:     "531306",
		Title:    "Rim of the World",
		Overview: "Stranded at a summer camp when aliens attack the planet, four teens with nothing in common embark on a perilous mission to save the world.",
		Director: &Director{
			FirstName: "Joseph",
			LastName:  "McGinty Nichol",
		},
	})
	// Movie 2
	movies = append(movies, Movie{
		Id:       generateId(),
		Isbn:     "181808",
		Title:    "Star Wars: The Last Jedi",
		Overview: "Rey develops her newly discovered abilities with the guidance of Luke Skywalker, who is unsettled by the strength of her powers. Meanwhile, the Resistance prepares to do battle with the First Order.",
		Director: &Director{
			FirstName: "Rian",
			LastName:  "Johnson",
		},
	})
	// Movie 3
	movies = append(movies, Movie{
		Id:       generateId(),
		Isbn:     "401650",
		Title:    "DC Super Hero Girls: Hero of the Year",
		Overview: "Wonder Woman, Supergirl, Batgirl, Harley Quinn, Bumblebee, Poison Ivy and Katana band together to navigate the twists and turns of high school in DC Super Hero Girls: Hero of the Year.",
		Director: &Director{
			FirstName: "Cecilia",
			LastName:  "Aranovich",
		},
	})
	// Movie 4
	movies = append(movies, Movie{
		Id:       generateId(),
		Isbn:     "49026",
		Title:    "The Dark Knight Rises",
		Overview: "Following the death of District Attorney Harvey Dent, Batman assumes responsibility for Dent's crimes to protect the late attorney's reputation and is subsequently hunted by the Gotham City Police Department. Eight years later, Batman encounters the mysterious Selina Kyle and the villainous Bane, a new terrorist leader who overwhelms Gotham's finest. The Dark Knight resurfaces to protect a city that has branded him an enemy.",
		Director: &Director{
			FirstName: "Christopher",
			LastName:  "Nolan",
		},
	})
	// Movie 5
	movies = append(movies, Movie{
		Id:       generateId(),
		Isbn:     "271110",
		Title:    "Captain America: Civil War",
		Overview: "Following the events of Age of Ultron, the collective governments of the world pass an act designed to regulate all superhuman activity. This polarizes opinion amongst the Avengers, causing two factions to side with Iron Man or Captain America, which causes an epic battle between former allies.",
		Director: &Director{
			FirstName: "Anthony",
			LastName:  "Russo",
		},
	})
	// Get the port from environment variables or use the default port 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r := mux.NewRouter()
	// Define routes with proper error handling in all HTTP handlers
	r.HandleFunc("/movies", getMovies).Methods(http.MethodGet)
	r.HandleFunc("/movie/{id}", getMovie).Methods(http.MethodGet)
	r.HandleFunc("/movies", createMovie).Methods(http.MethodPost)
	r.HandleFunc("/movie/{id}", updateMovie).Methods(http.MethodPut)
	r.HandleFunc("/movie/{id}", deleteMovie).Methods(http.MethodDelete)
	// Start the server
	fmt.Println("Server starting on port", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

// generateId function creates a new unique ID using the uuid package
func generateId() string {
	return uuid.New().String()
}

// findMovieById finds a movie by ID, returning an error if not found
func findMovieById(id string) (*Movie, error) {
	for _, item := range movies {
		if item.Id == id {
			return &item, nil
		}
	}
	return nil, errors.New("movie not found")
}

// validateMovie checks if the required fields are present
func validateMovie(movie Movie) error {
	if movie.Isbn == "" || movie.Title == "" || movie.Overview == "" {
		return errors.New("isbn, title, and overview are required fields")
	}
	if movie.Director == nil || movie.Director.FirstName == "" || movie.Director.LastName == "" {
		return errors.New("director's first and last name are required")
	}
	return nil
}

// getMovies function returns all movies in JSON format
func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(movies); err != nil {
		http.Error(w, "Failed to encode movies", http.StatusInternalServerError)
	}
}

// getMovie function returns a movie by ID or a 404 error if not found
func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	movieId := params["id"]
	// Find the movie by ID
	movie, err := findMovieById(movieId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	// Encode and return the movie
	if err := json.NewEncoder(w).Encode(movie); err != nil {
		http.Error(w, "Failed to encode movie", http.StatusInternalServerError)
	}
}

// createMovie function validates input and creates a new movie
func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	// Decode request body into the movie struct
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		http.Error(w, "Invalid input, please provide a valid movie object", http.StatusBadRequest)
		return
	}
	// Validate movie fields
	if err := validateMovie(movie); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	movie.Id = generateId()        // Assign a new unique ID to the movie
	movies = append(movies, movie) // Add the movie to the list
	// Return the newly created movie
	if err := json.NewEncoder(w).Encode(movie); err != nil {
		http.Error(w, "Failed to encode movie", http.StatusInternalServerError)
	}
}

// updateMovie function allows partial updates on a movie
func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	movieId := params["id"]
	// Find the movie to update
	movie, err := findMovieById(movieId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	// Decode updated fields from the request body
	var updatedFields map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&updatedFields); err != nil {
		http.Error(w, "Invalid input, unable to parse JSON", http.StatusBadRequest)
		return
	}
	// Apply updates to the movie fields
	if isbn, ok := updatedFields["isbn"].(string); ok && isbn != "" {
		movie.Isbn = isbn
	}
	if title, ok := updatedFields["title"].(string); ok && title != "" {
		movie.Title = title
	}
	if overview, ok := updatedFields["overview"].(string); ok && overview != "" {
		movie.Overview = overview
	}
	if directorMap, ok := updatedFields["director"].(map[string]interface{}); ok {
		if firstName, ok := directorMap["firstName"].(string); ok && firstName != "" {
			movie.Director.FirstName = firstName
		}
		if lastName, ok := directorMap["lastName"].(string); ok && lastName != "" {
			movie.Director.LastName = lastName
		}
	}
	// Return the updated movie
	if err := json.NewEncoder(w).Encode(movie); err != nil {
		http.Error(w, "Failed to encode updated movie", http.StatusInternalServerError)
	}
}

// deleteMovie function removes a movie by its ID with proper error handling
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	movieId := params["id"]
	// Find the movie to delete
	for index, item := range movies {
		if item.Id == movieId {
			// Remove the movie from the slice
			movies = append(movies[:index], movies[index+1:]...)
			w.WriteHeader(http.StatusNoContent) // Return 204 No Content on successful deletion
			return
		}
	}
	// If movie not found, return 404 error
	http.Error(w, "Movie not found", http.StatusNotFound)
}
