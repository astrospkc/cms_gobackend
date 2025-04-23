package connect

import (
	"context"
	"fmt"

	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const (
	dbName  = "CMS_portfolio"
	colCollection = "collections"
	colNameUsers = "users"
	colNameProjects = "projects"
	colNameBlogs = "blogs"
	colNameLinks = "links"
	colNameGithub = "github"
	colNameSketches ="sketches"
	colNameDesigns  ="designs"
	colNameAPI  = "apikey"
)

var UsersCollection *mongo.Collection
var ColCollection    *mongo.Collection
var ProjectCollection *mongo.Collection
var BlogsCollection *mongo.Collection
var LinksCollection *mongo.Collection
var GithubCollection *mongo.Collection
var SkectchesCollection *mongo.Collection
var DesignsCollection *mongo.Collection
var TeaCollection *mongo.Collection
var APIKeyCollection *mongo.Collection


func Connect(){
	err:=godotenv.Load(".env.prod")
	if err!=nil{
		log.Fatal("Error handling .env file")
	}
	var uri string 
	// uri := "mongodb+srv://punamkumari399:RDQX28rR3RIh3V9m@cluster0.g24hw.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
	if uri = os.Getenv("MONGODB_URI"); uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environment variable. See\n\t https://docs.mongodb.com/drivers/go/current/usage-examples/")
	}

	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	ctx, cancel:= context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	client, err := mongo.Connect(opts)

	if err != nil {
		log.Fatal("the error while connecting client", err)
		return
	}
	// defer func() {
	// 	fmt.Println("Disconnecting")
	// 	if err = client.Disconnect(context.TODO()); err != nil {
	// 		log.Fatal("error while disconnecting: ", err)
	// 		return
	// 	}
	// }()
	
	err = client.Ping(ctx, nil)
	
	if err != nil{
		log.Fatal("ping :", err)
		return
	}
	UsersCollection = client.Database(dbName).Collection(colNameUsers)
	indexModel:=mongo.IndexModel{
		Keys:bson.D{{Key:"email", Value:  -1}},
		Options:options.Index().SetUnique(true),
	}
	_, err =UsersCollection.Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil{
		
		log.Fatal("error occured : ", err)
	}
	
	ColCollection  = client.Database(dbName).Collection(colCollection)
	ProjectCollection = client.Database(dbName).Collection(colNameProjects)
	BlogsCollection = client.Database(dbName).Collection(colNameBlogs)
	LinksCollection = client.Database(dbName).Collection(colNameLinks)
	GithubCollection = client.Database(dbName).Collection(colNameGithub)
	APIKeyCollection = client.Database(dbName).Collection(colNameAPI)
	// TeaCollection = client.Database("CMS_portfolio").Collection("tea")
	
	// insert tea database

	fmt.Println("Set up is done")
	

	
}