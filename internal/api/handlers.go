package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/louisxponce/vehicleapi/internal/data"
)

type ApiHandler struct {
	Data *data.DataAccess
}

func NewApiHandler(dataAccess *data.DataAccess) *ApiHandler {
	return &ApiHandler{Data: dataAccess}
}

// GetAll
func (h *ApiHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	brand := strings.ToLower(r.URL.Query().Get("brand"))
	model := strings.ToLower(r.URL.Query().Get("model"))
	yearStr := strings.TrimSpace(r.URL.Query().Get("year"))

	vehicles, err := h.Data.GetVehicles(brand, model, yearStr)
	if err != nil {
		log.Printf("Data access error. %v", err)
		writeJSONError(w, http.StatusInternalServerError, "Data access error")
		return
	}
	if err := json.NewEncoder(w).Encode(vehicles); err != nil {
		log.Printf("Failed to encode response: %v", err)
		writeJSONError(w, http.StatusInternalServerError, "Failed to encode response")
	}
}

// GetSingle
func (h *ApiHandler) GetSingle(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	log.Println("Requested id:", id)
	v, err := h.Data.GetVehicle(id)
	if err != nil {
		log.Printf("Data access error. %v", err)
		writeJSONError(w, http.StatusInternalServerError, "Data access error")
		return
	}
	if v == nil {
		log.Printf("Couldn't find id: %s", id)
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}

func writeJSONError(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{
		"error": message,
	})
}
