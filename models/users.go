package models

import "gopkg.in/mgo.v2/bson"

const (
	CollectionUser = "users"
)

type Address struct {
	Street string `json:"street" bosn:"street"`
	City   string `json:"city" bson:"city"`
	Planet string `json:"planet" bosn:"planet"`
	Phone  string `json:"phone" bosn:"phone"`
}

type User struct {
	Id        bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	FirstName string        `json:"fname" form:"fname" binding:"required" bson:"fname"`
	LastName  string        `json:"lname" form:"lname" binding:"required" bson:"lname"`
	Age       int           `json:"age" form:"age" binding:"required" bson:"age"`
	Email     string        `json:"email" binding:"required" bson:"email"`
	Address   []Address     `json:"address" bson:"address"`
}
