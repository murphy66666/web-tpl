package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web-tpl/app"
	"web-tpl/app/http/models"
)

func Index(c *gin.Context) {
	var rel []models.User

	err := app.DBW().Find(&rel).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  rel,
	})

	return
}

func Login(c *gin.Context) {
	//Auth()
}

func Logout(c *gin.Context) {
	//Auth()
}

func Info(c *gin.Context) {
	//Auth()
}

func Update(c *gin.Context) {
	//Auth()
}
