package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/Crampustallin/houses/internal/database"
	models "github.com/Crampustallin/houses/internal/models/objects"
	"github.com/Crampustallin/houses/internal/utils"
)

func GetListHandler(w http.ResponseWriter, r *http.Request) {
	list, err := GetList(10, database.QueryProperties)
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

	dat, err := io.ReadAll(r.Body)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "couldn't read request")
		return
	}

	err = json.Unmarshal(dat, &data)
	if err !=  nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "something went wrong...")
		return
	}

	err = insertProp(&data)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "error while inserting data")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"OK": "saved"})
}
