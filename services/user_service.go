package services

import (
	"context"

	"github.com/tieubaoca/simple-chat-server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindUserByUsername(username string) (models.User, error) {
	db := client.Database("saas").Collection("user")
	var result models.User
	err := db.FindOne(context.TODO(), bson.D{{"username", username}}).Decode(&result)
	return result, err
}

func InsertUser(user models.User) (*mongo.InsertOneResult, error) {
	db := client.Database("saas").Collection("user")
	return db.InsertOne(context.TODO(), user)

}
