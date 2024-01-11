package routers

import (
	"github.com/gin-gonic/gin"
	"web-tpl/app/http/controllers/home"
	"web-tpl/app/http/controllers/qrcode"
	"web-tpl/app/http/controllers/user"
)

// Reg 路由只负责注册
func Reg(r *gin.Engine) {
	r.GET("/", home.Index)
	r.POST("/v1/qrcode", qrcode.Gen)
	r.GET("/v1/userlist", home.UserList)

	//用户模块
	r.GET("/user/index", user.Index)
	//r.POST("/user/login","middleware.ServerName",user.Login) //用户登录
	//r.POST("/user/logout",user.Logout) //用户退出
	//r.GET("/user/info",user.Info) //获取用户信息
	//r.POST("/user/info",user.Update) //修改用户信息
}
