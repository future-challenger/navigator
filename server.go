package main

import (
	"fmt"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"navigator/controllers"
)

var m = martini.Classic()

func prepare() {
	m.Use(martini.Recovery())
	m.Use(martini.Logger())
	m.Use(render.Renderer())

	// routers
	var r = martini.NewRouter()
	indexController := new(controllers.IndexController)
	feedbackController := new(controllers.FeedbackController)
	booksController := new(controllers.BooksController)

	r.Get("/nav/index/:name", indexController.List)
	//feedback
	r.Get("/nav/feedback", feedbackController.FeedbackList)
	r.Post("/nav/feedback", feedbackController.Feedback)
	// r.Post("/nav/register", feedbackController.Register)
	// books
	r.Get("/nav/books", booksController.BookList)
	r.Post("/nav/book/new", booksController.AddBook)
	r.Post("/nav/book/update", booksController.UpdateBook)

	m.Action(r.Handle)
}

func talk() {
	fmt.Println("hello")
}

func main() {
	talk()
	prepare()

	m.RunOnAddr(":9090")
}
