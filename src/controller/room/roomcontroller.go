package roomcontroller

import (
	modelroom "app/src/model/room"

	"encoding/json"
	"net/http"

	"github.com/gorilla/sessions"
)

var (
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

type Room struct {
	Room  string
	Admin string `json:",omitempty"`
	Name  string `json:",omitempty"`
}

var room Room
var rooms []Room

func CreateRoom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	json.NewDecoder(r.Body).Decode(&room)
	modelroom.RoomCreate(room.Admin, room.Name)
	data := Room{
		Name: room.Name,
	}

	json.NewEncoder(w).Encode(data)
}

func JoinedUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	json.NewDecoder(r.Body).Decode(&room)
	modelroom.UserJoined(room.Name, room.Room)
	data := Room{
		Room: room.Room,
		Name: room.Name,
	}

	json.NewEncoder(w).Encode(data)
}

func GetAllRoom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")

	roomsFind := modelroom.FindAllRooms()
	json.NewEncoder(w).Encode(&roomsFind)

}
