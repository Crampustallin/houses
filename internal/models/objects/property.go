package models

type Property struct {
	Id uint `json:"id"`
	PropertyTypeId uint `json:"property_type_id"`
	PropertyType string `json:"property_type" `
	Address string `json:"address"`
	Price uint `json:"price"`
	Rooms uint `json:"rooms"`
	Area float32 `json:"area"`
	Description string `json:"description"`
}
