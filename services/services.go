package services

import "go.mongodb.org/mongo-driver/mongo"

var client *mongo.Client

func InitDB(_client *mongo.Client) {
	client = _client
}
