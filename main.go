package main

import (
	"fmt"

	database "github.com/Lukmanern/go-starter/database/sql-db"
)

// app runner

func main() {
	fmt.Println("Hellowww World :D")
	db := database.LoadSQLDatabase()
	if db == nil {
		fmt.Println("error")
	}

	fmt.Println("Hellowww World :D")
}
