package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func HandleResponses(c *gin.Context) {

	c.Next()
	// 请求完成后执行你的代码
	fmt.Printf("end,,,,,,,,,,,\n")
}
