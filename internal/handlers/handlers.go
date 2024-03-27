package handlers

import (
	"net/http"

	"github.com/Crampustallin/houses/internal/database"
	"github.com/Crampustallin/houses/internal/utils"
)

func GetListHandler(w http.ResponseWriter, r *http.Request) {
	list, err := GetList(10, database.QueryProperties)
	if err != nil {
		respondWithError(w, 500, err.Error())
		return
	}
	if list == nil {
		respondWithJSON(w, 200, map[string]string{"message": "message"})
		return
	}
	respondWithJSON(w, 200, list)
}

func FindHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	list, err := GetList(id, database.QueryPropertiesById)
	if err != nil {
		respondWithError(w, 500, err.Error())
		return
	}
	if list == nil {
		respondWithJSON(w, 200, map[string]string{"message": "message"})
		return
	}
	respondWithJSON(w, 200, list)
}

func PostHandler(w http.ResponseWriter, r *http.Request) {

}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) error {
	response, err := utils.ToJson(payload) 
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(code)
	w.Write(response)
	return nil
}

func respondWithError(w http.ResponseWriter, code int, message string) error {
	return respondWithJSON(w, code, map[string]string{"error":message}) 
}
