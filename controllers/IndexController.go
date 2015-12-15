package controllers

import (
	"github.com/martini-contrib/render"
	"net/http"
)

type IndexController struct {
}

func (c *IndexController) List(req *http.Request, r render.Render) {
	r.JSON(200, map[string]interface{}{"hello": "world"})
}
