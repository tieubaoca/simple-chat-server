package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/tieubaoca/simple-chat-server/services"
)

func FindMessageByChatRoom(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	w.Header().Set("Content-Type", "application/json")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	messages, err := services.FindMessagesByChatRoomId(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(messages)
}
