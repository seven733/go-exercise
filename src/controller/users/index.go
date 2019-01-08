package users

import (
	"net/http"

	GLOBAL "go-exercise/src/global"
	userService "go-exercise/src/service/users"

	utils "go-exercise/src/utils"

	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	var query *GLOBAL.SafeDict = GLOBAL.NewSafeDict(map[string]interface{}{})

	query.Put("lname", c.Query("lname"))
	users := userService.List(query)
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func Create(c *gin.Context) {
	var user userService.UserSchema
	c.BindJSON(&user)
	field, err := utils.ValidateBody(user)
	if err != nil {
		c.String(400, field+" 校验失败")
		return
	}
	userService.CreateOne(user)
	c.String(http.StatusOK, "success")
}

func Detail(c *gin.Context) {
	var id = c.Param("userId")
	user := userService.Detail(id)
	c.JSON(http.StatusOK, user)
}

func Edit(c *gin.Context) {
	var id = c.Param("userId")
	var result bool = userService.Edit(id, nil)
	c.JSON(http.StatusOK, gin.H{"success": result})
}

func Remove(c *gin.Context) {
	var id = c.Param("userId")
	var result = userService.Remove(id)
	c.JSON(http.StatusOK, gin.H{"success": result})
}
