package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
	Id       primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	ChatRoom string             `json:"chatRoom"`
	Sender   string             `json:"sender"`
	Content  string             `json:"content"`
}
