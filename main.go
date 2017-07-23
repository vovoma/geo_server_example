package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Geolocation struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	UserID    string  `json:"userId"`
	Timestamp int     `json:"timestamp"`
}

var DATA_FILES_DIR string
var locations []Geolocation

func serve(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(locations)
	}

	if r.Method == http.MethodPost {
		var location Geolocation
		err := json.NewDecoder(r.Body).Decode(&location)
		if err != nil {
			log.Fatal(err)
			return
		}
		filename := fmt.Sprintf("user%s_t%d", location.UserID, location.Timestamp)
		file, err := os.Create(DATA_FILES_DIR + filename)
		if err != nil {
			log.Fatal(err)
		}
		data, _ := json.Marshal(location)
		file.Write(data)
		log.Printf("File create %s\n", filename)

		locations = append(locations, location)
		json.NewEncoder(w).Encode(location)
		defer r.Body.Close()
	}
	return
}
func init() {
	DATA_FILES_DIR = os.Getenv("GEOLOCATION_DATA_FILES_DIR")
	if DATA_FILES_DIR == "" {
		DATA_FILES_DIR = "./app/data/"
	}
}
func main() {
	http.HandleFunc("/geolocation", serve)
	log.Fatal(http.ListenAndServe(":4000", nil))
}
