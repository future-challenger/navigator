package models

import (
	"encoding/json"
	"fmt"
	// "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"navigator/conf"
)

const (
	userPrimaryKey     = "usercode"
	userCollectionName = "user"
)

type DepartmentModel struct {
	DeptCode string `json: "deptCode"`
	DeptName string `json: "deptName"`
}

type UserModel struct {
	UserCode        string `json: "userCode"`
	UserName        string `json: "userName"`
	Password        string `json: "password"`
	DepartmentModel `json: "deptInfo"`
	Gender          string `json: "gender"`
}

func NewUserModel(dict string) *UserModel {
	var u UserModel
	err := json.Unmarshal([]byte(dict), &u)
	if err != nil {
		return nil
	}

	return &u
}

func (model *UserModel) ModelList(conditions bson.M) []UserModel {
	confInstance := conf.ConfigInstance()
	mgoUri := conf.ConfigInstance().MongoDBConnectionString()

	session, err := confInstance.GetMgoSession(mgoUri)
	if err != nil {
		fmt.Printf("can not connect to server %v\n", err)
		panic(err)
	}
	defer session.Close()

	var userModelList []UserModel
	collection := session.DB(confInstance.DBName).C(userCollectionName)
	err = collection.Find(conditions).All(&userModelList)

	return userModelList
}

func (m *UserModel) GetModel(modelId string) (*UserModel, error) {
	confInstance := conf.ConfigInstance()
	mgoUri := conf.ConfigInstance().MongoDBConnectionString()

	var userModel UserModel
	session, err := confInstance.GetMgoSession(mgoUri)
	if err != nil {
		fmt.Printf("can not connect to server %v\n", err)
		panic(err)
	}
	defer session.Close()

	collection := session.DB(confInstance.DBName).C(userCollectionName)
	err = collection.Find(bson.M{primaryKey: modelId}).One(&userModel)
	if err != nil {
		fmt.Printf("insert error - %v\n", err)
		return &UserModel{}, err
	}

	return &userModel, nil
}

func (m *UserModel) InsertModel(model *UserModel) error {
	confInstance := conf.ConfigInstance()
	mgoUri := confInstance.MongoDBConnectionString()
	session, err := confInstance.GetMgoSession(mgoUri)
	if err != nil {
		fmt.Printf("can not connect to server %v\n", err)
		panic(err)
	}
	defer session.Close()

	f := session.DB(confInstance.DBName).C(userCollectionName)
	err = f.Insert(model)

	return err
}

func (m *UserModel) UpdateModel(selector bson.M, model *UserModel) error {
	confInstance := conf.ConfigInstance()
	mgoUri := confInstance.MongoDBConnectionString()

	session, err := confInstance.GetMgoSession(mgoUri)
	if err != nil {
		fmt.Printf("can not connect to server %v\n", err)
		panic(err)
	}
	defer session.Close()

	f := session.DB(confInstance.DBName).C(userCollectionName)
	err = f.Update(selector, model)

	return err
}
