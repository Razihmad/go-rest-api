package main

import (
	"fmt"

	"github.com/Razihmad/go-rest-api/config"
	"github.com/Razihmad/go-rest-api/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	config.ConnectDB()
	fmt.Println("connected successfully")
	app := fiber.New()
	routes.SetUP(app)
	app.Listen(":3000")

}
