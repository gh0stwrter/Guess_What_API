package database

import (
	"gopkg.in/mgo.v2/bson"
	"upper.io/db.v3"
	"upper.io/db.v3/mongo"
)

type Reponse struct {
	Name string `bson:"name"`
}

type Classement struct {
	IDUser       bson.ObjectId `bson:"iduser"`
	HighestScore int           `bson:"highestscore"`
}
type Room struct {
	Name string `bson:"name"`
}

var settings = mongo.ConnectionURL{
	Database: `guess_what`,
	Host:     `127.0.0.1`,
}
var sess, err = mongo.Open(settings)

func DatabaseSession() db.Database {
	return sess
}

func SetClassement(user *Classement, id bson.ObjectId) {
	classementCollection := sess.Collection("classment")
	classementCollection.Insert(Classement{
		IDUser: id,
		// Implemente Logic count all scores of players
	})
	defer sess.Close()
}
