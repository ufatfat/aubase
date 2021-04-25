package controller

import (
	"AUBase/service"
	"AUBase/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func GetWorkNum (c *gin.Context) {
	num := service.GetWorkNum()
	c.JSON(http.StatusOK, gin.H{
		"msg": "获取作品数量成功！",
		"data": gin.H{
			"num": num,
		},
	})
}


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
	userID, err := util.GetIDFromToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	workInfo, err := service.GetWorkInfo(workID, userID)
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
		"data": workInfo,
		"token": util.UpdateToken(token),
	})
}