package models

import (
	"fmt"
	"strconv"
	"testing"
)

func TestFeedbackModel(t *testing.T) {
	defer func() {
		ex := recover()
		if ex != nil {
			fmt.Errorf("%v", ex)
		}
	}()

	f := FeedbackModel{FeedbackCode: "123", Desc: "testing"}
	f.InsertModel(&f)
	temp, err := f.GetModel("123")
	if err != nil {
		fmt.Println("error is ", err)
		t.Fail()
	}

	if temp.FeedbackCode != f.FeedbackCode {
		t.Fail()
	}
}

// func TestExist(t *testing.T) {
// 	f := new(FeedbackModel)
// 	model := f.GetModel("123")

// 	if model == nil {
// 		t.Error("not found")
// 	} else {
// 		t.Log("Feedback code:- " + model.FeedbackCode)
// 	}
// }

func TestModelList(t *testing.T) {
	f := new(FeedbackModel)
	feedbackList := f.ModelList(nil)

	fmt.Println("Length is " + strconv.Itoa(len(feedbackList)))

	if len(feedbackList) == 0 {
		t.Error("length is ZERO")
	}

	if len(feedbackList) != 8 {
		t.Error("Not equals to SIX")
	}

	fmt.Println("Successful")
}
