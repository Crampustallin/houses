package models

type Property struct {
	Id int `json:"id"`
	PropertyTypeId int `json:"property_type_id"`
	PropertyType string `json:"property_type" `
	Address string `json:"address"`
	Price float64 `json:"price"`
	Rooms int `json:"rooms"`
	Area float32 `json:"area"`
	Description string `json:"description"`
}

type Prop struct {
	Id int `json:"id"`
	PropertyTypeId int `json:"property_type_id"`
	AddressId int `json:"address_id"`
	Price float64 `json:"price"`
	Rooms int `json:"rooms"`
	Area float32 `json:"area"`
	Description string `json:"description"`
}
