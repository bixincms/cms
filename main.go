package main

import (
	"cms/controller/admin"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "暂时不知道写什么")
	})
	manager := r.Group("/manager")
	manager.GET("/login", admin.Login)
	r.Run(":80")
}
