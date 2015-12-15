package conf

import (
	"fmt"
	"gopkg.in/mgo.v2"
	// "gopkg.in/mgo.v2/bson"
)

type Configuration struct {
	DBName string
	DBHost string
	DBPort string
}

func ConfigInstance() *Configuration {
	config := Configuration{
		DBName: "NDB",
		DBHost: "localhost",
		DBPort: "27017",
	}
	return &config
}

func (c *Configuration) MongoDBConnectionString() string {
	if c == nil {
		return ""
	}

	connectionString := "mongodb://" + c.DBHost + ":" + c.DBPort + "/" + c.DBName

	return connectionString
}

var mgoSession *mgo.Session = nil

func (c *Configuration) GetMgoSession(mgoUri string) (*mgo.Session, error) {
	if mgoSession == nil {
		var err error
		mgoSession, err = mgo.Dial(mgoUri)
		if err != nil {
			fmt.Printf("can not connect to server %v\n", err)
			// panic(err)
			return nil, err
		}
		// defer session.Close()
	}

	if mgoSession.Ping() == nil {
		var err error
		mgoSession, err = mgo.Dial(mgoUri)
		if err != nil {
			return nil, err
		}
	}

	return mgoSession.Copy(), nil
}
