package models

import (
	"encoding/json"
	"fmt"
	// "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"navigator/conf"
)

const (
	borrowCollectionName = "rule"
)

type BorrowModel struct {
	BorrowedByWho string `json: "borrowedByWho"`
	BorrowedBook  string `json: "borrowedBook"`
	BorrowedAt    string `json: "borrowedAt"`
	BorrowStatus  string `json: "borrowStatus"` // 0: borrowed, 1: returned, 2: over due, 3: available
}

func NewBorrowModel(dict string) *BorrowModel {
	var bo BorrowModel
	err := json.Unmarshal([]byte(dict), &bo)
	if err != nil {
		return nil
	}

	return &bo
}

func (model *BorrowModel) ModelList(conditions bson.M) []BorrowModel {
	confInstance := conf.ConfigInstance()
	mgoUri := conf.ConfigInstance().MongoDBConnectionString()

	session, err := confInstance.GetMgoSession(mgoUri)
	if err != nil {
		fmt.Printf("can not connect to server %v\n", err)
		panic(err)
	}
	defer session.Close()

	var borrowModelList []BorrowModel
	collection := session.DB(confInstance.DBName).C(borrowCollectionName)
	err = collection.Find(conditions).All(&borrowModelList)

	return borrowModelList
}

func (m *BorrowModel) GetModel(modelId string) (*BorrowModel, error) {
	confInstance := conf.ConfigInstance()
	mgoUri := conf.ConfigInstance().MongoDBConnectionString()

	var borrowModel BorrowModel
	session, err := confInstance.GetMgoSession(mgoUri)
	if err != nil {
		fmt.Printf("can not connect to server %v\n", err)
		panic(err)
	}
	defer session.Close()

	collection := session.DB(confInstance.DBName).C(borrowCollectionName)
	err = collection.Find(bson.M{primaryKey: modelId}).One(&borrowModel)
	if err != nil {
		fmt.Printf("insert error - %v\n", err)
		return &BorrowModel{}, err
	}

	return &borrowModel, nil
}

func (m *BorrowModel) InsertModel(model *BorrowModel) error {
	confInstance := conf.ConfigInstance()
	mgoUri := confInstance.MongoDBConnectionString()
	session, err := confInstance.GetMgoSession(mgoUri)
	if err != nil {
		fmt.Printf("can not connect to server %v\n", err)
		panic(err)
	}
	defer session.Close()

	f := session.DB(confInstance.DBName).C(borrowCollectionName)
	err = f.Insert(model)

	return err
}

func (m *BorrowModel) UpdateModel(selector bson.M, model *BorrowModel) error {
	confInstance := conf.ConfigInstance()
	mgoUri := confInstance.MongoDBConnectionString()

	session, err := confInstance.GetMgoSession(mgoUri)
	if err != nil {
		fmt.Printf("can not connect to server %v\n", err)
		panic(err)
	}
	defer session.Close()

	f := session.DB(confInstance.DBName).C(borrowCollectionName)
	err = f.Update(selector, model)

	return err
}
