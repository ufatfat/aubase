package router

import (
	"aubase/controller"
	"aubase/middleware"
	"aubase/service"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/gin-gonic/gin"
	"reflect"
	"strconv"
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
		vote := api.Group("/vote")
		{
			vote.Use(middleware.IsSignedOut, middleware.IsTurnOpen)
			vote.PUT("", controller.VoteForWork)
			//vote.GET("/num",controller.GetVotedNum)
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
				default:
					controller.GetWorkToVoteByID(c)
				}
			})
		}
		activity := api.Group("/activity")
		{
			activity.Use(middleware.IsSignedOut)
			activity.GET("/turn", controller.GetTurnInfo)
		}
		upload := api.Group("/upload")
		{
			upload.POST("", controller.UploadImage)
			upload.POST("/info", controller.CreateWork)
			upload.GET("/group", controller.GetGroups)
		}
		api.GET("/stats", func (c *gin.Context) {
			activityID, _ := c.Get("activityID")
			workInfos := service.GetStats(activityID.(uint32))
			f := excelize.NewFile()
			idx := f.NewSheet("Sheet1")
			cols := "ABCDEFGHIJ"
			colsHeader := []string{"分类", "组别", "名称", "报名序列号", "负责人姓名", "负责人单位", "设计人员", "指导老师", "联系电话", "联系邮箱"}
			colsKey := []string{"Class", "WorkGroup", "WorkName", "SeqID", "LeaderName", "LeaderOrg", "Designers", "Teacher", "Phone", "Email"}
			for k := range cols {
				if err := f.SetCellValue("Sheet1", string(cols[k]) + "1", colsHeader[k]); err != nil {
					fmt.Println(err.Error())
				}
			}
			for i := range workInfos {
				for k := range colsKey {
					v := reflect.ValueOf(workInfos[i]).FieldByName(colsKey[k])
					if colsKey[k] == "Class" {
						fmt.Println(v, v.Uint())
						if v.Uint() == 0 {
							if err := f.SetCellValue("Sheet1", string(cols[k])+strconv.FormatInt(int64(i+2), 10), "高校"); err != nil {
								fmt.Println(err.Error())
							}
						} else {
							if err := f.SetCellValue("Sheet1", string(cols[k])+strconv.FormatInt(int64(i+2), 10), "社会"); err != nil {
								fmt.Println(err.Error())
							}
						}
					} else {
						if err := f.SetCellValue("Sheet1", string(cols[k])+strconv.FormatInt(int64(i+2), 10), v); err != nil {
							fmt.Println(err.Error())
						}
					}
				}
			}
			f.SetActiveSheet(idx)
			c.Header("content-disposition", `attachment; filename=stats.xlsx`)
			buf, _ := f.WriteToBuffer()
			c.Data(200, "application/octet-stream", buf.Bytes())
		})
	}


	return
}