package controller

import (
	"context"
	"errors"
	"fmt"
	"gobackend/connect"
	"log"
	"reflect"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func GetClaimsEmail(user interface{}) (string, error) {
	claims, ok := user.(jwt.MapClaims)
	if !ok {
		return "", errors.New("invalid JWT claims format")
	}

	email, ok := claims["aud"].(string)
	if !ok {
		return "", errors.New("invalid or missing 'aud' field")
	}

	return email, nil
}

func GetUserViaEmail(email string) (UserResponse, error)  {
	fmt.Println("users_email: ", email)
	var foundUser UserResponse
	err := connect.UsersCollection.FindOne(context.TODO(), bson.M{"email":email}).Decode(&foundUser)
	if err!=nil{
		log.Fatal("No such email exist")
	}
		fmt.Println(reflect.TypeOf(foundUser), foundUser)
	return foundUser, err
	
}
