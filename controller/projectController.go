package controller

import (
	"context"

	"gobackend/connect"
	"gobackend/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// import "github.com/gofiber/fiber/v2"

func CreateProject() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// first get the user email , for inserting to that userid
		user := c.Locals("user")
		claims,ok := user.(jwt.MapClaims)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid JWT claims format",
			})
		}
		email, ok := claims["aud"].(string)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid or missing  aud field",
			})
		}

		var p models.Project
		if err := c.BodyParser(&p); err!=nil{
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":"Invalid request body",
			})
		}

		user_info,err := GetUserViaEmail(email)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "User not found"})
		}
		project := models.Project{
			UserId: user_info.Id.Hex(),
			Title: p.Title,
			Description: p.Description,
			Tags:p.Tags,
			Thumbnail: p.Tags,
			GithubLink: p.GithubLink,
			LiveDemoLink: p.LiveDemoLink,
		}

		_,err = connect.ProjectCollection.InsertOne(context.Background(), project)
		if err!=nil{
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "looks like some information is missing , try again",
		})

		
	}
	return c.JSON(fiber.Map{"success": "created",})
	}
}