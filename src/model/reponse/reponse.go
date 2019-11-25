package reponse

import (
	model "app/src/model"
	"gopkg.in/mgo.v2/bson"
	"fmt"
)

var reponse Reponse
var orm = model.DatabaseSession()

type Reponse struct {
	IDReponse bson.ObjectId `bson:"_id"`
	Name string `bson:"name"`
}

func GetRandomReponse() {
	responseCollection := orm.Collection("reponse")
	res := responseCollection.Find()
	res.All(&reponse)
	fmt.Println(reponse)
}
