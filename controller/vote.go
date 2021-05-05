package controller

import (
	"AUBase/model"
	"AUBase/service"
	"AUBase/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)


func VoteForWork (c *gin.Context) {
	token := c.GetHeader("Authorization")[7:]
	userID, err := util.GetIDFromToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	/*w := c.PostForm("work_id")
	workID, err := strconv.ParseUint(w, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	p := c.PostForm("positive")
	positive, err := strconv.ParseBool(p)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}*/
	var voteInfo model.VoteInfo
	if err = c.BindJSON(&voteInfo); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg": "参数错误",
		})
		return
	}
	fmt.Println(voteInfo)
	if err = service.VoteForWork(voteInfo.WorkID, userID, voteInfo.Negative); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "投票成功！",
		"is_negative": voteInfo.Negative,
		"token": util.UpdateToken(token),
	})
}

func GetVotedNum (c *gin.Context) {
	token := c.GetHeader("Authorization")[7:]
	userID, err := util.GetIDFromToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	a := c.Query("activity")
	t := c.Query("turn")
	activityID, err := strconv.ParseUint(a, 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	turnID, err := strconv.ParseUint(t, 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	votedNum, err := service.GetVotedNum(userID, uint32(activityID), uint32(turnID))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "查询成功！",
		"data": gin.H{
			"voted_num": votedNum,
		},
	})
}