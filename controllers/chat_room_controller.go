package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/tieubaoca/simple-chat-server/models"
	"github.com/tieubaoca/simple-chat-server/services"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindChatRoomById(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	log.Println("FindChatRoomById: " + id)
	w.Header().Set("Content-Type", "application/json")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	chatRoom, err := services.FindChatRoomById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(chatRoom)
}

func FindChatRoomsByMember(w http.ResponseWriter, r *http.Request) {
	member := r.URL.Query().Get("member")
	w.Header().Set("Content-Type", "application/json")
	if member == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	chatRooms, err := services.FindChatRoomsByMember(member)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(chatRooms)
}

func FindChatRoomByMembers(w http.ResponseWriter, r *http.Request) {
	member1 := r.URL.Query().Get("member1")
	member2 := r.URL.Query().Get("member2")
	w.Header().Set("Content-Type", "application/json")
	if member1 == "" || member2 == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	chatRoom, err := services.FindChatRoomByMembers([]string{member1, member2})
	if err != nil && err != mongo.ErrNoDocuments {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	json.NewEncoder(w).Encode(chatRoom)
}

func InsertChatRoom(w http.ResponseWriter, r *http.Request) {
	var chatRoom models.ChatRoom
	err := json.NewDecoder(r.Body).Decode(&chatRoom)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	result, err := services.InsertChatRoom(
		bson.M{
			"name":   chatRoom.Name,
			"member": chatRoom.Member,
		},
	)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(result)
}
