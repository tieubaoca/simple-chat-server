package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ChatRoom struct {
	Id     primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Name   string             `json:"name"`
	Member []string           `json:"member"`
}
