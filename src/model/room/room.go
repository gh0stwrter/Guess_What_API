package modelroom

import (
	model "app/src/model"
	"gopkg.in/mgo.v2/bson"
)

var orm = model.DatabaseSession()
var room Room

type Room struct {
	IDRoom bson.ObjectId `bson:"_id"`
	Admin  bson.ObjectId `bson:"admin"`
	Name   string        `bson:"name"`
}

func CreateRoom(admin bson.ObjectId, name string) string {
	roomCollection := orm.Collection("room")

	roomCollection.Insert(Room{
		Admin: admin,
		Name:  name,
	})
	return "Room Create"
}
