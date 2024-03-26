package utils

import (
	"testing" 

	"github.com/Crampustallin/houses/internal/models/objects"
)

func ToJsonTest (t *testing.T) {
	needleJSON := `{
  "id": 1,
  "property_type_id": 123,
  "property_type": "Квартира",
  "address": "Ул. Ленина, д. 5, кв. 3",
  "price": 1200000,
  "rooms": 3,
  "area": 85.5,
  "description": "Просторная квартира с видом на реку"
  }`
  testProperty := models.Property{
        Id:             1,
        PropertyTypeId: 2,
        PropertyType:   "Квартира",
        Address:        "Ул. Ленина, д. 5, кв. 3",
        Price:          1200000,
        Rooms:          1,
        Area:           45.5,
        Description:    "Студия в новом доме с ремонтом и мебелью",
    }

    result, err := ToJson(testProperty)
    if err != nil {
	    t.Fatal("Failed to convert json", err)
    }

    if needleJSON != string(result) {
	    t.Fatalf("JSON objects doesn't math got:\n %v\n\nexpect:%v\n", result, needleJSON)
    }

}
