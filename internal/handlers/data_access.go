package handlers

import (
	"database/sql"
	"errors"

	"github.com/Crampustallin/houses/internal/database"
	"github.com/Crampustallin/houses/internal/models/objects"
)




func GetList[T string | int](n T,  methodQ func(database.DB, T)(*sql.Rows, error)) ([]models.Property, error) {
	db := &database.Db
	if db == nil {
		return nil, errors.New("No connection to data base")
	}
	
	rows, err := methodQ(*db, n)
	if err != nil {
		return nil, err
	}
	defer rows.Close()


	props := make(map[int]models.Prop)
	addr := make(map[int]string)
	propType := make(map[int]string)

	for rows.Next() {
		var prop models.Prop
		if err := rows.Scan(&prop.Id, &prop.PropertyTypeId,
	&prop.AddressId, &prop.Price, &prop.Rooms, &prop.Area, &prop.Description); err != nil {
			return nil, err
		}
		props[prop.Id] = prop
		addr[prop.AddressId] = ""
		propType[prop.PropertyTypeId] = ""
	}

	if len(props) == 0 {
		return nil, nil
	}

	ids := make([]interface{}, len(addr), len(addr))

	i := 0

	for key := range addr {
		ids[i] = key
		i++
	}

	rows, err = database.QueryAddress(*db, ids)

	for rows.Next() {
		var addId int
		var name string
		if err := rows.Scan(&addId, &name); err != nil {
			return nil, err
		}
		addr[addId] = name
	}

	ids = make([]interface{}, len(propType), len(propType))

	i = 0
	for key := range propType {
		ids[i] = key
		i++
	}

	rows, err = database.QueryPropertyTypes(*db, ids)

	for rows.Next() {
		var typeId int
		var typeName string
		if err := rows.Scan(&typeId, &typeName); err != nil {
			return nil, err
		}
		propType[typeId] = typeName
	}


	properties := make([]models.Property, len(props), len(props))

	for _, val := range props {
		property := models.Property{Id: val.Id, 
		PropertyTypeId: val.PropertyTypeId, 
		PropertyType: propType[val.PropertyTypeId], 
		Address: addr[val.AddressId], 
		Price: val.Price, 
		Rooms: val.Rooms, 
		Area: val.Area, 
		Description: val.Description}
		properties = append(properties, property)
	}

	return properties, nil;
}
