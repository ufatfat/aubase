package controller

import (
	"AUBase/service"
	"AUBase/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func GetWorkToVote (c *gin.Context) {
	token, _ := c.Cookie("token")
	userID, err := util.GetIDFromToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	workToVote, err := service.GetWorkToVote(userID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusOK, gin.H{
				"msg": "没有未评作品！",
			})
			return
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"msg": err.Error(),
			})
			return
		}
	}
	c.JSON(http.StatusOK, workToVote)
}

func VoteForWork (c *gin.Context) {
	token, _ := c.Cookie("token")
	userID, err := util.GetIDFromToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	w := c.PostForm("work_id")
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
	}
	if err = service.VoteForWork(workID, userID, positive); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "投票成功！",
		"status": positive,
	})
}