package controller

import (
	"aubase/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
	if workInfo.WorkGroup == "End" {
		c.JSON(http.StatusOK, gin.H{
			"msg": "当前轮次作品已浏览完！",
		})
		return
	}
	if workInfo.WorkGroup == "No" {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "当前轮次无此作品！",
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
	turnID, _ := c.Get("turnID")
	w := c.Param("workID")
	if w == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg": "参数错误！",
		})
	}
	workID, err := strconv.ParseUint(w, 10, 32)
	workInfo, err := service.GetWorkToVoteByID(userID.(uint32), turnID.(uint32), uint32(workID))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	if workInfo.WorkID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "当前轮次无此作品！",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "查询成功！",
		"data": workInfo,
	})
}

func GetWorkByGroup (c *gin.Context) {
	userID, _ := c.Get("userID")
	turnID, _ := c.Get("turnID")

	g := c.Query("groupID")
	groupID, err := strconv.ParseUint(g, 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg": "参数错误",
		})
		return
	}
	get := c.Query("get")
	workInfos, err := service.GetWorkByGroup(uint32(groupID), userID.(uint32), turnID.(uint32), get)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "查询成功！",
		"data": workInfos,
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