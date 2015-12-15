package models

import (
	"fmt"
	"strconv"
	"testing"
)

func TestBookModel(t *testing.T) {
	defer func() {
		ex := recover()
		if ex != nil {
			fmt.Errorf("%v", ex)
		}
	}()

	b := BookModel{BookCode: "123", BookName: "支付战争", Author: "Jackson"}
	b.InsertModel(&b)
	temp, err := b.GetModel("123")
	if err != nil {
		fmt.Println("error is ", err)
		t.Fail()
	}

	if temp.BookCode != b.BookCode {
		t.Fail()
	}
}

func TestBookModelList(t *testing.T) {
	b := new(BookModel)
	bookModelList := b.ModelList(nil)

	fmt.Println("Length is " + strconv.Itoa(len(bookModelList)))

	if len(bookModelList) == 0 {
		t.Error("length is ZERO")
	}

	if len(bookModelList) != 1 {
		t.Error("Not equals to SIX")
	}

	fmt.Println("Successful")
}
