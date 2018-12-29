package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/seven/demo/db"
	"github.com/seven/demo/middlewares"
	articles "github.com/seven/demo/service/articles"
	"github.com/seven/demo/service/users"
	"golang.org/x/sync/errgroup"
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

var (
	g errgroup.Group
)

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

func router() http.Handler {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	// 将日志写入文件
	fmt.Fprintln(gin.DefaultWriter, "add record")

	r := gin.Default()
	r.Use(middlewares.PrintTest)
	r.Use(middlewares.Connect)
	r.Use(middlewares.HandleResponses)

	r.GET("/ping", pong)
	r.GET("/users", getUsers)
	r.GET("/users/:name", getUserById)
	r.GET("/users/:name/*action", userAction)
	r.GET("/welcome", welcome)

	r.POST("/users", createUser)

	v1 := r.Group("/v1")
	{
		v1.POST("/login", loginEndpoint)
	}

	v2 := r.Group("/v2")
	{
		v2.GET("/articles", articles.List)
		v2.POST("/articles", articles.CreateOne)

		v2.GET("/users", users.List)
		v2.POST("/users", users.CreateOne)
		v2.GET("/users/:userId", users.Detail)
		v2.DELETE("/users/:userId", users.Remove)
		v2.PUT("/users/:userId", users.Edit)
	}

	return r
}

func init() {
	db.Connect()
}

func main() {
	server01 := &http.Server{
		Addr:         ":2333",
		Handler:      router(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server02 := &http.Server{
		Addr:         ":2334",
		Handler:      router(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	g.Go(func() error {
		return server01.ListenAndServe()
	})

	g.Go(func() error {
		return server02.ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
