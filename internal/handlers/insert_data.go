package handlers

import (
	"github.com/Crampustallin/houses/internal/database"
	"github.com/Crampustallin/houses/internal/models/objects"
)

func insertProp(p *models.Prop) error {
	db := &database.Db
	_, err := database.InsertProperty(*db, *p)
	if err != nil {
		return err
	}
	return nil
}
