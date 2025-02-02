package handlers

import (
    "encoding/json"
    "net/http"
    "receipt-processor/models"
    "receipt-processor/storage"
    "receipt-processor/utils"

    "github.com/google/uuid"
)

func ProcessReceipt(w http.ResponseWriter, r *http.Request) {
    var receipt models.Receipt
    err := json.NewDecoder(r.Body).Decode(&receipt)
    if err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

    points := utils.CalculatePoints(receipt)
    id := uuid.New().String()

    storage.Store.Save(id, points)

    response := map[string]string{"id": id}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

