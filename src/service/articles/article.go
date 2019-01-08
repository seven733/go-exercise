package articles

import (
	"fmt"
	"net/http"

	"go-exercise/src/models"

	"github.com/gin-gonic/gin"
	mgo "github.com/globalsign/mgo"
)

func CreateOne(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	article := models.Article{}
	err := c.Bind(&article)
	if err != nil {
		c.Error(err)
		c.String(http.StatusBadRequest, "schema error")
	}

	fmt.Println("article", article)
	err = db.C(models.CollectionArticle).Insert(article)
	if err != nil {
		c.Error(err)
	}
	c.String(http.StatusCreated, "success")
}

func List(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	articles := []models.Article{}

	err := db.C(models.CollectionArticle).Find(nil).All(&articles)
	if err != nil {
		c.Error(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"articles": articles,
	})
}
