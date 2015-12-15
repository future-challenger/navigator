package controllers

import (
	"fmt"
	"github.com/martini-contrib/render"
	"gopkg.in/mgo.v2/bson"
	"navigator/conf"
	"navigator/models"
	"net/http"
)

type RuleController struct {
}

func (controller *RuleController) RuleList(req *http.Request, r render.Render) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Errorf("RuleController error:- %v", err)
			r.JSON(200, map[string]interface{}{"state": conf.ErrorCode["NORMAL_ERROR"], "message": "fail"})
		}
	}()

	ruleModel := new(models.RuleModel)
	modelList := ruleModel.ModelList(nil)

	r.JSON(200, map[string]interface{}{"state": conf.ErrorCode["ALL_OK"], "message": "sucess", "data": modelList})
}

func (controller *RuleController) GetRule(req *http.Request, r render.Render) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Errorf("RuleController error:- %v", err)
			r.JSON(200, map[string]interface{}{"state": conf.ErrorCode["NORMAL_ERROR"], "message": "fail"})
		}
	}()

	req.ParseForm()
	ruleId := req.FormValue("ruleId")
	tempRuleModel := new(models.RuleModel)
	ruleModel, _ := tempRuleModel.GetModel(ruleId)
	if ruleModel == nil {
		fmt.Println("No rule with such id : ", ruleId)
		r.JSON(200, map[string]interface{}{"state": conf.ErrorCode["NORMAL_ERROR"], "message": "fail"})
	} else {
		r.JSON(200, map[string]interface{}{"state": conf.ErrorCode["ALL_OK"], "message": "sucess", "data": ruleModel})
	}
}

func (controller *RuleController) UpdateRule(req *http.Request, r render.Render) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Errorf("RuleController error:- %v", err)
			r.JSON(200, map[string]interface{}{"state": conf.ErrorCode["NORMAL_ERROR"], "message": "fail"})
		}
	}()

	req.ParseForm()
	ruleModelJson := req.FormValue("ruleJson")
	ruleModel := models.NewRuleModel(ruleModelJson)
	err := ruleModel.UpdateModel(bson.M{"rulecode": ruleModel.RuleCode}, ruleModel)
	if err == nil {
		r.JSON(200, map[string]interface{}{"state": conf.ErrorCode["ALL_OK"], "message": "sucess", "data": ruleModel})
	} else {
		r.JSON(200, map[string]interface{}{"state": conf.ErrorCode["NORMAL_ERROR"], "message": "fail"})
	}
}

func (controller *RuleController) InsertRule(req *http.Request, r render.Render) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Errorf("RuleController error:- %v", err)
			r.JSON(200, map[string]interface{}{"state": conf.ErrorCode["NORMAL_ERROR"], "message": "fail"})
		}
	}()

	req.ParseForm()
	ruleModelJson := req.FormValue("ruleJson")
	ruleModel := models.NewRuleModel(ruleModelJson)
	err := ruleModel.InsertModel(ruleModel)
	if err == nil {
		r.JSON(200, map[string]interface{}{"state": conf.ErrorCode["ALL_OK"], "message": "sucess", "data": ruleModel})
	} else {
		r.JSON(200, map[string]interface{}{"state": conf.ErrorCode["NORMAL_ERROR"], "message": "fail"})
	}
}
