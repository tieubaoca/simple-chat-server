package services

import (
	"context"

	"github.com/tieubaoca/simple-chat-server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindMessagesByChatRoomId(chatRoomId string) ([]models.Message, error) {
	db := client.Database("saas").Collection("message")
	ojId, err := primitive.ObjectIDFromHex(chatRoomId)
	if err != nil {
		return nil, err
	}
	result, err := db.Find(context.TODO(), bson.D{{"chatRoom", ojId}})
	if err != nil {
		return nil, err
	}
	var messages []models.Message
	if err = result.All(context.TODO(), &messages); err != nil {
		return nil, err
	}
	return messages, nil
}

func InsertMessage(message models.Message) (*mongo.InsertOneResult, error) {
	db := client.Database("saas").Collection("message")
	return db.InsertOne(context.TODO(), message)
}
