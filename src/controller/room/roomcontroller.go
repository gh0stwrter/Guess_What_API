package roomcontroller

import (
	modelroom "app/src/model/room"

	"encoding/json"
	"net/http"

	"github.com/gorilla/sessions"
	"gopkg.in/mgo.v2/bson"
)

var (
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

type Room struct {
	IDRoom  bson.ObjectId `json:",omitempty"`
	AdminID string        `json:",omitempty"`
	Name    string        `json:",omitempty"`
	Data    []Room        `json:",omitempty"`
	Message string        `json",omitempty"`
}

var room Room
var rooms []Room

func CreateRoom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	json.NewDecoder(r.Body).Decode(&room)
	modelroom.CreateRoom(room.AdminID, room.Name)

	data := Room{
		IDRoom: room.IDRoom,
		Name:   room.Name,
	}

	json.NewEncoder(w).Encode(data)
}

func GetAllRoom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	
	roomsFind := modelroom.FindAllRooms()
	json.NewEncoder(w).Encode(&roomsFind)

}
