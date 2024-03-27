package utils

import (
	"net/http"
	"encoding/json"
)

func ToJson(property interface{}) ([]byte, error) {
	return json.Marshal(property)
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) error {
	response, err := ToJson(payload) 
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(code)
	w.Write(response)
	return nil
}

func RespondWithError(w http.ResponseWriter, code int, message string) error {
	return RespondWithJSON(w, code, map[string]string{"error":message}) 
}
