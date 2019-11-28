package modelroom

import (
	model "app/src/model"
	"gopkg.in/mgo.v2/bson"
)

var orm = model.DatabaseSession()
var room Room
var rooms []Room
var roomCollection = orm.Collection("room")

type User struct {
	Name string
}
type Room struct {
	IDRoom  bson.ObjectId `bson:"_id"`
	Admin   string        `bson:"_admin"`
	Players []User        `bson:"players"`
	Turn    string        `bson:"turn"`
}

func CreateRoom(admin string, name string) string {

	roomCollection.Insert(Room{
		Admin: admin,
	})

	return "Room Create"
}

func FindAllRooms() []Room {
	res := roomCollection.Find()
	res.All(&rooms)
	return rooms
}
