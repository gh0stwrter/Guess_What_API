package modelroom

import (
	model "app/src/model"
	"fmt"
	"gopkg.in/mgo.v2/bson"

	"upper.io/db.v3"
)

var orm = model.DatabaseSession()
var room Room
var rooms []Room
var roomCollection = orm.Collection("room")

type Room struct {
	Admin   string   `bson:"_admin"`
	Players []string `bson:"players"`
	Name    string   `bson:"name"`
	Turn    string   `bson:"turn"`
}

func RoomCreate(admin string, name string) string {

	fmt.Println(roomCollection.Insert(Room{
		Admin: admin,
		Name:  name,
	}))

	return "Room Create"
}
func UserJoined(name string, id string) {
	data := bson.ObjectIdHex(id)
	res := roomCollection.Find(db.Cond{"_id": data})
	err := res.One(&room)
	if err != nil {
		fmt.Println(err)
	}
	player := []string{name}
	players := append(player, name)
	res.Update(Room{
		Players: players,
	})
}

func FindAllRooms() []Room {
	res := roomCollection.Find()
	res.All(&rooms)
	return rooms
}
