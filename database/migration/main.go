package main

import (
	"fmt"
	"log"

	"github.com/Lukmanern/go-starter/app/entity"
	"github.com/Lukmanern/go-starter/config"
	connector "github.com/Lukmanern/go-starter/database/sql-db"
)

func main() {
	// load environment variables
	conf := config.ReadConfig("./.env")
	fmt.Println(conf.GetDatabaseURI())
	connector.LoadGormDatabase()
	db := connector.LoadGormDatabase()
	err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").AutoMigrate(
		&entity.User{},
		&entity.Todo{},
	)
	if err != nil {
		log.Panic("error in migration db :", err)
	}
}
