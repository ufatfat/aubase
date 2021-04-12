package controller

import (
	"AUBase/service"
	"AUBase/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
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
	c.JSON(http.StatusOK, gin.H{
		"msg": "登录成功！",
		"data": gin.H{
			"name": userInfo.Name,
			"token": token,
		},
		"token": util.UpdateToken(token),
	})
}

func SignOut (c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "登出成功！",
	})
}

func MyVotedWork (c *gin.Context) {
	token := c.GetHeader("Authorization")[7:]
	id, err := util.GetIDFromToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	work, err := service.MyVotedWork(id)
	c.JSON(http.StatusOK, gin.H{
		"data": work,
		"token": util.UpdateToken(token),
	})
}

func ChangePassword (c *gin.Context) {

}