// Connects to MongoDB and sets a Stable API version
package main

import (
	"fmt"
	"gobackend/connect"
	"gobackend/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)



func main() {
	
	
	app := fiber.New()
	app.Use(cors.New(cors.Config{
        AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
        AllowOrigins:     "http://localhost:3000/",
    
        AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
    }))
	connect.Connect()
	
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	routes.RegisterNormalRoutes(app)
	routes.RegisterAPIKeyRoutes(app)







	// app.Post("/api/createUser", controller.CreateUser)
	app.Listen(":8080")

	
}
