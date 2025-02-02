package handlers

import (
    "encoding/json"
    "net/http"
    "receipt-processor/storage"

    "github.com/gorilla/mux"
)

func GetPoints(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    points, exists := storage.Store.Get(id)
    if !exists {
        http.Error(w, "Receipt not found", http.StatusNotFound)
        return
    }

    response := map[string]int{"points": points}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

