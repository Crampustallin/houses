package middleware

import (
	"net/http"
	"encoding/json"

	models "github.com/Crampustallin/houses/internal/models/objects"
	"github.com/Crampustallin/houses/internal/utils"
)

func ValidateProp(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var data models.Prop
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		}

		if data.PropertyTypeId == 0 || data.AddressId == 0 || data.Price == 0 {
			utils.RespondWithError(w, http.StatusBadRequest, "Missing required fields")
			return
		}
			
		next.ServeHTTP(w,r)
	})
}
