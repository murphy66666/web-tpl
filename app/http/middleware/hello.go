package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func HelloA(c *gin.Context) {
	fmt.Println("-- before HelloA --")
	c.Next()
	fmt.Println("-- after HelloA --")
}

func HelloB(c *gin.Context) {
	fmt.Println("-- before HelloB --")
	c.Next()
	fmt.Println("-- after HelloB --")
}
