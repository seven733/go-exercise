package users

import (
	"github.com/gin-gonic/gin"
	"github.com/seven/demo/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func List(c *gin.Context) []models.User {
	db := c.MustGet("db").(*mgo.Database)
	users := []models.User{}

	err := db.C(models.CollectionUser).Find(bson.M{}).All(&users)
	if err != nil {
		c.Error(err)
	}
	return users
}

func CreateOne(c *gin.Context, data User) bool {
	db := c.MustGet("db").(*mgo.Database)
	user := map[string]interface{}{
		"fname":   data.FirstName,
		"lname":   data.LastName,
		"age":     data.Age,
		"email":   data.Email,
		"address": data.Address,
	}
	err := db.C(models.CollectionUser).Insert(&user)
	if err != nil {
		c.Error(err)
	}
	return true
}
