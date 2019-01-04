package users

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/seven/demo/models"
	userService "github.com/seven/demo/service/users"
	validator "gopkg.in/go-playground/validator.v9"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var validate *validator.Validate

func List(c *gin.Context) {
	// var users = userService.List(c)
	db := c.MustGet("db").(*mgo.Database)
	fmt.Println("db", db)
	users := []models.User{}

	err := db.C(models.CollectionUser).Find(bson.M{}).All(&users)
	if err != nil {
		c.Error(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func Create(c *gin.Context) {

	// body, _ := ioutil.ReadAll(c.Request.Body)
	// fmt.Println("---body/--- \r\n " + string(body))

	var user userService.User
	c.BindJSON(&user)
	validate = validator.New()
	err := validate.Struct(user)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		fmt.Println("validationErrors", validationErrors)
		if len(validationErrors) > 0 {
			err := validationErrors[0]
			c.String(http.StatusBadRequest, err.Namespace()+" validation failed")
		}
		return
	}
	userService.CreateOne(c, user)
	c.String(http.StatusOK, "success")
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
