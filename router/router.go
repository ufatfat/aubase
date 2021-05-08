package router

import (
	"aubase/controller"
	"aubase/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter () (r *gin.Engine) {
	r = gin.Default()
	api := r.Group("/api")
	{
		api.Use(middleware.IsActivityOpen)
		user := api.Group("/user")
		{
			user.POST("/signin", middleware.IsSignedIn, controller.SignIn)
			user.Use(middleware.IsSignedOut)
			user.POST("/signout", controller.SignOut)
			//user.PUT("/password", controller.ChangePassword)
			//user.GET("/myvote", controller.MyVotedWork)
		}
		/*vote := api.Group("/vote")
		{
			vote.Use(middleware.IsSignedOut)
			vote.PUT("", controller.VoteForWork)
			vote.GET("/num",controller.GetVotedNum)
		}*/
		work := api.Group("/work")
		{
			work.Use(middleware.IsSignedOut, middleware.IsTurnOpen)
			work.GET("", controller.GetWorkToVote)
			work.GET("/:workID", func (c *gin.Context) {
				p := c.Param("workID")
				switch p {
				case "num":
					controller.GetWorkNum(c)
				case "range":
					controller.GetWorkRange(c)
				default:
					controller.GetWorkToVoteByID(c)
				}
			})
		}/**/
		activity := api.Group("/activity")
		{
			activity.Use(middleware.IsSignedOut)
			activity.GET("/turn", controller.GetTurnInfo)
		}
		upload := api.Group("upload")
		{
			//upload.Use(middleware.IsSignedOut)
			upload.POST("", controller.UploadImage)
		}
	}

	return
}