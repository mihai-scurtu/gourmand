package main

import (
	"log"

	"github.com/mihai-scurtu/gourmand"
)

func main() {
	app := gourmand.NewApp()

	log.Println("Running crawler...")

	err := app.Run()
	if err != nil {
		log.Println(err)
	}

	log.Println("Finished.")
}
