package models

import (
	"encoding/json"
	"fmt"
	// "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"navigator/conf"
)

const (
	ruleCollectionName = "rule"
)

type RuleModel struct {
	RuleCode  string `json: "ruleCode"`
	TimeSpan  string `json: "timeSpan"`
	BookLimit string `json: "bookLimit"`
	CreatedAt string `json: "createdAt"`
	UpdatedAt string `json: "updatedAt"`
}

func NewRuleModel(dict string) *RuleModel {
	var r RuleModel
	err := json.Unmarshal([]byte(dict), &r)
	if err != nil {
		return nil
	}

	return &r
}

func (model *RuleModel) ModelList(conditions bson.M) []RuleModel {
	confInstance := conf.ConfigInstance()
	mgoUri := conf.ConfigInstance().MongoDBConnectionString()

	session, err := confInstance.GetMgoSession(mgoUri)
	if err != nil {
		fmt.Printf("can not connect to server %v\n", err)
		panic(err)
	}
	defer session.Close()

	var ruleModelList []RuleModel
	collection := session.DB(confInstance.DBName).C(ruleCollectionName)
	err = collection.Find(conditions).All(&ruleModelList)

	return ruleModelList
}

func (m *RuleModel) GetModel(modelId string) (*RuleModel, error) {
	confInstance := conf.ConfigInstance()
	mgoUri := conf.ConfigInstance().MongoDBConnectionString()

	var ruleModel RuleModel
	session, err := confInstance.GetMgoSession(mgoUri)
	if err != nil {
		fmt.Printf("can not connect to server %v\n", err)
		panic(err)
	}
	defer session.Close()

	collection := session.DB(confInstance.DBName).C(ruleCollectionName)
	err = collection.Find(bson.M{primaryKey: modelId}).One(&ruleModel)
	if err != nil {
		fmt.Printf("insert error - %v\n", err)
		return &RuleModel{}, err
	}

	return &ruleModel, nil
}

func (m *RuleModel) InsertModel(model *RuleModel) error {
	confInstance := conf.ConfigInstance()
	mgoUri := confInstance.MongoDBConnectionString()
	session, err := confInstance.GetMgoSession(mgoUri)
	if err != nil {
		fmt.Printf("can not connect to server %v\n", err)
		panic(err)
	}
	defer session.Close()

	f := session.DB(confInstance.DBName).C(ruleCollectionName)
	err = f.Insert(model)

	return err
}

func (m *RuleModel) UpdateModel(selector bson.M, model *RuleModel) error {
	confInstance := conf.ConfigInstance()
	mgoUri := confInstance.MongoDBConnectionString()

	session, err := confInstance.GetMgoSession(mgoUri)
	if err != nil {
		fmt.Printf("can not connect to server %v\n", err)
		panic(err)
	}
	defer session.Close()

	f := session.DB(confInstance.DBName).C(ruleCollectionName)
	err = f.Update(selector, model)

	return err
}
