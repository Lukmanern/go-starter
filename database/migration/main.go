package main

import (
	"fmt"
	"log"

	todoEntity "github.com/Lukmanern/go-starter/app/module/todo/entity"
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
		&todoEntity.User{},
		&todoEntity.Todo{},
	)
	if err != nil {
		log.Panic("\n\n :: -1 :: ", err)
	}
}
