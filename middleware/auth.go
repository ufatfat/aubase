package middleware

import (
	"AUBase/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func IsSignedIn (c *gin.Context) {
	a := c.GetHeader("Authorization")
	if a != "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg": "您已登录！",
		})
	}
}

func IsSignedOut (c *gin.Context) {
	a := c.GetHeader("Authorization")
	if a == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"msg": "请先登录！",
		})
		return
	}
	div := strings.Index(a, " ")
	if div == -1 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"msg": "身份验证失败，请重新登录！",
		})
		return
	}

	// 获取Authorization的scheme
	if scheme := a[:div]; scheme != "Bearer" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"msg": "身份验证失败，请重新登录！",
		})
		return
	}
	if ok := util.ValidateToken(a[div+1:]); !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"msg": "身份验证失败，请重新登录！",
		})
		return
	}
}