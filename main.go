package main

import (
	"fmt"

	userController "github.com/Lukmanern/go-starter/app/module/user/controller"
	"github.com/Lukmanern/go-starter/config"
)

// app runner
func main() {
	fmt.Println("Hellowww World :D")
	userController.PrintName()

	conf := config.ReadConfig("./.env")
	fmt.Println(conf)
}
