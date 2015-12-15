package controllers

import (
	// "fmt"
	"fmt"
	"github.com/martini-contrib/render"
	"navigator/conf"
	"navigator/models"
	"net/http"
)

type FeedbackController struct {
}

func (c *FeedbackController) FeedbackList(req *http.Request, r render.Render) {

	f := new(models.FeedbackModel)
	r.JSON(200, map[string]interface{}{"data": f.ModelList(nil)})
}

func (c *FeedbackController) Feedback(req *http.Request, r render.Render) {
	req.ParseForm()
	feedbackVal := req.FormValue("feedbackKey")
	action := req.FormValue("action")
	fmt.Println("feedback value:- " + feedbackVal + " action:- " + action)

	var err error
	feedbackModel := models.NewFeedbackModel(feedbackVal)
	if action == "insert" {
		err = feedbackModel.InsertModel(feedbackModel)
	} else if action == "update" {
		err = feedbackModel.UpdateModel(map[string]string{"feedbackCode": feedbackModel.FeedbackCode}, feedbackModel)
	} else {
		fmt.Printf("FeedbackController: there's no such action")
		r.JSON(200, map[string]string{"state": conf.ErrorCode["NO_SUCH_ACTION"], "message": "fail"})
	}

	if err != nil {
		r.JSON(200, map[string]string{"state": conf.ErrorCode["NORMAL_ERROR"], "message": "fail"})
	} else {
		r.JSON(200, map[string]string{"state": conf.ErrorCode["ALL_OK"], "message": "sucess"})
	}
}
