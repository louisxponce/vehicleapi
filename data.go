package main

import (
	"encoding/json"
	"log"
	"os"
)

func loadData() {
	log.Printf("Loading data into memory...")
	data, err := os.ReadFile("data.json")
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(data, &rawVehicleData)

}
