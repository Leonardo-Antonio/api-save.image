package main

import (
	"log"

	"github.com/Leonardo-Antonio/api-save.image/src/app"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln(err)
	}
	app := app.NewServer()
	app.Middlewares()
	app.Routers()
	app.Listeing()
}
