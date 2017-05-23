package models

import "gopkg.in/mgo.v2/bson"

type Chat struct {
	Id bson.ObjectId `json:"id" bson:"_id"`
	OutputFormat string `json:"output_format" bson:"output_format"`
	Prefix string `json:"output_format" bson:"output_format"`
}