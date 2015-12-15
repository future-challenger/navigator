package models

import (
	"encoding/json"
	"fmt"
	// "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"navigator/conf"
)

type DeviceInfo struct {
	SystemType string `json: "systemType"`    // ios or android
	SystemVer  string `json: "systemVersion"` // system version eg, ios 9.1
	AppVer     string `json: "appVersion"`    // app version eg, 2.9.0
}

type FeedbackModel struct {
	FeedbackCode string     `json: "feedbackCode"`
	UserName     string     `json: "userName"`
	PhoneNum     string     `json: "phoneNum"`
	DeviceInfo   DeviceInfo `json: "deviceInfo"`
	ImageList    []string   `json: "imageList"`
	Desc         string     `json: "description"`
	FeedbackType string     `json: "feedbackType"` // normal feedback or a problem in use
	Other        string     `json: "other"`        // other infomation
}

func NewFeedbackModel(dict string) *FeedbackModel {
	var f FeedbackModel
	err := json.Unmarshal([]byte(dict), &f)
	if err != nil {
		return nil
	}

	return &f
}

func (m *FeedbackModel) ModelList(conditions map[string]string) []FeedbackModel {
	confInstance := conf.ConfigInstance()
	mgoUri := conf.ConfigInstance().MongoDBConnectionString()
	// var f FeedbackModel
	session, err := confInstance.GetMgoSession(mgoUri)
	if err != nil {
		fmt.Printf("can not connect to server %v\n", err)
		panic(err)
	}
	defer session.Close()

	var feedbackModelList []FeedbackModel

	collection := session.DB(confInstance.DBName).C("feedback")
	err = collection.Find(conditions).All(&feedbackModelList)

	return feedbackModelList
}

func (m *FeedbackModel) GetModel(modelId string) (*FeedbackModel, error) {
	confInstance := conf.ConfigInstance()
	mgoUri := conf.ConfigInstance().MongoDBConnectionString()
	var f FeedbackModel
	session, err := confInstance.GetMgoSession(mgoUri)
	if err != nil {
		fmt.Printf("can not connect to server %v\n", err)
		panic(err)
	}
	defer session.Close()

	collection := session.DB(confInstance.DBName).C("feedback")
	err = collection.Find(bson.M{"feedbackcode": modelId}).One(&f)
	if err != nil {
		fmt.Printf("insert error - %v\n", err)
		return &FeedbackModel{}, err
	}

	return &f, nil
}

func (m *FeedbackModel) InsertModel(model *FeedbackModel) error {
	confInstance := conf.ConfigInstance()
	mgoUri := confInstance.MongoDBConnectionString()
	session, err := confInstance.GetMgoSession(mgoUri)
	if err != nil {
		fmt.Printf("can not connect to server %v\n", err)
		panic(err)
	}
	defer session.Close()

	f := session.DB(confInstance.DBName).C("feedback")
	err = f.Insert(model)

	return err
}

func (m *FeedbackModel) UpdateModel(selector map[string]string, model *FeedbackModel) error {
	confInstance := conf.ConfigInstance()
	mgoUri := confInstance.MongoDBConnectionString()
	session, err := confInstance.GetMgoSession(mgoUri)
	if err != nil {
		fmt.Printf("can not connect to server %v\n", err)
		panic(err)
	}
	defer session.Close()

	f := session.DB(confInstance.DBName).C("feedback")
	err = f.Update(selector, model)

	return err
}
