package main

import (
	"flag"
	"light-challenge/app"
	"light-challenge/config"
	"log"
)

func main() {
	configPath := flag.String("config", "./policy.yaml", "File path to the policy yaml")
	config.LoadPolicy(*configPath)
	app := app.New()

	err := app.Start()
	if err != nil {
		log.Fatal(err.Error())
	}
}
