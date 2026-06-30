package main

import (
	"log"
	"tlms/internal/bootstraps"
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

	app := bootstraps.NewSeedApp(db)

	// seed Office
	if err := seeder.InitOfficeSeed(db); err != nil {
		log.Fatal(err)
	}

	// seed casbin
	if err = seeder.InitCasbinSeed(app.Authz, db); err != nil {
		log.Fatal(err)
	}

	// seed user admin
	if err = seeder.InitUserAdmin(app, db); err != nil {
		log.Fatal(err)
	}

	log.Printf("data seed done.")
}
