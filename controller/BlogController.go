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
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Blog struct{
	Id 			primitive.ObjectID	`bson:"id" json:"id"`
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
func ReadBlogWIthId() fiber.Handler{
	return func(c *fiber.Ctx) error {
		p_id:= c.Params("blogid")
		if p_id==""{
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":"Please provide project id",
			})
		}
		objID, err := primitive.ObjectIDFromHex(p_id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid project ID format",})
		}

		var blog Blog
		err = connect.BlogsCollection.FindOne(context.TODO(), bson.M{"id":objID}).Decode(&blog)
		if err!=nil{
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":"Failed to find the blogs with given id",
			})
		}
		return c.JSON(blog)
	}
}


func setBlog(upd *Blog) (bson.M, error){
	data, err := bson.Marshal(upd)
	if err!=nil{
		return nil, err
	}
	var m bson.M
	if err:= bson.Unmarshal(data, &m); err!=nil{
		return nil, err
	}
	return m, nil
}
func UpdateBlogWithBlogId() fiber.Handler{
	return func(c *fiber.Ctx) error {
		b_id := c.Params("blogid")
		if b_id==""{
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":"Please provide blog id",
			})
		}

		var upd Blog
		if err := c.BodyParser(&upd); err!=nil{
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":"invalid JSON",
			})
		}

		setBlog, err := setBlog(&upd)
		if err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to prepare update"})
        }
        if len(setBlog) == 0 {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "No fields provided to update"})
        }

		objId, err := primitive.ObjectIDFromHex(b_id)
		if err!=nil{
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":"Failed to convert in primitive type",
			})
		}

		filter := bson.M{"id":objId} 
		update:= bson.M{"$set":setBlog}
		opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
		
		var updatedBlog models.Blog
		err = connect.BlogsCollection.FindOneAndUpdate(context.TODO(), filter, update, opts).Decode(&updatedBlog)
		if err !=nil{
			if err ==mongo.ErrNoDocuments{
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error":"Blog not found"})
			}
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":"Update failed"})
		}
		return c.JSON(updatedBlog)
	}
}
func DeleteBlog() fiber.Handler{
	return func(c *fiber.Ctx) error {
		p_id:= c.Params("blogid")
		if p_id==""{
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":"Please provide blog id",
			})
		}
		objId, err := primitive.ObjectIDFromHex(p_id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Blog ID format",})
		}

		filter := bson.M{"id":objId}
		
		result, err := connect.BlogsCollection.DeleteOne(context.TODO(),filter)
		if err!=nil{
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"eror":"Blog was not deleted successful"})
		}

		return c.JSON(result)

	}
}