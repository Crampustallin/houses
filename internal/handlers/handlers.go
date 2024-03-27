package handlers

import (
	"net/http"

	"github.com/Crampustallin/houses/internal/database"
	"github.com/Crampustallin/houses/internal/utils"
)

func GetListHandler(w http.ResponseWriter, r *http.Request) {
	list, err := GetList(10, database.QueryProperties)
	if err != nil {
		utils.RespondWithError(w, 500, err.Error())
		return
	}
	if list == nil {
		utils.RespondWithJSON(w, 200, map[string]string{"message": "message"})
		return
	}
	utils.RespondWithJSON(w, 200, list)
}

func FindHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	list, err := GetList(id, database.QueryPropertiesById)
	if err != nil {
		utils.RespondWithError(w, 500, err.Error())
		return
	}
	if list == nil {
		utils.RespondWithJSON(w, 200, map[string]string{"message": "message"})
		return
	}
	utils.RespondWithJSON(w, 200, list)
}

func PostHandler(w http.ResponseWriter, r *http.Request) {

}

