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
		vote.PUT("", controller.VoteForWork)
	}
	work := r.Group("/work")
	{
		vote.Use(middleware.IsSignedOut)
		work.GET("", controller.GetWorkToVote)
		work.GET("/:workID", func (c *gin.Context) {
			p := c.Param("workID")
			switch p {
			case "num":
				controller.GetWorkNum(c)
			default:
				controller.GetWorkToVoteByID(c)
			}
		})
	}

	return
}