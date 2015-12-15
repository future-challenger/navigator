package models

import (
	// "encoding/json"
	// "fmt"
	// "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	// "navigator/conf"
)

type ModelInterface interface {
	ModelList(conditions bson.M) []interface{}
	GetModel(modelId string) interface{}
	InsertModel(model interface{})
	SearchModelList(string)
	UpdateModel(selector map[string]string, model *interface{}) error
}
