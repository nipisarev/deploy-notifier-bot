package main

import (
	"net/http"
	"github.com/nipisarev/deploy-notifier-bot/models"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func Index(c *gin.Context) {
	c.String(http.StatusOK, "welcome")
}

func HostIndex(c *gin.Context) {
	c.Header("Content-Type", "application/json; charset=UTF-8")
	db := getDB(c)
	results := []models.Host{}

	if err := db.C("hosts").Find(nil).All(&results); err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	c.JSON(http.StatusOK, results)
}

func GetHost(c *gin.Context) {
	c.Header("Content-Type", "application/json; charset=UTF-8")

	name := c.Param("name")
	h := models.Host{}
	db := getDB(c)
	if err := db.C("hosts").Find(bson.M{"name": name}).One(&h); err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}
	c.JSON(http.StatusOK, h)
}

func AddHost(c *gin.Context) {
	var host models.Host
	c.BindJSON(&host)

	host.Id = bson.NewObjectId()
	host.Hipchat.Id = bson.NewObjectId()
	host.Slack.Id = bson.NewObjectId()
	host.Jira.Id = bson.NewObjectId()

	db := getDB(c)
	err := db.C("hosts").Insert(&host)
	if err != nil {
		Must(err)
	}
	count,_ := db.C("hosts").Count()
	c.Header("Content-Type", "application/json; charset=UTF-8")
	c.JSON(http.StatusCreated, count)
}

func getDB(c *gin.Context) *mgo.Database {
	return c.MustGet("mongo").(*mgo.Database)
}
