package models

import (
	"encoding/json"
	"fmt"
	// "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"navigator/conf"
)

const (
	primaryKey         = "bookcode"
	bookCollectionName = "book"
)

type BookModel struct {
	BookCode   string `json: "bookCode"`
	BookName   string `json: "bookName"`
	Author     string `json: "author"`
	BorrowedAt string `json: "borrowedAt"`
	BorrowedBy string `json: "borrowedBy"`
}

func NewBookModel(dict string) *BookModel {
	var b BookModel
	err := json.Unmarshal([]byte(dict), &b)
	if err != nil {
		return nil
	}

	return &b
}

/*
type ModelInterface interface {
	ModelList(conditions map[string]string) []interface{}
	GetModel(modelId string) interface{}
	InsertModel(model interface{})
	SearchModelList(string)
	UpdateModel(selector map[string]string, model *interface{}) error
}
*/

func (m *BookModel) ModelList(conditions bson.M) []BookModel {
	confInstance := conf.ConfigInstance()
	mgoUri := conf.ConfigInstance().MongoDBConnectionString()

	session, err := confInstance.GetMgoSession(mgoUri)
	if err != nil {
		fmt.Printf("can not connect to server %v\n", err)
		panic(err)
	}
	defer session.Close()

	var bookModelList []BookModel
	collection := session.DB(confInstance.DBName).C(bookCollectionName)
	err = collection.Find(conditions).All(&bookModelList)

	return bookModelList
}

func (m *BookModel) GetModel(modelId string) (*BookModel, error) {
	confInstance := conf.ConfigInstance()
	mgoUri := conf.ConfigInstance().MongoDBConnectionString()

	var bookModel BookModel
	session, err := confInstance.GetMgoSession(mgoUri)
	if err != nil {
		fmt.Printf("can not connect to server %v\n", err)
		panic(err)
	}
	defer session.Close()

	collection := session.DB(confInstance.DBName).C(bookCollectionName)
	err = collection.Find(bson.M{primaryKey: modelId}).One(&bookModel)
	if err != nil {
		fmt.Printf("insert error - %v\n", err)
		return &BookModel{}, err
	}

	return &bookModel, nil
}

func (m *BookModel) InsertModel(model *BookModel) error {
	confInstance := conf.ConfigInstance()
	mgoUri := confInstance.MongoDBConnectionString()
	session, err := confInstance.GetMgoSession(mgoUri)
	if err != nil {
		fmt.Printf("can not connect to server %v\n", err)
		panic(err)
	}
	defer session.Close()

	f := session.DB(confInstance.DBName).C(bookCollectionName)
	err = f.Insert(model)

	return err
}

func (m *BookModel) UpdateModel(selector bson.M, model *BookModel) error {
	confInstance := conf.ConfigInstance()
	mgoUri := confInstance.MongoDBConnectionString()

	session, err := confInstance.GetMgoSession(mgoUri)
	if err != nil {
		fmt.Printf("can not connect to server %v\n", err)
		panic(err)
	}
	defer session.Close()

	f := session.DB(confInstance.DBName).C(bookCollectionName)
	err = f.Update(selector, model)

	return err
}
