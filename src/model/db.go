package database

import (
	"upper.io/db.v3"
	"upper.io/db.v3/mongo"
)

var settings = mongo.ConnectionURL{
	Database: `guess_whatdb`,
	Host:     `127.0.0.1`,
}
var sess, err = mongo.Open(settings)

func DatabaseSession() db.Database {
	return sess
}
