package models

import "gopkg.in/mgo.v2/bson"

type Issue struct {
	Id bson.ObjectId `json:"id" bson:"_id"`
	UpdateStatus bool `json:"update_status" bson:"update_status"`
	UpdateLabel bool `json:"update_label" bson:"update_label"`
}
