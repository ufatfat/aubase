package controller

import (
	"aubase/model"
	"aubase/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func VoteForWork (c *gin.Context) {
	activityID, _ := c.Get("activityID")
	if activityID == nil {
		fmt.Println("a")
	}
	userID, _ := c.Get("userID")
	if userID == nil {
		fmt.Println("u")
	}
	turnID, _ := c.Get("turnID")
	if turnID == nil {
		fmt.Println("t")
	}
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