package controllers

import (
	"fmt"
	"github.com/martini-contrib/render"
	"gopkg.in/mgo.v2/bson"
	"navigator/conf"
	"navigator/models"
	"net/http"
)

type BooksController struct {
}

func (controller *BooksController) BookList(req *http.Request, r render.Render) {
	defer func() {
		fmt.Println("book list")
		if err := recover(); err != nil {
			r.JSON(200, map[string]interface{}{"state": conf.ErrorCode["NORMAL_ERROR"], "message": "fail"})
		}
	}()

	req.ParseForm()
	var bookKeyword = req.FormValue("bookName")
	fmt.Println("book name:- ", bookKeyword)

	var searchCondition bson.M
	if len(bookKeyword) > 0 {
		// searchCondition = map[string]string{"bookname": "/" + bookKeyword + "/"}
		searchCondition = bson.M{"bookname": bson.RegEx{bookKeyword, ""}}
	} else {
		//TODO: it's wrong if book name is empty. return error result json here
		// or just do nothing here.
		searchCondition = bson.M{"bookname": ""}
	}

	bookModel := new(models.BookModel)
	bookModelList := bookModel.ModelList(searchCondition)
	r.JSON(200, map[string]interface{}{"state": conf.ErrorCode["ALL_OK"], "message": "sucess", "data": bookModelList})
}

func (controller *BooksController) AddBook(req *http.Request, r render.Render) {
	defer func() {
		if err := recover(); err != nil {
			r.JSON(200, map[string]interface{}{"state": conf.ErrorCode["NORMAL_ERROR"], "message": "fail"})
		}
	}()
	req.ParseForm()
	newBookJson := req.FormValue("newBook")

	fmt.Println("new book json:- ", newBookJson)

	bookModel := models.NewBookModel(newBookJson)
	bookModel.InsertModel(bookModel)
	r.JSON(200, map[string]interface{}{"state": conf.ErrorCode["ALL_OK"], "message": "sucess"})
}

func (controller *BooksController) UpdateBook(req *http.Request, r render.Render) {
	defer func() {
		if err := recover(); err != nil {
			r.JSON(200, map[string]interface{}{"state": conf.ErrorCode["NORMAL_ERROR"], "message": "fail"})
		}
	}()
	req.ParseForm()
	newBookJson := req.FormValue("newBook")

	fmt.Println("new book json:- ", newBookJson)

	bookModel := models.NewBookModel(newBookJson)
	bookModel.UpdateModel(bson.M{"bookcode": bookModel.BookCode}, bookModel)
	r.JSON(200, map[string]interface{}{"state": conf.ErrorCode["ALL_OK"], "message": "sucess"})
}
