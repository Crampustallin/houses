package database

import (
	"database/sql"
	"strconv"
	"strings"

	models "github.com/Crampustallin/houses/internal/models/objects"
	_ "github.com/lib/pq"
)

type DB interface {
	Query(query string, args ...any) (*sql.Rows, error) 
	Exec(query string, args ...any) (sql.Result, error)
}

var Db *sql.DB

func NewDB(con string) error {
	var err error
	Db, err = sql.Open("postgres", con)
	if err != nil {
		return err
	}
	return nil
}

func QueryProperties(db DB, lim int) (*sql.Rows, error) {
	query := `SELECT 
	p.property_id, 
	p.property_type_id, 
	p.address_id, 
	p.price, 
	p.rooms, 
	p.area, 
	p.description FROM properties p`

	if lim > 0 {
		query += " LIMIT " + strconv.Itoa(lim)
	}

	return db.Query(query)
}

func QueryPropertiesById(db DB, id string) (*sql.Rows, error) {
	query := `SELECT 
	p.property_id, 
	p.property_type_id, 
	p.address_id, 
	p.price, 
	p.rooms, 
	p.area, 
	p.description FROM properties p
	WHERE p.property_id = ` + id +";"

	return db.Query(query)
}

func QueryAddress(db DB, ids []interface{}) (*sql.Rows, error) {
	query := `SELECT
	a.address_id,
	CONCAT(a.street, ', д.', a.house_number, ', кв. ', a.apartment_number) AS address
	FROM addresses a
	WHERE a.address_id IN (` + FormatParamQuery(ids) + ")"

	return db.Query(query, ids...)
}

func QueryPropertyTypes(db DB, ids []interface{}) (*sql.Rows, error) {
	query := `SELECT 
	p.property_type_id,
	p.property_type_name FROM property_types p
	WHERE p.property_type_id IN (` + FormatParamQuery(ids) + ")"

	return db.Query(query, ids...)
}

func InsertProperty(db DB, prop models.Prop) (sql.Result, error) {
	query := `INSERT INTO properties (property_type_id, address_id, price, rooms, area, description)
	VALUES ($1, $2, $3, $4, $5, $6)`

	return db.Exec(query, prop.PropertyTypeId, prop.AddressId, prop.Price, prop.Rooms, prop.Area, prop.Description)
}

func FormatParamQuery[T any](params []T) string {
	length := len(params)
	placeholders := make([]string, length, length)
	for i := range placeholders {
		placeholders[i] = "$" + strconv.Itoa(i+1)
	}
	return strings.Join(placeholders, ",")
}
