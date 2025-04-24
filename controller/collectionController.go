package controller

import (
	"context"
	"fmt"
	"gobackend/connect"
	"gobackend/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func CreateCollection() fiber.Handler{
	return func(c *fiber.Ctx) error {
		user := c.Locals("user")
		email,err := GetClaimsEmail(user)
		if err!= nil{
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":err.Error(),
			})
		}
		
		var col models.Collection
		if err := c.BodyParser(&col); err!=nil{
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":"Invalid request body",
			})
		}
		user_info, err := GetUserViaEmail(email)
		if err!=nil{
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":"User not found",
			})
		}
		hex:= user_info.Id
		collection := models.Collection{
			Id:primitive.NewObjectID(),
			UserId: hex,
			Title: col.Title,
			Description: col.Description,
		}

		_,err = connect.ColCollection.InsertOne(context.Background(),collection)
		if err!=nil{
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":"Try with different name or check for missing details",
			})
		}
		return c.JSON(fiber.Map{
			"succcess":"created",
		})

	}
}

func GetAllCollection() fiber.Handler{
	return func(c *fiber.Ctx) error{
		user := c.Locals("user")
		email,err := GetClaimsEmail(user)
		if err!= nil{
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":err.Error(),
			})
		}

		userInfo, err:= GetUserViaEmail(email)
		if err!=nil{
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":"Failed to fetch user_info",
			})
		}

		cursor, err := connect.ColCollection.Find(context.TODO(), bson.M{"user_id":userInfo.Id})
		if err!=nil{
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "No collection could be found",
			})
		}
		var collections []models.Collection
		if err := cursor.All(context.TODO(), &collections); err!=nil{
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to parse project data",
			})
		}
		fmt.Println(collections)
		return c.JSON(collections)
	}
}