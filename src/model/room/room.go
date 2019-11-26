package modelroom

import (
	model "app/src/model"
	modeluser "app/src/model/user"

)

var orm = model.DatabaseSession()
var room Room
var rooms []Room
var roomCollection = orm.Collection("room")

type Room struct {
	Admin  string `bson:"_admin"`
	Name   string `bson:"name"`
	Enable bool   `bson:enable`
}



func CreateRoom(admin string, name string) string {

	roomCollection.Insert(Room{
		Admin: admin,
		Name:  name,
	})

	return "Room Create"
}

func FindAllRooms() []Room {
	res := roomCollection.Find()
	res.All(&rooms)
	return rooms
}
