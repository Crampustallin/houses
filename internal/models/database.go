package models

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type DB interface {
	Query(query string, args ...any) (*sql.Rows, error) 
}

func NewDB(con string) (*sql.DB, error) {
	db, err := sql.Open("postgres", con)
	if err != nil {
		return nil, err
	}
	return db, nil
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
		query += " LIMIT " + string(lim)
	}

	return db.Query(query)
}
