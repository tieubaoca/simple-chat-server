package services

import (
	"context"

	"github.com/tieubaoca/simple-chat-server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindChatRoomById(id string) (models.ChatRoom, error) {
	db := client.Database("saas").Collection("chat_room")
	var result models.ChatRoom
	obId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return result, err
	}
	err = db.FindOne(context.TODO(), bson.D{{"_id", obId}}).Decode(&result)
	return result, err
}

func FindChatRoomsByMember(member string) ([]models.ChatRoom, error) {
	db := client.Database("saas").Collection("chat_room")
	cursor, err := db.Find(context.TODO(), bson.D{{"member", member}})
	if err != nil {
		return nil, err
	}

	var results []models.ChatRoom
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}
	return results, err
}

func FindChatRoomByMembers(members []string) (models.ChatRoom, error) {
	db := client.Database("saas").Collection("chat_room")
	var result models.ChatRoom
	err := db.FindOne(context.TODO(), bson.D{{"member", bson.D{
		{"$all", members},
	}}}).Decode(&result)
	return result, err
}

func InsertChatRoom(chatRoom models.ChatRoom) (*mongo.InsertOneResult, error) {
	db := client.Database("saas").Collection("chat_room")
	return db.InsertOne(context.TODO(), chatRoom)
}
