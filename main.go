package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Product is a struct
type Product struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	Name        string             `json:"product_name" bson:"product_name"`
	Price       float64            `json:"product_price" bson:"product_price"`
	Currency    string             `json:"currency" bson:"currency"`
	Quantity    string             `json:"product_quantity" bson:"product_quantity"`
	Discount    int                `json:"product_discount,omitempty" bson:"product_discount,omitempty"`
	Vendor      string             `json:"product_vendor" bson:"product_vendor"`
	Accessoires []string           `json:"product_accessoires,omitempty" bson:"product_accessoires,omitempty"`
	SKuID       string             `json:"SKuID" bson:"SKuID"`
}

var oppo = Product{
	ID:          primitive.NewObjectID(),
	Name:        "Oppo F5",
	Price:       1592.99,
	Currency:    "MAD",
	Quantity:    "50",
	Vendor:      "OPPO Inc.",
	Accessoires: []string{"Cable", "Showcase", "Charger"},
	SKuID:       "12",
}

var iPhone12 = Product{
	ID:          primitive.NewObjectID(),
	Name:        "iPhone 12 pro max",
	Price:       15000.99,
	Currency:    "MAD",
	Quantity:    "50",
	Discount:    10,
	Vendor:      "Appel Inc.",
	Accessoires: []string{"Cable", "Showcase", "Charger"},
	SKuID:       "12",
}

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		fmt.Println(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		fmt.Println(err)
	}
	db := client.Database("tronics")
	collection := db.Collection("products")
	res, err := collection.InsertOne(context.Background(), iPhone12)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.InsertedID.(primitive.ObjectID).Timestamp())
}
