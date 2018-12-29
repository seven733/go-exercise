package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/seven/demo/db"
)

func PrintTest(c *gin.Context) {
	// 请求执行前
	fmt.Println("begin....")
	c.Next()
}

// Connect middleware clones the database session for each request and
// makes the `db` object available for each handler
func Connect(c *gin.Context) {
	s := db.Session.Clone()

	defer s.Close()

	c.Set("db", s.DB(db.Mongo.Database))
	c.Next()
}
