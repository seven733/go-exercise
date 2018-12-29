package users

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/seven/demo/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func List(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	users := []models.User{}

	err := db.C(models.CollectionUser).Find(nil).All(&users)
	if err != nil {
		c.Error(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func CreateOne(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	user := models.User{}
	err := c.Bind(&user)
	if err != nil {
		c.Error(err)
		c.String(http.StatusBadRequest, "user schema error")
		return
	}

	fmt.Println("user", user)
	err = db.C(models.CollectionUser).Insert(user)
	if err != nil {
		c.Error(err)
	}
	c.String(http.StatusCreated, "create user success")
}

func Detail(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	var user models.User
	id := bson.ObjectIdHex(c.Param("userId"))
	err := db.C(models.CollectionUser).Find(bson.M{"_id": id}).One(&user)
	if err != nil {
		c.Error(err)
	}
	c.JSON(http.StatusOK, user)
}

func Edit(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	id := bson.ObjectIdHex(c.Param("userId"))
	doc := bson.M{
		"$set": bson.M{
			"age":   44,
			"email": "test@gmail.com",
		},
	}
	err := db.C(models.CollectionUser).Update(bson.M{"_id": id}, doc)
	if err != nil {
		c.Error(err)
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}

func Remove(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	id := bson.ObjectIdHex(c.Param("userId"))

	err := db.C(models.CollectionUser).Remove(bson.M{"_id": id})
	if err != nil {
		c.Error(err)
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}
