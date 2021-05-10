package controller

import (
	"aubase/model"
	"aubase/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func VoteForWork (c *gin.Context) {
	activityID, _ := c.Get("activityID")
	userID, _ := c.Get("userID")
	turnID, _ := c.Get("turnID")
	var voteInfo model.VoteInfo
	if err := c.BindJSON(&voteInfo); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	if err := service.VoteForWork(activityID.(uint32), userID.(uint32), turnID.(uint32), voteInfo); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "投票成功",
	})

}

func VoteDone (c *gin.Context) {
	userID, _ := c.Get("userID")
	turnID, _ := c.Get("turnID")
	if err := service.VoteDone(userID.(uint32), turnID.(uint32)); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "提交成功！",
	})
}

func CheckIsDone (c *gin.Context) {
	userID, _ := c.Get("userID")
	turnID, _ := c.Get("turnID")
	ok := service.CheckIsDone(userID.(uint32), turnID.(uint32))
	c.JSON(http.StatusOK, gin.H{
		"msg": "查询成功！",
		"data": gin.H{
			"is_done": ok,
		},
	})
}

func GetVotedNum (c *gin.Context) {
	userID, _ := c.Get("userID")
	turnID, _ := c.Get("turnID")
	num := service.GetVotedNum(userID.(uint32), turnID.(uint32))
	c.JSON(http.StatusOK, gin.H{
		"msg": "查询成功！",
		"data": gin.H{
			"num": num,
		},
	})
}