// Connects to MongoDB and sets a Stable API version
package main

import (
	"fmt"
	"gobackend/connect"
	"gobackend/controller"

	"github.com/gofiber/fiber/v2"
)



func main() {
	
	connect.Connect()
	app := fiber.New()

	
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	app.Post("/body", controller.CreateUser())
	app.Post("/login", controller.Login())

	// app.Post("/api/createUser", controller.CreateUser)
	app.Listen(":6000")

	
}
