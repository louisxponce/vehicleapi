package data

import (
	"database/sql"
	"log"
	"strings"

	"github.com/louisxponce/vehicleapi/internal/models"
	_ "github.com/mattn/go-sqlite3"
)

type DataAccess struct {
	DB *sql.DB
}

func NewDataAccess() *DataAccess {

	db, err := sql.Open("sqlite3", "vehicles.db")
	if err != nil {
		log.Fatalf("Could not open the database: %v", err)
	}

	// Use WAL for better concurrency
	if _, err := db.Exec("PRAGMA journal_mode=WAL;"); err != nil {
		log.Fatalf("Failed to enable WAL mode: %v", err)
	}

	// Increase cache size (~100 MB per connection)
	_, err = db.Exec("PRAGMA cache_size = -102400;") // 100 MB per connection
	if err != nil {
		log.Fatalf("Failed to set cache size: %v", err)
	}

	row := db.QueryRow("PRAGMA cache_size;")
	var size int
	if err := row.Scan(&size); err != nil {
		log.Printf("Failed to read back cache_size: %v", err)
	} else {
		log.Printf("Effective cache_size: %d pages (~%.2f MB)", size, float64(-size)*4096/1024/1024)
	}

	// Store temp tables in memory
	if _, err := db.Exec("PRAGMA temp_store = MEMORY;"); err != nil {
		log.Fatalf("Failed to set temp_store: %v", err)
	}

	// Map DB file into memory (256 MB)
	if _, err := db.Exec("PRAGMA mmap_size = 268435456;"); err != nil {
		log.Printf("mmap_size not supported or ignored: %v", err)
	}

	// Set connection limits
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)

	log.Printf("Database initialized and tuned for memory.")

	return &DataAccess{DB: db}
}

// GetVehicles
func (r *DataAccess) GetVehicles(brand string, model string, yearStr string) ([]models.Vehicle, error) {
	var args []any
	var clauses []string
	if strings.Contains(brand, "*") {
		clauses = append(clauses, "brand_lc LIKE ?")
		args = append(args, strings.ToLower(strings.ReplaceAll(brand, "*", "%")))
	} else if brand != "" {
		clauses = append(clauses, "brand_lc = ?")
		args = append(args, brand)
	}

	if strings.Contains(model, "*") {
		clauses = append(clauses, "model_lc LIKE ?")
		args = append(args, strings.ToLower(strings.ReplaceAll(model, "*", "%")))
	} else if model != "" {
		clauses = append(clauses, "model_lc = ?")
		args = append(args, model)
	}

	if yearStr != "" {
		clauses = append(clauses, "year = ?")
		args = append(args, yearStr)
	}

	query := "SELECT id, brand, model, year FROM vehicle"
	if len(clauses) > 0 {
		query += " WHERE " + strings.Join(clauses, " AND ")
	}
	log.Printf("%v", query)

	// Prepare the database row
	rows, err := r.DB.Query(query, args...)
	if err != nil {
		log.Printf("Database error. %v", err)
		return nil, err
	}

	vehicles := []models.Vehicle{}
	for rows.Next() {
		var v models.Vehicle
		err := rows.Scan(&v.Id, &v.Brand, &v.Model, &v.Year)
		if err != nil {
			log.Printf("Scan error: %v", err)
			continue
		}
		vehicles = append(vehicles, v)
	}
	return vehicles, nil
}

// GetVechicle
func (r *DataAccess) GetVehicle(id string) (*models.Vehicle, error) {
	row := r.DB.QueryRow("SELECT id, brand, model, year FROM vehicle WHERE id = ?", id)
	var v models.Vehicle
	err := row.Scan(&v.Id, &v.Brand, &v.Model, &v.Year)
	if err == sql.ErrNoRows {
		return nil, err
	}
	if err != nil {
		log.Printf("DB error: %v", err)
		return nil, err
	}
	return &v, nil
}
