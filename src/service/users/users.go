package users

import (
	"go-exercise/src/models"

	GLOBAL "go-exercise/src/global"

	"github.com/globalsign/mgo/bson"
)

func List(query *GLOBAL.SafeDict) []models.User {
	db := GLOBAL.MongoSession
	var users []models.User

	filter := bson.M{}
	if lname, ok := query.Get("lname"); ok {
		filter = bson.M{
			"lname": lname,
		}
	}

	err := db.C(models.CollectionUser).Find(filter).All(&users)
	if err != nil {
		panic(err.Error())
	}
	return users
}

func CreateOne(data UserSchema) bool {
	db := GLOBAL.MongoSession
	user := map[string]interface{}{
		"fname":   data.FirstName,
		"lname":   data.LastName,
		"age":     data.Age,
		"email":   data.Email,
		"address": data.Address,
	}
	err := db.C(models.CollectionUser).Insert(&user)
	if err != nil {
		panic(err.Error())
	}
	return true
}

func Detail(id string) models.User {
	db := GLOBAL.MongoSession
	var user models.User
	bsonId := bson.ObjectIdHex(id)
	err := db.C(models.CollectionUser).Find(bson.M{"_id": bsonId}).One(&user)
	if err != nil {
		panic(err.Error())
	}
	return user
}

func Edit(id string, data interface{}) bool {
	db := GLOBAL.MongoSession
	bsonId := bson.ObjectIdHex(id)
	doc := bson.M{
		"$set": bson.M{
			"age":   44,
			"email": "test@gmail.com",
		},
	}
	err := db.C(models.CollectionUser).Update(bson.M{"_id": bsonId}, doc)
	if err != nil {
		panic(err.Error())
	}
	return true
}

func Remove(id string) bool {
	db := GLOBAL.MongoSession
	bsonId := bson.ObjectIdHex(id)

	err := db.C(models.CollectionUser).Remove(bson.M{"_id": bsonId})
	if err != nil {
		panic(err.Error())
	}
	return true
}
