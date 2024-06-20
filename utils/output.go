package utils

import (
	"encoding/json"
	"net/http"
)

type Responses struct {
	Error   string
	Success bool
	Data    interface{}
}

func CreateOutput(w http.ResponseWriter, err error, success bool, data any) {
	response := Responses{Error: err.Error(), Success: success, Data: data}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&response)
}
