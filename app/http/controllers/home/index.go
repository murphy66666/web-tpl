package home

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web-tpl/app"
	"web-tpl/app/http/models"
)

func Index(ctx *gin.Context) {
	app.Log().Info("999")
	//ctx.JSON(http.StatusOK, map[string]any{
	//	"code": 0,
	//	"msg":  "success",
	//	"data": "666",
	//})
	//callback 值只允许字母数字下线 过滤处理不好 要加过滤处理限制
	ctx.JSONP(http.StatusOK, map[string]any{
		"code": 0,
		"msg":  "success",
		"data": "666",
	})
}

func UserList(ctx *gin.Context) {
	var rel []models.User

	err := app.DBR().Find(&rel).Error
	if err != nil {
		ctx.JSON(200, gin.H{
			"code": 1,
			"msg":  err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"code": 0,
		"data": rel,
	})
}
