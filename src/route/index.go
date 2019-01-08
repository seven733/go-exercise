package route

import (
	"fmt"
	"io"
	"net/http"
	"os"

	userController "go-exercise/src/controller/users"

	"go-exercise/src/middlewares"

	"go-exercise/src/service/articles"

	"github.com/gin-gonic/gin"
)

func InitRoute() http.Handler {

	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	// 将日志写入文件
	fmt.Fprintln(gin.DefaultWriter, "add record")

	r := gin.Default()
	r.Use(middlewares.PrintTest)
	r.Use(middlewares.Connect)
	r.Use(middlewares.HandleResponses)

	api := r.Group("/api")
	{
		api.GET("/articles", articles.List)
		api.POST("/articles", articles.CreateOne)

		api.GET("/users", userController.List)
		api.POST("/users", userController.Create)
		api.GET("/users/:userId", userController.Detail)
		api.DELETE("/users/:userId", userController.Remove)
		api.PUT("/users/:userId", userController.Edit)
	}

	return r
}
