package JSON

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const citiBikeURL = "https://gbfs.citibikenyc.com/gbfs/en/station_status.json"

type stationData struct {
	LastUpdated int `json:"last_updated"`
	TTL         int `json:"ttl"`
	Data        struct {
		Stations []station `json:"stations"`
	} `json:"data"`
}

type station struct {
	StationID         string `json:"station_id"`
	NumBikesAvailable int    `json:"num_bikes_available"`
	NumBikesDisable   int    `json:"num_bikes_disable"`
	NumDocksAvailable int    `json:"num_docks_available"`
	NumDocksDisable   int    `json:"num_docks_disable"`
	IsInstalled       int    `json:"is_installed"`
	IsRenting         int    `json:"is_renting"`
	IsReturning       int    `json:"is_returning"`
	LastReported      int    `json:"last_reported"`
	HasAvailableKeys  bool   `json:"has_available_keys"`
}

func Simple() {
	//get json
	response, err := http.Get(citiBikeURL)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	//read body into byte
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	//unmarshal object
	var sd stationData
	if err := json.Unmarshal(body, &sd); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n\n", sd.Data.Stations[0])

	output, err := json.Marshal(sd)
	if err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile("./Data/DataCollection/citibike.json", output, 0644); err != nil {
		log.Fatal(err)
	}
}
