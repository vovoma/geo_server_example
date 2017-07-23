package main

import (
	"encoding/json"
	"log"
	"net/http"
	"fmt"
	"os"
)

type Geolocation struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	UserID    string  `json:"userId"`
	Timestamp int     `json:"timestamp"`
}

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
		filename := fmt.Sprintf("user%s_t%d",location.UserID,location.Timestamp)
		file, err := os.Create("./app/data/"+filename)
		if err != nil {
			log.Fatal(err)
		}
		data, _ := json.Marshal(location)
		file.Write(data)
		log.Printf("File create %s\n",filename)

		locations = append(locations, location)
		json.NewEncoder(w).Encode(location)
		defer r.Body.Close()
	}
	return
}

func main() {
	http.HandleFunc("/geolocation", serve)
	log.Fatal(http.ListenAndServe(":4000", nil))
}
