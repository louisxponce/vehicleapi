package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

func getAll(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	brand := strings.ToLower(r.URL.Query().Get("brand"))
	model := strings.ToLower(r.URL.Query().Get("model"))
	yearStr := r.URL.Query().Get("year")

	// Build up the SELECT statement depending on the different filters
	var args []interface{}
	clauses := []string{}

	if brand != "" {
		clauses = append(clauses, "LOWER(brand) LIKE ?")
		args = append(args, "%"+brand+"%")
	}

	if model != "" {
		clauses = append(clauses, "LOWER(model) LIKE ?")
		args = append(args, "%"+model+"%")
	}

	if yearStr != "" {
		clauses = append(clauses, "LOWER(year) LIKE ?")
		args = append(args, "%"+yearStr+"%")
	}

	query := "SELECT id, brand, model, year FROM vehicle"
	if len(clauses) > 0 {
		query += " WHERE " + strings.Join(clauses, " AND ")
	}
	log.Printf("%v", query)

	// Prepare the database row
	rows, err := db.Query(query, args...)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		log.Printf("Database error. %v", err)
		return
	}

	vehicles := []Vehicle{}
	for rows.Next() {
		var v Vehicle
		err := rows.Scan(&v.Id, &v.Brand, &v.Model, &v.Year)
		if err != nil {
			log.Printf("Scan error: %v", err)
			continue
		}
		vehicles = append(vehicles, v)
	}
	json.NewEncoder(w).Encode(vehicles)
}

func getSingle(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	log.Println("Requested id:", id)

	row := db.QueryRow("SELECT id, brand, model, year FROM vehicle WHERE id = ?", id)
	var v Vehicle
	err := row.Scan(&v.Id, &v.Brand, &v.Model, &v.Year)
	if err == sql.ErrNoRows {
		// log.Printf("No record found.")
		http.NotFound(w, r)
		return
	}
	if err != nil {
		http.Error(w, "Database error.", http.StatusInternalServerError)
		log.Printf("DB error: %v", err)
		return
	}
	// log.Printf("Record found. %v", v)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}
