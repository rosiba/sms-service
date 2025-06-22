package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"sms-service/api"
	"sms-service/internal/db"
	"sms-service/internal/delivery"
	"sms-service/internal/postgres"
)

func main() {
	log.Println("starting sms-service")

	log.Println("loading environment variables")
	if err := readConfig(); err != nil {
		log.Fatalf("failed to read config: %v", err)
	}

	log.Println("starting a connection to database")
	conn, err := db.Connect()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer conn.Close()

	if err := db.InitDB(conn); err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}

	mr := postgres.NewMessageRepository(conn)
	ds := delivery.NewDeliveryService(mr)
	go ds.Run()

	log.Println("starting http server")
	h := api.NewHandler(mr, ds)
	router, err := api.NewRouter(h)
	if err != nil {
		log.Fatalf("failed to create router: %v", err)
	}
	port := fmt.Sprintf(":%s", os.Getenv("WEB_PORT"))
	if err := router.Run(port); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}

func readConfig() error {
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("failed to load .env file: %v", err)
	}

	return nil
}
