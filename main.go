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

	// User section

	app.Post("/api/createUser", controller.CreateUser())
	app.Post("/api/login", controller.Login())
	app.Get("/api/getUser",middleware.FetchUser(),  controller.GetUser())
	
	
	// Project section 

	app.Post("/api/project/createProject", middleware.FetchUser(), controller.CreateProject())
	app.Put("/api/project/updateProject/:projectid", middleware.FetchUser(), controller.UpdateProject())
	app.Get("/api/project/readProject", middleware.FetchUser(), controller.ReadProject())
	app.Get("/api/project/readProjectWithId/:projectid", middleware.FetchUser(), controller.FindOneViaPID())
	app.Delete("/api/project/deleteProject/:projectid", middleware.FetchUser(), controller.DeleteProject())

	// Blog section

	app.Post("/api/blog/createBlog", middleware.FetchUser(), controller.CreateBlog())
	app.Get("/api/blog/readBlog", middleware.FetchUser(), controller.ReadBlog())
	app.Get("/api/blog/readBlog/:blogid", middleware.FetchUser(), controller.ReadBlogWIthId())
	app.Put("/api/blog/updateBlog/:blogid", middleware.FetchUser(), controller.UpdateBlogWithBlogId())
	app.Delete("/api/blog/deleteBlog/:blogid", middleware.FetchUser(), controller.DeleteBlog())

	// Link section
	// this section can be made as your second brain
	app.Post("/api/link/createLink", middleware.FetchUser(), controller.CreateLink())
	app.Get("/api/link/readLink", middleware.FetchUser(), controller.ReadLink())
	app.Get("/api/link/readLink/:linkid", middleware.FetchUser(), controller.ReadLinkWithLinkId())
	app.Put("/api/link/updateLink/:linkid", middleware.FetchUser(), controller.UpdateLinkWithLinkId())
	app.Delete("/api/link/deleteLink/:linkid", middleware.FetchUser(), controller.DeleteLinkWithLinkId())
	app.Delete("/api/link/deleteAllLink", middleware.FetchUser(), controller.DeleteAllLinks())





	



	





	// app.Post("/api/createUser", controller.CreateUser)
	app.Listen(":6000")

	
}
