package controller

import (
	"aubase/service"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
	"strconv"
)

func GetFile (c *gin.Context) {
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
}

func GetVoteStats (c *gin.Context) {
	turnID, _ := c.Get("turnID")
	votedInfos := service.GetVotedInfos(turnID.(uint32))
	c.JSON(http.StatusOK, gin.H{
		"msg": "查询成功！",
		"data": votedInfos,
	})
}

func GetOrder (c *gin.Context) {
	turnID, _ := c.Get("turnID")
	votedInfos := service.GetOrder(turnID.(uint32))
	c.JSON(http.StatusOK, gin.H{
		"msg": "查询成功！",
		"data": votedInfos,
	})
}