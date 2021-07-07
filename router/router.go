package router

import (
	"aubase/controller"
	"aubase/middleware"
	"aubase/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter () (r *gin.Engine) {
	r = gin.Default()
	r.GET("/index", func (c *gin.Context) {
		l := service.GetWorkRange(1)
		for k := range l {
			service.GenIndex(l[k], k)
		}
	})
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
		vote := api.Group("/vote")
		{
			vote.Use(middleware.IsSignedOut, middleware.IsTurnOpen)
			vote.PUT("", controller.VoteForWork)
			vote.POST("/done", controller.VoteDone)
			vote.GET("/done", controller.CheckIsDone)
			vote.GET("/num",controller.GetVotedNum)
		}
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
				case "group":
					controller.GetWorkByGroup(c)
				case "search":
					controller.GetWorkIDByWorkIndex(c)
				default:
					controller.GetWorkToVoteByID(c)
				}
			})
		}
		activity := api.Group("/activity")
		{
			activity.GET("/turn", controller.GetTurnInfo)
		}
		upload := api.Group("/upload")
		{
			upload.POST("", controller.UploadImage)
			upload.POST("/info", controller.CreateWork)
			upload.GET("/group", controller.GetGroups)
		}
		admin := api.Group("/admin")
		{
			admin.GET("/file", controller.GetFile)
			admin.GET("/votes", middleware.IsTurnOpen, controller.GetVoteStats)
			admin.GET("/order", middleware.IsTurnOpen, controller.GetOrder)
			admin.POST("/votes/:userID", middleware.IsTurnOpen, controller.AdminVote)
			admin.GET("/users", controller.GetUsers)
		}
		api.GET("/test", func (c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"data": service.GetImages(),
			})
		})

	}

	return
}