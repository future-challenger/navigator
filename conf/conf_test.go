package conf

import "testing"

func TestConfiguration(t *testing.T) {
	t.Log("start test configuration")

	configuration := ConfigInstance()
	if configuration.MongoDBConnectionString() == "" {
		t.Error("Connection string is empty")
	}

	if "mongodb://localhost:27017/NDB" != configuration.MongoDBConnectionString() {
		t.Errorf("Connection string is not right with := %v", configuration.MongoDBConnectionString())
	}
}
