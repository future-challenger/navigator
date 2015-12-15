package models

import (
	"fmt"
	"strconv"
	"testing"
)

/*
    type UserModel struct {
    UserCode     string `json: "userCode"`
    UserName     string `json: "userName"`
    Password     string `json: "password"`
    DetpartModel `json: "deptInfo"`
    Gender       string `json: "gender"`
}

type DepartmentModel struct {
    DeptCode string `json: "deptCode"`
    DeptName string `json: "deptName"`
}
*/

func TestUserModel(t *testing.T) {
	defer func() {
		ex := recover()
		if ex != nil {
			fmt.Errorf("%v", ex)
		}
	}()

	u := UserModel{UserCode: "123", UserName: "不是人", Password: "123",
		DepartmentModel: DepartmentModel{"456", "开发部"}, Gender: "妹纸"}
	u.InsertModel(&u)
	temp, err := u.GetModel("123")
	if err != nil {
		fmt.Println("error is ", err)
		t.Fail()
	}

	if temp.UserCode != u.UserCode {
		t.Fail()
	}
}

func TestUserModelList(t *testing.T) {
	u := new(UserModel)
	userModelList := u.ModelList(nil)

	fmt.Println("Length is " + strconv.Itoa(len(userModelList)))

	if len(userModelList) == 0 {
		t.Error("length is ZERO")
	}

	if len(userModelList) != 2 {
		t.Error("Not equals to 2")
	}

	fmt.Println("Successful")
}
