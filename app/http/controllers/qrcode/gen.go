package qrcode

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
	"net/http"
)

// 验证器
type params struct {
	URL string `form:"url" binding:"required,url"`
}

func Gen(ctx *gin.Context) {
	var p params

	if err := ctx.ShouldBind(&p); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  err.Error(),
		})
		return
	}

	//生成二维码
	var png []byte
	png, err := qrcode.Encode(p.URL, qrcode.Medium, 256)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 2,
			"msg":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": gin.H{
			"url": "data:image/png;base64," + base64.StdEncoding.EncodeToString(png), //base64编码
		},
	})
}
