package utils

import (
	"encoding/json"

	"github.com/Crampustallin/houses/internal/models"
)

func ToJson(property models.Property) ([]byte, error) {
	return json.Marshal(property)
}
