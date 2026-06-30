package main

import (
	"log"
	"tlms/internal/database"
	"tlms/internal/seeder"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	if err = seeder.SeedData(db); err != nil {
		log.Fatal(err)
	}

	log.Printf("data seed done.")
}
