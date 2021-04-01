package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func IsSignedIn (c *gin.Context) {
	if _, err := c.Cookie("token"); err == nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"msg": "您已登录！",
		})
	}
}

func IsSignedOut (c *gin.Context) {
	if _, err := c.Cookie("token"); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"msg": "请先登录！",
		})
	}
}