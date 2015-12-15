package controllers

import (
	"fmt"
	"github.com/martini-contrib/render"
	"gopkg.in/mgo.v2/bson"
	"navigator/conf"
	"navigator/models"
	"net/http"
)

const (
	loginControllerTag = "LoginController"
)

/*
   UserCode        string `json: "userCode"`
   UserName        string `json: "userName"`
*/

type LoginController struct {
}

func (controller *LoginController) Login(req *http.Request, r render.Render) {
	defer func() {
		fmt.Errorf("LoginController error :- %v" + " Error")
		if err := recover(); err != nil {
			r.JSON(200, map[string]interface{}{"state": conf.ErrorCode["NORMAL_ERROR"], "message": "fail"})
		}
	}()

	req.ParseForm()
	//TODO: md5 -> password later
	userName := req.FormValue("userName")
	password := req.FormValue("pwd")

	// userModel := models.NewUserModel("{\"userCode\": \"" + userName + "\"\"password\":\"" + password + "\"}")
	userModel := new(models.UserModel)
	userModelList := userModel.ModelList(bson.M{"username": userName, "password": password})
	if len(userModelList) > 0 {
		r.JSON(200, map[string]interface{}{"state": conf.ErrorCode["ALL_OK"], "message": "success", "data": userModelList[0]})
	} else {
		r.JSON(200, map[string]interface{}{"state": conf.ErrorCode["NORMAL_ERROR"], "message": "fail"})
	}
}
