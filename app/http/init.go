package http

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"web-tpl/app"
	"web-tpl/app/core/valid"
	"web-tpl/app/http/middleware/logger"
	"web-tpl/app/http/routers"
)

func NewServer() error {
	r := gin.New()
	if app.Config.WebServerLog.Enable {
		r.Use(logger.New(app.Config.WebServerLog, app.Config.Env, app.Config.HomeDir))
	}
	r.Use(gin.Recovery())
	//全局初始化配置

	//验证器新增函数
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("valid", valid.Mobile)
	}

	//跨域 全局配置
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://127.0.0.1:5173"},
		AllowMethods:     []string{"POST", "PUT", "GET"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	//路由加载
	routers.Reg(r)

	return r.Run(app.Config.HTTP.Listen)
}
