package router

import (
	"AUBase/controller"
	"AUBase/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter () (r *gin.Engine) {
	r = gin.Default()

	user := r.Group("/user")
	{
		user.POST("/signin", middleware.IsSignedIn, controller.SignIn)
		user.Use(middleware.IsSignedOut)
		user.POST("/signout", controller.SignOut)
		user.PUT("/password", controller.ChangePassword)
		user.GET("/myvote", controller.MyVotedWork)
	}
	vote := r.Group("/vote")
	{
		vote.Use(middleware.IsSignedOut)
		vote.GET("/:workID", controller.GetWorkToVote)
		vote.PUT("", controller.VoteForWork)
	}

	return
}