package main

import (
	"fmt"

	user_module "github.com/Lukmanern/go-starter/app/module/user/controller"
)

// app runner
func main() {
	fmt.Println("Hellowww World :D")
	user_module.PrintName()
}
