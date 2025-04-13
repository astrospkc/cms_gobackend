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

type ProjectUpdate struct {
    Title        *string `json:"title,omitempty" bson:"title,omitempty"`
    Description  *string `json:"description,omitempty" bson:"description,omitempty"`
    Tags         *string `json:"tags,omitempty" bson:"tags,omitempty"`
    Thumbnail    *string `json:"thumbnail,omitempty" bson:"thumbnail,omitempty"`
    GithubLink   *string `json:"githublink,omitempty" bson:"githublink,omitempty"`
    LiveDemoLink *string `json:"livedemolink,omitempty" bson:"livedemolink,omitempty"`
    // (no CreatedAt here)
}

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
		fmt.Println("users_email while create project: ", email)
		user_info,err := GetUserViaEmail(email)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "User not found"})
		}
		hex:=user_info.Id
		fmt.Println("hex: ", hex)
		project := models.Project{
			Id:primitive.NewObjectID(),
			UserId: hex,
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

func ReadProject() fiber.Handler{
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
		// we can't get project list using email , we need user_id for it.
		userInfo, err:= GetUserViaEmail(email)
		if err!=nil{
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":"Failed to fetch user_info",
			})
		}
		
		
		// var project_info models.Project
		cursor,err := connect.ProjectCollection.Find(context.TODO(), bson.M{"user_id":userInfo.Id})

		// cursor,err := connect.ProjectCollection.Find(context.TODO(), bson.M{"email":email})
		if err!=nil{
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "No project could be found",
			})
		}
		
		var projects []models.Project
		if err := cursor.All(context.TODO(), &projects); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to parse project data",
			})
		}
		fmt.Println(projects)
		return c.JSON(projects)
	}
}

func setDoc(upd *ProjectUpdate) (bson.M, error){
	data, err:= bson.Marshal(upd)
	if err!=nil{
		return nil,err
	}

	var m bson.M
	if err:= bson.Unmarshal(data, &m); err!=nil{
		return nil, err
	}
	return m, nil
}

func UpdateProject() fiber.Handler {
	return func(c *fiber.Ctx) error{
		p_id := c.Params("projectid")
		if p_id == ""{
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":"Please provide project id",
			})
		}
		
		var upd ProjectUpdate
		if err := c.BodyParser(&upd); err!=nil{
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":"Invalid JSON",
			})
		}

		// build the $set doc
		setDoc, err:=  setDoc(&upd)
		 if err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to prepare update"})
        }
        if len(setDoc) == 0 {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "No fields provided to update"})
        }

		fmt.Println("setDoc: ", setDoc)
		objId, err := primitive.ObjectIDFromHex(p_id)
		if err !=nil{
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":"Failed to convert in primitive type",
			})
		}


		fmt.Println("object Id: ", objId)
		// lets find first to check it is actually working or not
		var foundProj models.Project
		err = connect.ProjectCollection.FindOne(context.TODO(),bson.M{"id":objId}).Decode(&foundProj)
		if err!=nil{
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":"Could not find the project",
			})
		}
		fmt.Println("foundProj: ", foundProj)


		fmt.Printf("pid: %T \n", objId)
		filter:=bson.M{"id":objId}
		update:=bson.M{"$set":setDoc}
		opts:=options.FindOneAndUpdate().SetReturnDocument(options.After)

		var updatedDoc models.Project
		err = connect.ProjectCollection.FindOneAndUpdate(context.TODO(),filter, update, opts).Decode(&updatedDoc)
		if err != nil {
            if err == mongo.ErrNoDocuments {
                return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Project not found"})
            }
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Update failed"})
        }

        return c.JSON(updatedDoc)
	}
}

// func UpdateProject() fiber.Handler{
	
// 	return func(c *fiber.Ctx) error {
// 		user := c.Locals("user")
// 		claims,ok := user.(jwt.MapClaims)
// 		if !ok {
// 			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 				"error": "Invalid JWT claims format",
// 			})
// 		}
// 		email, ok := claims["aud"].(string)
// 		if !ok {
// 			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 				"error": "Invalid or missing  aud field",
// 			})
// 		}

// 		var update models.Project
// 		if err := c.BodyParser(&update); err!=nil{
// 			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 				"error":"Invalid request body",
// 			})
// 		}

// 		user_info,err := GetUserViaEmail(email)
// 		if err != nil {
// 			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "User not found"})
// 		}
// 		project := models.Project{
// 			UserId: user_info.Id.Hex(),
// 			Title: update.Title,
// 			Description: update.Description,
// 			Tags:update.Tags,
// 			Thumbnail: update.Thumbnail,
// 			GithubLink: update.GithubLink,
// 			LiveDemoLink: update.LiveDemoLink,
// 		}
// 		filer := bson.M{"_id":user_info.Id}
// 		result, err := connect.BlogsCollection.UpdateByID(context.TODO(), filer, project)
// 		if err!=nil{
// 			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 				"error": "Failed to update project",
// 			})
// 		}

// 		if result.MatchedCount==0{
// 			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
// 				"error": "Project not found",
// 			})
// 		}
// 		return c.JSON(result)
// 		// return c.JSON(fiber.Map{
// 		// 	"success": "Blog updated successfully",
// 		// })
// 	}
// }



// func DeleteProject() fiber.Handler{}