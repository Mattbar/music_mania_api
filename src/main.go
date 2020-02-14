package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)
// Song struct
type Song struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Album   string `json:"album"`
	Key 		string `json:"key"`
	URL 		string `json:"url"`
}

// type Files struct {
// 	Path	[]string `json:"path"`
// }

// Init books var as a slice Book struct
var songs []Song

func main() {
	router := mux.NewRouter()
	
	headers := handlers.AllowedHeaders([]string{"Origin", "X-Requested-With", "Content-Type", "Authorization", " X-Auth-Token", "Accept", "Accept-Language", "Content-Language", "Origin"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD"})
	origins := handlers.AllowedOrigins([]string{"*"})

	// Mock books
	songs = append(songs, Song{ID: "1", Title: "funkorama", Artist: "kevin macleod", Album: "Album 1", Key:"kevin macleod/Album 1/funkorama.mp3"})
	songs = append(songs, Song{ID: "2", Title: "React", Artist: "peter", Album: "FakeAll2"})

	router.HandleFunc("/songs", getSongs).Methods("GET")
	router.HandleFunc("/songs/{id}", getSong).Methods("GET")

	log.Println("Now server is running on port 8800")
	log.Fatal(http.ListenAndServe("8800", handlers.CORS(headers, methods, origins)(router)))
}

func getSongs(w http.ResponseWriter, r *http.Request) {
	log.Println("get all called")
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(songs)
}

func getSong(w http.ResponseWriter, r *http.Request) {
	log.Println("GET SONG CALLED")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range songs {
		if item.ID == params["id"] {
			if item.Key != "" {
				url := GetURL(item.Key)
				item.URL = url
			}

			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Song{})
}

// func addSong(w http.ResponseWriter, r *http.Request) {
// 	log.Println("Add  called")
// 	w.Header().Set("Content-Type", "application/json")
// 	var song Song
// 	err := json.NewDecoder(r.Body).Decode(&song)
// 	if err != nil {
// 		log.Fatal(err)
//     http.Error(w, err.Error(), http.StatusBadRequest)
//     return
//   }
// 	song.ID = strconv.Itoa(rand.Intn(1000))
// 	songs = append(songs, song)
// 	json.NewEncoder(w).Encode(song)
// }

// func addArist(w http.ResponseWriter, r *http.Request){
// 	log.Println("ADD a Artist called")
// 	w.Header().Set("Content-Type", "application/json")
// 	var files Files
// 	err := json.NewDecoder(r.Body).Decode(&files)
// 		if err != nil {
// 			log.Fatal(err)
//     	http.Error(w, err.Error(), http.StatusBadRequest)
//     	return
// 		}
// 	log.Println(files.Path)
// 	addToBucket(files.Path)
// 	json.NewEncoder(w).Encode(files)

// }

// func updateSong(w http.ResponseWriter, r *http.Request) {
// 	log.Println("update book called")
// }
// func removeSong(w http.ResponseWriter, r *http.Request) {
// 	log.Println("remove book called")
// }
