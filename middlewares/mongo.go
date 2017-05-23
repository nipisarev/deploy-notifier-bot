package middlewares

import (
	"gopkg.in/mgo.v2"
	"github.com/gin-gonic/gin"
)

var SESSION *mgo.Session

func Monger() gin.HandlerFunc {

	return func(c *gin.Context) {
		s := SESSION.Clone()
		defer s.Close()

		c.Set("db", s.DB("deployment"))
		c.Next()
	}
}
