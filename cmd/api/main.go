package main

import (
	"log"
	"tlms/internal/bootstraps"
	"tlms/internal/database"
	"tlms/internal/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	app := bootstraps.NewApp(db)

	r := gin.Default()

	routes.SetupRoutes(r, app)

	if err := r.Run(":2323"); err != nil {
		log.Fatal(err)
	}

}
