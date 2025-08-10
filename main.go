package main

import (
	"book-manager/config"
	"book-manager/routes"
)

func main() {
	config.InitDB()
	r := routes.SetupRouter()
	r.Run(":8080") 
}
