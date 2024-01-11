package home

import (
	"github.com/gin-gonic/gin"
	"web-tpl/app/http/controllers/home/params"
	"web-tpl/app/utils/rsp"
)

func Add(c *gin.Context) {
	var p params.Add
	err := c.ShouldBind(&p)
	if err != nil {
		rsp.JSONErr(c, rsp.WithMsg(err.Error()))
		return
	}
	rsp.JSONOk(c)
}
