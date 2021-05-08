package middleware

import (
	"aubase/service"
	"aubase/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
	token := a[div+1:]
	if ok := util.ValidateToken(token); !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"msg": "身份验证失败，请重新登录！",
		})
		return
	}
	userID, err := util.GetIDFromToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.Set("userID", userID)
}

func IsActivityOpen (c *gin.Context) {
	c.Header("Access-Control-Allow-Headers", "X-Requested-With")
	a := c.GetHeader("Activity")
	if a == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg": "本比赛暂未开放！",
		})
		return
	}
	activityID, err := strconv.ParseUint(a, 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	if ok := service.CheckActivityOpen(uint32(activityID)); !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg": "本比赛暂未开放！",
		})
		return
	}
	c.Set("activityID", uint32(activityID))
}

func IsTurnOpen (c *gin.Context) {
	t := c.Query("turn")
	if t == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg": "本轮次暂未开放！",
		})
		return
	}
	turnID, err := strconv.ParseUint(t, 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	if ok := service.CheckTurnOpen(uint32(turnID)); !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg": "本轮次暂未开放！",
		})
		return
	}
	c.Set("turnID", uint32(turnID))
}