package utils

import (
	"encoding/json"

	"github.com/Crampustallin/houses/internal/models/objects"
)

func ToJson(property models.Property) ([]byte, error) {
	return json.Marshal(property)
}
