package controller

import (
	"context"

	"fmt"

	"gobackend/connect"
	"gobackend/models"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)





func InsertItem(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	
	docs := []interface{}{
	models.Tea{Type: "Masala", Category: "black", Toppings: []string{"ginger", "pumpkin spice", "cinnamon"}, Price: 6.75},
	models.Tea{Type: "Gyokuro", Category: "green", Toppings: []string{"berries", "milk foam"}, Price: 5.65},
	models.Tea{Type: "English Breakfast", Category: "black", Toppings: []string{"whipped cream", "honey"}, Price: 5.75},
	models.Tea{Type: "Sencha", Category: "green", Toppings: []string{"lemon", "whipped cream"}, Price: 5.15},
	models.Tea{Type: "Assam", Category: "black", Toppings: []string{"milk foam", "honey", "berries"}, Price: 5.65},
	models.Tea{Type: "Matcha", Category: "green", Toppings: []string{"whipped cream", "honey"}, Price: 6.45},
	models.Tea{Type: "Earl Grey", Category: "black", Toppings: []string{"milk foam", "pumpkin spice"}, Price: 6.15},
	models.Tea{Type: "Hojicha", Category: "green", Toppings: []string{"lemon", "ginger", "milk foam"}, Price: 5.55},
	}
	result, err := connect.TeaCollection.InsertMany(context.TODO(), docs)
	if err!=nil{
		log.Fatal("Error inserting doucment", err)
	}
	fmt.Println("Inserted document ids: ", result.Acknowledged)

	// now getting all the avaerage
	groupStage := bson.D{
	{"$group", bson.D{
		{"_id", "$category"},
		{"average_price", bson.D{{"$avg", "$price"}}},
		{"type_total", bson.D{{"$sum", 1}}},
	}}}
	cursor, err := connect.TeaCollection.Aggregate(context.TODO(), mongo.Pipeline{groupStage})
	if err!=nil{
		panic(err)
	}

	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err!=nil{
		panic(err)
	}
	for _, result := range results {
	fmt.Printf("Average price of %v tea options: $%v \n", result["_id"], result["average_price"])
	fmt.Printf("Number of %v tea options: %v \n\n", result["_id"], result["type_total"])
	
	}
}