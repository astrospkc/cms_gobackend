package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"gobackend/connect"
	"gobackend/models"
	"log"
	"net/http"
)

type Response struct{
	Message   string   `json:"message"`
	User		models.User `json:"user"`
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

// insert user , update user, delete user, read user
func CreateUser(w http.ResponseWriter, r *http.Request){
	// w.Header().Set("Content-Type", "application/json")
	// w.Header().Set("Allow-Content-Allow-Methods","POST")
	fmt.Println("create user")
	fmt.Println("request: ", r)
	fmt.Println("the r body", r.Body)

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err!=nil{
		log.Fatal("Decode went wrong", err)
		return
	}
	response := Response{
		Message:"Created successful",
		User : user,
	}
	w.WriteHeader(http.StatusCreated)
	err = InsertUser(user)
	if err!= nil{
		http.Error(w,"Failed to create user", http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(user)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("user created successfully")
	fmt.Println("User info inserted", response)
}