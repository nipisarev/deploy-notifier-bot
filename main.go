package main

import (
	"log"
	"gopkg.in/mgo.v2"
	"github.com/nipisarev/deploy-notifier-bot/middlewares"
)

func main() {
	router := NewRouter()
	log.Fatal(router.Run(":8181"))
}

func init() {
	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	middlewares.SESSION = session
}

func Must(err error) {
	if err != nil {
		panic(err.Error())
	}
}
