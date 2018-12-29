package db

import (
	"fmt"

	mgo "gopkg.in/mgo.v2"
)

var (
	// Session stores mongo session
	Session *mgo.Session

	// Mongo stores the mongodb connection string informatio
	Mongo *mgo.DialInfo
)

const MongodbUrl = "mongodb://localhost:27017/go-demo"

func Connect() {
	mongo, err := mgo.ParseURL(MongodbUrl)
	s, err := mgo.Dial(MongodbUrl)
	if err != nil {
		fmt.Println("can't contect mongo", err)
		panic(err.Error())
	}
	s.SetSafe(&mgo.Safe{})
	fmt.Println("Connected to", MongodbUrl)
	Session = s
	Mongo = mongo
}
