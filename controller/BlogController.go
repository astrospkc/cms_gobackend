package controller

import (
	"context"
	"fmt"
	"gobackend/connect"
	"gobackend/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Blog struct{
	
	UserId		string 	`bson:"user_id" json:"user_id"`
	Title		string	`bson:"title" json:"title"`
	Content		string	`bson:"content" json:"content"`
	Tags		string	`bson:"tags,omitempty" json:"tags"`
	CoverImage 	string	`bson:"coverImage,omitempty" json:"coverImage"`
	
}

func CreateBlog() fiber.Handler{
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

		var p models.Blog
		if err := c.BodyParser(&p); err!=nil{
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":"Invalid request body",
			})
		}
		fmt.Println("users_email while create project: ", email)
		user_info,err := GetUserViaEmail(email)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "User not found"})
		}
		hex:=user_info.Id
		fmt.Println("hex: ", hex)
		blog := models.Blog{
			Id:primitive.NewObjectID(),
			UserId: hex.Hex(),
			Title: p.Title,
			Content: p.Content,
			Tags:p.Tags,
			CoverImage: p.CoverImage,
			
		}

		_,err = connect.BlogsCollection.InsertOne(context.Background(), blog)
		if err!=nil{
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "looks like some information is missing , try again",
		})

		
	}
	return c.JSON(fiber.Map{"success": "created",})
	}
}


func ReadBlog() fiber.Handler{
	return func(c *fiber.Ctx) error {
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
		user_info,err := GetUserViaEmail(email)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "User not found"})
		}

		fmt.Println("the user info : ", user_info)
		hex:= user_info.Id.Hex()
		cursor, err := connect.BlogsCollection.Find(context.TODO(),bson.M{"user_id":hex})
		if err!=nil{
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":"No Blogs could be found",
			})
		}

		var blogs []models.Blog
		if err:=cursor.All(context.TODO(), &blogs); err!=nil{
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":"Failed to parse blogs data",
			})
		}

		return c.JSON(blogs)
	}
}
// func ReadBlogWIthId() fiber.Handler{}
// func UpdateBlogWithEmail() fiber.Handler{}
// func DeleteBlog() fiber.Handler{}