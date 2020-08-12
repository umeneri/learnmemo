package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/k0kubun/pp"
)

func PrintAccessLog(c *gin.Context) {
	pp.Println(c.Request.Method)
	pp.Println(c.Request.RequestURI)
	c.Next()
}
