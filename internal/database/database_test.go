package database 

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockDB struct {
	mock.Mock
	T assert.TestingT
	expectedQuery string
}

func (m *MockDB) Query(query string, args ...any) (*sql.Rows, error) {
	assert.Equal(m.T, m.expectedQuery, query)
	return nil, nil
}

func (m *MockDB) Exec(query string, args ...any) (sql.Result, error) {
	return nil, nil
}

func TestQuerries(t *testing.T) {
	m := new(MockDB)
	m.T = t

	m.On("Query", mock.Anything, mock.AnythingOfType("[]interface{}")).Return(nil, nil)

	m.expectedQuery = `SELECT 
	p.property_id, 
	p.property_type_id, 
	p.address_id, 
	p.price, 
	p.rooms, 
	p.area, 
	p.description FROM properties p LIMIT 5`

	_, err := QueryProperties(m, 5)
	assert.NoError(t,err)

	m.expectedQuery = `SELECT 
	p.property_id, 
	p.property_type_id, 
	p.address_id, 
	p.price, 
	p.rooms, 
	p.area, 
	p.description FROM properties p`

	_, err = QueryProperties(m, 0)
	assert.NoError(t,err)

	m.expectedQuery = `SELECT 
	p.property_id, 
	p.property_type_id, 
	p.address_id, 
	p.price, 
	p.rooms, 
	p.area, 
	p.description FROM properties p
	WHERE p.property_id = 1;`

	_, err = QueryPropertiesById(m, "1")
	assert.NoError(t,err)

	m.expectedQuery = `SELECT
	a.address_id,
	CONCAT(a.street, ', д.', a.house_number, ', кв. ', a.apartment_number) AS address
	FROM addresses a
	WHERE a.address_id IN ($1,$2,$3)`

	params := []interface{}{1,2,3}
	_, err = QueryAddress(m, params)
	assert.NoError(t,err)

	m.expectedQuery = `SELECT 
	p.property_type_id
	p.property_type_name FROM property_types p
	WHERE p.property_type_id IN ($1,$2,$3)`  

	_, err = QueryPropertyTypes(m, params)
	assert.NoError(t,err)
}
