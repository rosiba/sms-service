package main

import (
	"log"
	"sms-service/api"
)

func main() {
	router := api.NewRouter()
	if err := router.Run(":8080"); err != nil {
		log.Panic(err)
	}
}
