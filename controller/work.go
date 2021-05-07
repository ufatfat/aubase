package controller

import (
	"aubase/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetWorkToVote (c *gin.Context) {
	userID, _ := c.Get("userID")
	activityID, _ := c.Get("activityID")
	turnID, _ := c.Get("turnID")
	workInfo, err := service.GetWorkToVote(userID.(uint32), activityID.(uint32), turnID.(uint32))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	if workInfo.WorkID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"msg": "当前轮次无此作品！",
		})
		return
	}
	if workInfo.WorkGroup == "End" {
		c.JSON(http.StatusOK, gin.H{
			"msg": "当前轮次作品已浏览完！",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "查询成功！",
		"data": workInfo,
	})
}
func GetWorkToVoteByID (c *gin.Context) {
	userID, _ := c.Get("userID")
	activityID, _ := c.Get("activityID")
	turnID, _ := c.Get("turnID")
	w := c.Param("workID")
	if w == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg": "参数错误！",
		})
	}
	workInfo, err := service.GetWorkToVote(userID.(uint32), activityID.(uint32), turnID.(uint32))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "查询成功！",
		"data": workInfo,
	})
}

func GetWorkRange (c *gin.Context) {
	turnID, _ := c.Get("turnID")
	workRange := service.GetWorkRange(turnID.(uint32))
	c.JSON(http.StatusOK, gin.H{
		"msg": "查询成功！",
		"data": workRange,
	})
}

func GetWorkNum (c *gin.Context) {
	turnID, _ := c.Get("turnID")
	workNum := service.GetWorkNum(turnID.(uint32))
	c.JSON(http.StatusOK, gin.H{
		"msg": "查询成功！",
		"data": gin.H{
			"num": workNum,
		},
	})
}