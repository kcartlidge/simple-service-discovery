package main

import (
	"log"

	config "github.com/kcartlidge/simples-config"
)

func reload() {
	log.Println("Reloading endpoints")
	c, err := config.CreateConfig("ssd.ini")
	if err != nil {
		log.Fatalln(err)
	}
	settings.mtx.Lock()
	settings.Endpoints = c.GetSection("ENDPOINTS")
	settings.mtx.Unlock()
	log.Println("Endpoints :", len(settings.Endpoints))
}
