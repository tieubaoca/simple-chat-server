package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/tieubaoca/simple-chat-server/controllers"
	"github.com/tieubaoca/simple-chat-server/models"
	"github.com/tieubaoca/simple-chat-server/services"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var dbClient *mongo.Client

var wsClients map[string]*websocket.Conn
var upgrader websocket.Upgrader
var broadcast = make(chan models.Message)

func main() {
	r := mux.NewRouter()
	//init ws
	wsClients = make(map[string]*websocket.Conn)
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	broadcast = make(chan models.Message)

	dbClient, err := mongo.NewClient(options.Client().ApplyURI("mongodb://admin:123456@localhost:27017/admin?connectTimeoutMS=10000&authSource=admin&authMechanism=SCRAM-SHA-1"))
	if err != nil {
		log.Println(err)
	}
	err = dbClient.Connect(context.TODO())
	if err != nil {
		log.Println(err)
	}
	log.Println("Connected to MongoDB!")

	initDb(dbClient)
	services.InitDB(dbClient)

	r.HandleFunc("/ws", handleWebsocket)
	r.HandleFunc("/chat-room/id", controllers.FindChatRoomById).Methods(http.MethodGet)
	r.HandleFunc("/chat-room/member", controllers.FindChatRoomsByMember).Methods(http.MethodGet)
	r.HandleFunc("/chat-room", controllers.InsertChatRoom).Methods(http.MethodPost)

	go handleMessages()

	log.Println("Server start on 8800")

	log.Fatal(http.ListenAndServe(":8800", r))

	defer dbClient.Disconnect(context.TODO())
}

func handleWebsocket(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	log.Println(username)
	if username == "" {
		log.Println("Username is required")
		return
	}
	if _, ok := wsClients[username]; ok {
		log.Println("Username is already taken")
		return
	}
	if user, _ := services.FindUserByUsername(username); user.Username == "" {
		if _, err := services.InsertUser(models.User{Username: username}); err != nil {
			return
		}
	}
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer ws.Close()
	wsClients[username] = ws
	for {
		var msg models.Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(wsClients, username)
			break
		}
		// Send the newly received message to the broadcast channel
		broadcast <- msg
	}
}

func handleMessages() {
	for {
		// Grab the next message from the broadcast channel
		msg := <-broadcast
		// Send it out to every client that is currently connected
		chatRoom, err := services.FindChatRoomById(msg.ChatRoom)
		if err != nil {
			log.Println(err)
		}
		_, err = services.InsertMessage(msg)
		if err != nil {
			log.Println(err)
			continue
		}
		for _, member := range chatRoom.Member {
			if ws, ok := wsClients[member]; ok {
				err := ws.WriteJSON(msg)
				if err != nil {
					log.Printf("error: %v", err)
					ws.Close()
					delete(wsClients, member)
				}
			}
		}
	}
}

func initDb(client *mongo.Client) {
	result, err := client.Database("saas").Collection("user").InsertMany(context.TODO(), []interface{}{
		bson.M{"username": "baodeptrai"},
		bson.M{"username": "tieubaoca"},
		bson.M{"username": "ldphong"},
		bson.M{"username": "teddy"},
	})
	if err != nil {
		log.Println(err)
	} else {

		log.Printf("Insert user success: %s", result.InsertedIDs)
	}
	result, err = client.Database("saas").Collection("chat_room").InsertMany(context.TODO(), []interface{}{
		bson.M{
			"name": "bao_bao",
			"member": []interface{}{
				"tieubaoca",
				"baodeptrai",
			},
		},
		bson.M{
			"name": "phong_bao",
			"member": []interface{}{
				"ldphong",
				"baodeptrai",
			},
		},
		bson.M{
			"name": "tung_bao",
			"member": []interface{}{
				"tieubaoca",
				"teddy",
			},
		},
	})

	if err != nil {
		log.Println(err)
	} else {

		log.Printf("Insert chat room success: %s", result.InsertedIDs)
	}
}
