package main

import (
	"log"
	"webserver/internal/app/webserver"
)

func main() {
	webserver := webserver.New()

	if err := webserver.Start(); err != nil {
		log.Fatal(err)
	}
}
