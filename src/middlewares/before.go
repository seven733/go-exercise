package middlewares

import (
	"fmt"

	"go-exercise/src/db"

	GLOBAL "go-exercise/src/global"

	"github.com/gin-gonic/gin"
)

func PrintTest(c *gin.Context) {
	// 请求执行前
	fmt.Println("begin....")
	c.Next()
}

func Connect(c *gin.Context) {
	s := db.Session.Clone()

	defer s.Close()

	GLOBAL.MongoSession = s.DB(db.Mongo.Database)
	c.Next()
}
