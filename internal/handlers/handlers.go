package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Crampustallin/houses/internal/database"
	models "github.com/Crampustallin/houses/internal/models/objects"
	"github.com/Crampustallin/houses/internal/utils"
)

func GetListHandler(w http.ResponseWriter, r *http.Request) {
	list, err := GetList(10, database.QueryProperties)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if list == nil {
		utils.RespondWithError(w, http.StatusNotFound, "not found")
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, list)
}

func FindHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	list, err := GetList(id, database.QueryPropertiesById)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "something went wrong...")
		return
	}
	if list == nil {
		utils.RespondWithError(w, http.StatusNotFound, "not found")
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, list)
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var data models.Prop

	err := json.NewDecoder(r.Body).Decode(&data) 
	if err !=  nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "error while handling body...")
		return
	}

	err = insertProp(&data)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "error while saving the data...")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"OK": "saved"})
}
