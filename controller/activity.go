package controller

import (
	"aubase/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func GetTurnInfo (c *gin.Context) {
	activityID, _ := c.Get("activityID")

	turnInfo, err := service.GetTurnInfo(activityID.(uint32))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{
				"msg": "当前比赛没有开放的轮次！",
			})
			return
		} else {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"msg": err.Error(),
			})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "查询成功！",
		"data": turnInfo,
	})

}