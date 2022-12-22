# simple-chat-server

## Description

A simple chat server developed by golang and Mongodb

## How to use

Custom the connection string to connect to MongoDb

```go
	dbClient, err := mongo.NewClient(options.Client().ApplyURI("mongodb://admin:123456@localhost:27017/admin?connectTimeoutMS=10000&authSource=admin&authMechanism=SCRAM-SHA-1"))
```

Install neccesary modules

```sh
    go get ./...
```

Start

```sh
    go run main.go
```

## API

```
http://localhost:8800/
```

Simple front end to send and receive message.
Submit username to connect to server web socket.
</br>
Send message to another user, if the chat room between 2 users does not exist, auto create a new
chat room and send message. Message auto broadcast to all user in the chat room.

```
    ws://localhost:8800/ws?username=[your-username]
```

Create a web socket connection to server with your username.
If user does not exist, it will create an user with your username
and insert to db.

```
http://localhost:8800/api/chat-room/id?id=[chat-room-id]
```

Get chat room by room id

```
http://localhost:8800/api/chat-room/member?member=[username]
```

Get all chat rooms by username

```
http://localhost:8800/api/chat-room/members?member1=[user1]&member2=[user2]
```

Get a chat room between two user

```
http://localhost:8800/api/chat-room
```

Post method to insert new chat room
