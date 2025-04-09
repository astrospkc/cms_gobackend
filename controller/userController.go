package controller

import (
	"context"
	"fmt"
	"gobackend/connect"
	"gobackend/models"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
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


var secretKey = []byte("secret")

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
    	c.Request().Header.SetCookie("token", tokenString)
		 err = bcrypt.CompareHashAndPassword(hashedPass, password)
    	fmt.Println(err)
		resp := &Response{
			Token:tokenString,
			User:user,
		}
		return c.JSON(resp)
	}
}




func InsertUser(user models.User) error{
	fmt.Println("this insertuser func is running")
	inserted, err := connect.UsersCollection.InsertOne(context.Background(), user)
	if err!= nil{
		log.Fatal(err)
		return err
	}
	fmt.Println("Inserted users in the db", inserted.InsertedID)
	return nil
}
