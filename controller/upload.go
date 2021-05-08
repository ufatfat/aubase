package controller

import (
	"aubase/config"
	"aubase/model"
	"aubase/service"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func UploadImage (c *gin.Context) {
	activityID, _ := c.Get("activityID")
	w := c.Query("work")
	file, err := c.FormFile("upload")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	fileInfo := model.FileInfo{
		ActivityID: strconv.FormatUint(uint64(activityID.(uint32)), 10),
		WorkID: w,
		ImageName: file.Filename,
	}
	filename := fileInfo.ActivityID + "_" + fileInfo.WorkID + "_" + fileInfo.ImageName
	if err = c.SaveUploadedFile(file, "tmp/" + filename); err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	if err = uploadFilesToOSS(&fileInfo); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "上传成功！",
	})
}

func uploadFilesToOSS (fileInfo *model.FileInfo) (err error) {
	// 创建OSSClient实例。
	client, err := oss.New(config.OSS_ENDPOINT, config.OSS_KEY_ID, config.OSS_KEY_SECRET)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	// 获取存储空间。
	bucket, err := client.Bucket("2021aubase")
	if err != nil {
		fmt.Println("Error:", err)
	}

	filePath := "tmp/" + fileInfo.ActivityID + "_" + fileInfo.WorkID + "_" + fileInfo.ImageName
	// 上传本地文件。
	if err = bucket.PutObjectFromFile("/" + fileInfo.ActivityID + "/" + fileInfo.WorkID + "/" + fileInfo.ImageName, filePath); err != nil {
		fmt.Println(err.Error())
		return
	} else {
		os.Remove(filePath)
		if strings.Index(filePath, "compressed") == -1 {
			service.AddImageToDB(fileInfo)
		}
	}
	return
}

func GetGroups (c *gin.Context) {
	activityID, _ := c.Get("activityID")
	groups := service.GetGroups(activityID.(uint32))
	c.JSON(http.StatusOK, gin.H{
		"msg": "查询成功！",
		"data": groups,
	})
}

func CreateWork (c *gin.Context) {
	var info model.CreateWork
	if err := c.BindJSON(&info); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	workID, err := service.CreateWork(&info)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "新建作品成功！",
		"data": gin.H{
			"work_id": workID,
		},
	})
}