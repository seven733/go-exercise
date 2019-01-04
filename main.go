package main

import (
	"fmt"
	"net/http"
	"time"

	"go-exercise/db"

	"go-exercise/route"

	"github.com/gin-gonic/gin"
	validator "gopkg.in/go-playground/validator.v9"
)

var map_db = map[string]string{
	"foo": "bar",
}

var validate *validator.Validate

type Address struct {
	Street string `validate:"required"`
	City   string `validate:"required"`
	Planet string `validate:"required"`
	Phone  string `validate:"required"`
}
type User struct {
	FirstName string     `json:"fname"`
	LastName  string     `json:"lname"`
	Age       int        `json:"age"validate:"gte=0,lte=130"`
	Email     string     `json:"email"validate:"required,email"`
	Addresses []*Address `json:"address"validate:"required,dive,required"`
}

var userDB = make([]User, 0)

func pong(c *gin.Context) {
	c.JSON(200, gin.H{"ping": "pong"})
}

func getUsers(c *gin.Context) {
	c.JSON(200, map_db)
}

func getUserById(c *gin.Context) {
	name := c.Param("name")
	user, ok := map_db[name]
	fmt.Println("user", user)
	if ok {
		c.String(200, user)
	}
}

func userAction(c *gin.Context) {
	name := c.Param("name")
	action := c.Param("action")
	c.String(200, name+" is "+action)
}

func welcome(c *gin.Context) {
	firstname := c.DefaultQuery("firstname", "Guest")
	lastname := c.Query("lastname")
	c.String(200, "Hello %s %s", firstname, lastname)
}

func createUser(c *gin.Context) {
	var user User
	c.BindJSON(&user)
	validate = validator.New()
	err := validate.Struct(user)
	if err != nil {
		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return
		}

		validationErrors := err.(validator.ValidationErrors)
		fmt.Println("validationErrors", validationErrors)
		if len(validationErrors) > 0 {
			err := validationErrors[0]
			c.String(http.StatusBadRequest, err.Namespace()+" validation failed")
		}
		return
	}
	userDB = append(userDB, user)
	c.JSON(200, userDB)
}

type Login struct {
	User     string `json:"user" binding:"required"`
	Password string `json:"password"  binding:"required"`
}

func loginEndpoint(c *gin.Context) {
	var json Login
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if json.User != "jack" || json.Password != "123" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "password error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "logged success"})
}

func init() {
	db.Connect()
}

func main() {
	server01 := &http.Server{
		Addr:         ":2333",
		Handler:      route.InitRoute(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server01.ListenAndServe()
}
