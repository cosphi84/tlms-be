package main

import (
	"fmt"
	"log"
	"os"
	"tlms/internal/bootstraps"
	"tlms/internal/database"
	"tlms/internal/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	port := os.Getenv("APP_PORT")
	if port == "" {
		panic("APP_PORT is not set")
	}
	ginMode := os.Getenv("GIN_MODE")
	if ginMode == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	app := bootstraps.NewApp(db)

	r := gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1"})

	routes.SetupRoutes(r, app)

	if err := r.Run(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatal(err)
	}

}
