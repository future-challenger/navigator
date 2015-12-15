package models

import (
// "encoding/json"
// "fmt"
// "gopkg.in/mgo.v2"
// "gopkg.in/mgo.v2/bson"
// "navigator/conf"
)

const (
// primaryKey = "rulecode"
)

type RuleModel struct {
	RuleCode string `json: ruleCode`
	TimeSpan string `json: "timeSpan"`
}
