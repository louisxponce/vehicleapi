package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func getAll(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	// brandFilter := strings.ToLower(r.URL.Query().Get("brandFilter"))

	// if brandFilter == "" {
	// 	log.Printf("no filter applied. returning unfiltered results.")
	// 	json.NewEncoder(w).Encode(rawVehicleData)
	// 	return
	// }

	brand := strings.ToLower(r.URL.Query().Get("brand"))
	model := strings.ToLower(r.URL.Query().Get("model"))
	yearStr := r.URL.Query().Get("year")

	if brand == "" && model == "" && yearStr == "" {
		log.Println("No filters applied. Returning unfiltered results.")
		json.NewEncoder(w).Encode(rawVehicleData)
		return
	}

	var year int
	if yearStr != "" {
		var err error
		year, err = strconv.Atoi(yearStr)
		if err != nil {
			http.Error(w, "Invalid year", http.StatusBadRequest)
			return
		}
	}

	log.Printf("brand: %s", brand)
	log.Printf("model: %s", model)
	log.Printf("year: %d", year)

	// Note. The slice below is initialized slice. It's possible to add to an uninitialzed slice but if the filter
	// doesn't match anything it will return null in the api and best practice is to return an empty arraay instead.
	filteredVehicleData := []Vehicle{}
	// var filteredVehicleData []Vehicle
	for _, v := range rawVehicleData {

		if brand != "" && !strings.Contains(strings.ToLower(v.Brand), brand) {
			continue
		}
		if model != "" && !strings.Contains(strings.ToLower(v.Model), model) {
			continue
		}
		if yearStr != "" && v.Year != year {
			continue
		}

		filteredVehicleData = append(filteredVehicleData, v)

		// if strings.Contains(strings.ToLower(v.Brand), brandFilter) {
		// 	filteredVehicleData = append(filteredVehicleData, v)
		// }
	}
	json.NewEncoder(w).Encode(filteredVehicleData)
}

func getSingle(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	log.Println("Requested id:", id)
	for _, v := range rawVehicleData {
		if v.Id == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(v)
			return
		}
	}
	http.NotFound(w, r)
}
