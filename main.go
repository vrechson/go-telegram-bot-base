package main

import (
	"log"

	_ "github.com/joho/godotenv/autoload"
	"github.com/whoismath/degausser/app"
	"github.com/whoismath/degausser/config"
)

func main() {
	c, err := config.Setup()
	if err != nil {
		log.Fatal(err)
	}

	a := app.CreateApp(c)
	a.Start()

}
