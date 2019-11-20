package main

import (
	modeluser "app/src/model/user"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan Message)
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Message struct {
	Username string `json:"username"`
}
type User struct {
	Name     string `json:",omitempty"`
	Password string `json:",omitempty"`
	Token    string `json:",omitempty"`
	Message  string `json:",omitempty"`
	Score    int    `json:",omitempty"`
}

var user User

func SignUpSignIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	json.NewDecoder(r.Body).Decode(&user)
	token := modeluser.Login(user.Name, user.Password)

	data := User{Token: token,
		Message: "Token successful",
	}
	json.NewEncoder(w).Encode(data)
}

func ScoreData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	json.NewDecoder(r.Body).Decode(&user)
	fmt.Println(user)
	res := modeluser.SetScore(user.Score, "5dd48dcba54d75554f479a0a")
	fmt.Println(res)
	data := User{
		Score: user.Score,
	}
	fmt.Println(data)
	json.NewEncoder(w).Encode(data)

}
func main() {
	// Create a simple file server
	r := mux.NewRouter()
	http.Handle("/", r)
	r.HandleFunc("/ws", handleConnections)
	r.HandleFunc("/score", ScoreData)
	r.HandleFunc("/sign-in", SignUpSignIn)
	log.Println("http server started on :8000")

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	defer ws.Close()

	clients[ws] = true

	for {
		var msg Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("erro: %v", err)
			delete(clients, ws)
			break
		}
		broadcast <- msg
	}
}
