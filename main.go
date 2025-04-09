// Connects to MongoDB and sets a Stable API version
package main

import (
	"fmt"
	"gobackend/connect"
	"gobackend/controller"

	"github.com/gofiber/fiber/v3"
)

func main() {
	
	connect.Connect()
	app := fiber.New()

	
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
	

	app.Get("api/helloUser", controller.HelloUser)
	// app.Post("/api/createUser", controller.CreateUser)
	app.Listen(":6000")

	
}
