package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Host struct {
	Id bson.ObjectId `json:"id" bson:"_id"`
	Name string `json:"name" bson:"name"`
	ProjectName string `json:"project_name" bson:"project_name"`
	RepoName string `json:"repo_name" bson:"repo_name"`
	Hipchat Chat `json:"hipchat" bson:"hipchat"`
	Slack Chat `json:"slack" bson:"slack"`
	Jira Issue `json:"jira" bson:"jira"`
}