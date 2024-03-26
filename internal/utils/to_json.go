package utils

import (
	"encoding/json"

)

func ToJson(property interface{}) ([]byte, error) {
	return json.Marshal(property)
}
