package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/tieubaoca/simple-chat-server/models"
	"github.com/tieubaoca/simple-chat-server/services"
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

func InsertChatRoom(w http.ResponseWriter, r *http.Request) {
	var chatRoom models.ChatRoom
	err := json.NewDecoder(r.Body).Decode(&chatRoom)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	result, err := services.InsertChatRoom(chatRoom)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(result)
}
