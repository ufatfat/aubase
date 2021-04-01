package controller

import (
	"AUBase/service"
	"AUBase/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func SignIn (c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username == "" || password == "" {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"msg": "用户名/密码不能为空！",
		})
		return
	}
	password = util.PasswordEncrypt(password)
	userInfo, err := service.SignIn(username, password)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{
				"msg": "用户名/密码错误！",
			})
			return
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"msg": err.Error(),
			})
			return
		}
	}
	token, err := util.GenToken(userInfo.ID, userInfo.Name)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.SetCookie("token", token, -1, "/", ".aubase.cn", true, true)
	c.JSON(http.StatusOK, gin.H{
		"msg": "登录成功！",
		"data": gin.H{
			"name": userInfo.Name,
		},
	})
}

func SignOut (c *gin.Context) {
	c.SetCookie("token", "", 0, "/", ".aubase.cn", true, true)
	c.JSON(http.StatusOK, gin.H{
		"msg": "登出成功！",
	})
}

func MyVotedWork (c *gin.Context) {
	p := c.Query("positive")
	positive, err := strconv.ParseBool(p)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	token, _ := c.Cookie("token")
	id, err := util.GetIDFromToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	work, err := service.MyVotedWork(id, positive)
	c.JSON(http.StatusOK, work)
}

func ChangePassword (c *gin.Context) {

}