package controllers

import (
	"github.com/martini-contrib/render"
	"net/http"
)

type RegisterController struct {
}

func (controller *RegisterController) Register(req *http.Request, r render.Render) {
	r.JSON(200, map[string]interface{}{"hello": "world"})
}
