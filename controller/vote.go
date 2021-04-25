package controller

import (
	"AUBase/model"
	"AUBase/service"
	"AUBase/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func GetWorkToVote (c *gin.Context) {
	token := c.GetHeader("Authorization")[7:]
	userID, err := util.GetIDFromToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	workToVote, err := service.GetWorkToVoteByUserID(userID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusOK, gin.H{
				"msg": "没有未投作品！",
			})
			return
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"msg": err.Error(),
			})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"data": workToVote,
		"token": util.UpdateToken(token),
	})
}

func GetWorkToVoteByID (c *gin.Context) {
	workID, err := strconv.ParseUint(c.Param("workID"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	token := c.GetHeader("Authorization")[7:]

	workToVote, err := service.GetWorkToVote(workID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusOK, gin.H{
				"msg": "没有此作品！",
			})
			return
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"msg": err.Error(),
			})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"data": workToVote,
		"token": util.UpdateToken(token),
	})
}

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
	if err = service.VoteForWork(voteInfo.WorkID, userID, voteInfo.Positive); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "投票成功！",
		"status": voteInfo.Positive,
		"token": util.UpdateToken(token),
	})
}