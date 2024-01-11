package routers

import (
	"github.com/gin-gonic/gin"
	"web-tpl/app/http/controllers/home"
	"web-tpl/app/http/middleware"
)

func RegTest(r *gin.Engine) {
	r.Use(middleware.HelloA)
	r.GET("/", middleware.HelloB, home.Index)
	r.GET("/v1/user/add", home.Add)

}
