package middlewares

import (
	"gopkg.in/mgo.v2"
	"github.com/gin-gonic/gin"
)

var (
	SESSION *mgo.Session
	DB      *mgo.Database
)

func Monger() gin.HandlerFunc {
	return func(c *gin.Context) {
		s := SESSION.Clone()
		defer s.Close()

		s.SetMode(mgo.Monotonic, false)
		c.Set("mongo", DB)
		c.Next()
	}
}
