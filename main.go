package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// import "golang-echo/app"

// Product describes an electronic product e.g. phone
type Product struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	Name        string             `json:"product_name" bson:"product_name"`
	Price       int                `json:"price" bson:"price"`
	Currency    string             `json:"currency" bson:"currency"`
	Quantity    string             `json:"quantity" bson:"quantity"`
	Discount    int                `json:"discount,omitempty" bson:"discount,omitempty"`
	Vendor      string             `json:"vendor" bson:"vendor"`
	Accessories []string           `json:"accessories,omitempty" bson:"accessories,omitempty"`
	SkuId       string             `json:"sku_id" bson:"sku_id"`
}

var iphone10 = Product{
	ID:          primitive.NewObjectID(),
	Name:        "iphone10",
	Price:       900,
	Currency:    "USD",
	Quantity:    "40",
	Vendor:      "apple",
	Accessories: []string{"charger", "headset", "slotopener"},
	SkuId:       "1234",
}

func main() {
	// app.Start()

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		fmt.Println(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		fmt.Println(err)
	}

	db := client.Database("golangEcho")
	collection := db.Collection("products")
	// res, err := collection.InsertOne(context.Background(), iphone10)
	res, err := collection.InsertOne(context.Background(), bson.D{
		{"name", "eric"},
		{"surname", "cartman"},
		{"hobbies", bson.A{"videogame", "alexa", "kfc"}},
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res.InsertedID.(primitive.ObjectID).Timestamp())
}
