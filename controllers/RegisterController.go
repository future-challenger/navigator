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
	registerControllerTag = "RegisterController"
)

type RegisterController struct {
}

func (controller *RegisterController) Register(req *http.Request, r render.Render) {
	defer func() {
		fmt.Errorf("LoginController error :- %v" + " Error")
		if err := recover(); err != nil {
			r.JSON(200, map[string]interface{}{"state": conf.ErrorCode["NORMAL_ERROR"], "message": "fail"})
		}
	}()

	req.ParseForm()

	userInfoJson := req.FormValue("userInfo")
	if len(userInfoJson) <= 0 {
		r.JSON(200, map[string]interface{}{"state": conf.ErrorCode["NORMAL_ERROR"], "message": "data error"})
		return
	}

	userModel := models.NewUserModel(userInfoJson)
	findUserList := userModel.ModelList(bson.M{"username": userModel.UserName})
	if len(findUserList) > 0 {
		r.JSON(200, map[string]interface{}{"state": conf.ErrorCode["NORMAL_ERROR"], "message": "user exists"})
		return
	}

	err := userModel.InsertModel(userModel)
	if err != nil {
		r.JSON(200, map[string]interface{}{"state": conf.ErrorCode["NORMAL_ERROR"], "message": "data error"})
	} else {
		r.JSON(200, map[string]interface{}{"state": conf.ErrorCode["ALL_OK"], "message": "success", "data": "OK"})
	}

}
