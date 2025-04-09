package controller

import (
	"context"
	"fmt"
	"gobackend/connect"
	"gobackend/models"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/v2/bson"
	"golang.org/x/crypto/bcrypt"
)

type Response struct{
	Token   string   `json:"token"`
	User		models.User `json:"user"`
}
type Person struct {
    Name string `json:"name" xml:"name" form:"name"`
    Pass string `json:"pass" xml:"pass" form:"pass"`
}


var secretKey = []byte(os.Getenv("JWT_SECRET"))

// first createtoken
func CreateToken(userid string) (string, error){
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userid,                    // Subject (user identifier)
		"iss": "One-Go",                  // Issuer
		"aud": userid,           // Audience (user role)
		"exp": time.Now().Add(time.Hour).Unix(), // Expiration time
		"iat": time.Now().Unix(),                 // Issued at
	})

	tokenString, err := claims.SignedString(secretKey)
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
			Name:d.Name,
			Email:d.Email,
			ProfilePic: d.ProfilePic,
			Password:string(hashedPass) ,
		}
		fmt.Println("user: ", user)
		inserted,err := connect.UsersCollection.InsertOne(context.Background(), user)
		if err!=nil{
			log.Fatal("looks like user can't be created", err)
			
		}
		fmt.Println("inserted user: ", inserted.InsertedID)
		tokenString,err := CreateToken(d.Name)
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
		tokenString, err := CreateToken(user.Name)
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

// from user response , please remove password section
