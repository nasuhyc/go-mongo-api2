package main

import (
	"go-mongo-api2/config"
	"go-mongo-api2/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	app := fiber.New()
	app.Use(cors.New())

	config.ConnectDB()
	port := os.Getenv("PORT")
	routes.Setup(app)
	err := app.Listen(":" + port)

	if err != nil {
		log.Fatal("Error app failed to start")
		panic(err)
	}
}
