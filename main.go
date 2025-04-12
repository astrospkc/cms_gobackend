// Connects to MongoDB and sets a Stable API version
package main

import (
	"fmt"
	"gobackend/connect"
	"gobackend/controller"
	"gobackend/middleware"

	"github.com/gofiber/fiber/v2"
)



func main() {
	
	connect.Connect()
	app := fiber.New()

	
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	app.Post("/api/createUser", controller.CreateUser())
	app.Post("/api/login", controller.Login())
	app.Get("/api/getUser",middleware.FetchUser(),  controller.GetUser())
	
	// Project section 
	app.Post("/api/project/createProject", middleware.FetchUser(), controller.CreateProject())

	// app.Post("/api/createUser", controller.CreateUser)
	app.Listen(":6000")

	
}
