package main

import (
	"log"
	"fmt"

	"github.com/Crampustallin/houses/internal/config"
	"github.com/Crampustallin/houses/internal/database"
	"github.com/Crampustallin/houses/internal/models"
)

func main() {
	conf := config.NewConfig()
	fmt.Println(conf.ConnectionString())
	err := database.NewDB(conf.ConnectionString())
	if err != nil {
		log.Fatal("error connecting to database ", err)
	}
	defer database.Db.Close()


	server := models.NewServer(":8080")

	err = server.Open()
	if err != nil {
		log.Fatal(err)
	}
}
