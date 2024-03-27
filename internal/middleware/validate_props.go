package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	models "github.com/Crampustallin/houses/internal/models/objects"
	"github.com/Crampustallin/houses/internal/utils"
)

func ValidateProp(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var data models.Prop
		buf, err := io.ReadAll(r.Body)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, "couldn't read reques")
		}
		body := io.NopCloser(bytes.NewBuffer(buf))
		if err := json.Unmarshal(buf, &data); err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		}

		if data.PropertyTypeId == 0 || data.AddressId == 0 || data.Price == 0 {
			utils.RespondWithError(w, http.StatusBadRequest, "Missing required fields")
			return
		}

		r.Body = body
			
		next.ServeHTTP(w,r)
	})
}
