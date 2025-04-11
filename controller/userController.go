package controller

import (
	"context"
	"fmt"
	"gobackend/connect"
	"gobackend/models"
	"log"
	"os"
	"reflect"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"golang.org/x/crypto/bcrypt"
)

// TODO: later on add Project , category , links, blog, media, resume, subscription, usersubscription, apikey , all of these in UserResponse
type UserResponse struct {
	Id   primitive.ObjectID `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	ProfilePic string `json:"profile_pic,omitempty"`
	Role 		string	`json:"role"`
	
}

type Response struct{
	Token   string   `json:"token"`
	User		models.User `json:"user"`
}

// var secretKey = []byte(os.Getenv("JWT_SECRET"))


// first createtoken
func CreateToken(userid string) (string, error){
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userid,                    // Subject (user identifier)
		"iss": "One-Go",                  // Issuer
		"aud": userid,           // Audience (user role)
		"exp": time.Now().Add(time.Hour).Unix(), // Expiration time
		"iat": time.Now().Unix(),                 // Issued at
	})
	secret :=[]byte(os.Getenv("JWT_SECRET"))
	tokenString, err := claims.SignedString(secret)
	if err!=nil{
		return "", err
	}
	fmt.Printf("Token claims added: %+v\n", claims)
	return tokenString, nil
}

// insert user , update user, delete user, read user

func CreateUser() fiber.Handler {
	return func(c *fiber.Ctx) error {

    	var d models.User
    	if err := c.BodyParser(&d); err!=nil{
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":"Invalid request body",
			})
		}

		
		if d.Email == "" || d.Password == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Email and Password are required",
			})
		}


		// hashing the password
		password:=[]byte(d.Password)
		hashedPass, err := bcrypt.GenerateFromPassword(password,bcrypt.DefaultCost)
		if err!=nil{
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Could not hash password",
			})
		}
		fmt.Println(string(hashedPass))


		if d.ProfilePic == "" {
			d.ProfilePic = "https://cdn.example.com/default-avatar.png"
		}
		user := models.User{
			Id:primitive.NewObjectID(),
			Name:d.Name,
			Email:d.Email,
			ProfilePic: d.ProfilePic,
			Password:string(hashedPass) ,
		}
		fmt.Println("user: ", user)
		inserted,err := connect.UsersCollection.InsertOne(context.Background(), user)
		if err!=nil{
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "looks like email address is already in use",
			})
			
		}
		fmt.Println("inserted user: ", reflect.TypeOf(inserted.InsertedID))
		
		tokenString,err := CreateToken(d.Email)
		if err!=nil{
			log.Println("failed to create token")
		}
		fmt.Println(tokenString, "token")
    	c.Cookie(&fiber.Cookie{
			Name: "token",
			Value:tokenString,
			HTTPOnly: true,
			Secure: true,
			Path:"/",
			MaxAge: 3600,
		})
		
		resp := &Response{
			Token:tokenString,
			User:user,
		}
		return c.JSON(resp)
	}
}


func Login() fiber.Handler{
	return func(c *fiber.Ctx) error{
		
		var d models.User
		if err := c.BodyParser(&d); err !=nil{
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":"Invalid request body",
			})
		}
		if d.Email==""||d.Password==""{
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":"Email and Password are required",
			})
		}
		var user models.User
		err := connect.UsersCollection.FindOne(context.TODO(), bson.M{"email":d.Email}).Decode(&user)
		if err!=nil{
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":"NO user with this email",
			})
		}
		fmt.Println("user: ", user)
		password := []byte(d.Password)
		err = bcrypt.CompareHashAndPassword([]byte(user.Password),password )
		if err!=nil{
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":"Password is incorrect, please try once more",
			})
		}
		tokenString, err := CreateToken(d.Email)
		if err!=nil{
			log.Println("failed to create token")
		}
		c.Cookie(&fiber.Cookie{
			Name: "token",
			Value:tokenString,
			HTTPOnly: true,
			Secure: true,
			Path:"/",
			MaxAge: 3600,
		})
		resp := &Response{
			Token:tokenString,
			User:user,
		}
		return c.JSON(resp)

	}
}


func GetUser() fiber.Handler{
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
		
		// check this email exist or not , and if exist , then provide the user details with every field 
		var foundUser UserResponse
		err := connect.UsersCollection.FindOne(context.TODO(), bson.M{"email":email}).Decode(&foundUser)
		if err!=nil{
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":"NO user with this email",
			})
		}
		fmt.Println(foundUser)
		return c.JSON(foundUser)
	}
}
// from user response , please remove password section
