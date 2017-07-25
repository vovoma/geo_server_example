package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type geolocation struct {
	Latitude  float64 `json:"latitude,omitempty"`
	Longitude float64 `json:"longitude,omitempty"`
	UserID    string  `json:"userId,omitempty"`
	Timestamp int     `json:"timestamp,omitempty"`
}

var locations []geolocation

func serve(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(locations)
	}

	if r.Method == http.MethodPost {
		var location geolocation
		err := json.NewDecoder(r.Body).Decode(&location)
		if err != nil {
			log.Fatal(err)
			return
		}
		filename := fmt.Sprintf("user%s_t%d", location.UserID, location.Timestamp)
		file, err := os.Create(dataFileDir + filename)
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
var dataFileDir string

func init() {
	dataFileDir = os.Getenv("GEOLOCATION_DATA_FILES_DIR")
	if dataFileDir == "" {
		dataFileDir = "./app/data/"
	}
}
func main() {
	http.HandleFunc("/geolocation", serve)
	log.Fatal(http.ListenAndServe(":4000", nil))
}
