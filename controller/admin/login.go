package admin

import (
	"cms/common"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Login(c *gin.Context) {
	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "token",
		Value:   "12425263",
		Expires: time.Now().Add(10 * time.Second),
	})
	c.JSON(200, common.Json(0, "签名通过"))
}
