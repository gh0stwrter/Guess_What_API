package roomcontroller

import (
	modelroom "app/src/model/room"

	"encoding/json"
	"fmt"
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
}

var room Room

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
	session, _ := store.Get(r, "session-Auth")
	isAuth := session.Values["authenticated"] == true
	if isAuth {
		roomsFind := modelroom.FindAllRooms()
		fmt.Println(roomsFind)
	}

	fmt.Println(session)

}
