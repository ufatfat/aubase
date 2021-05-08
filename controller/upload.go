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
)

func UploadImage (c *gin.Context) {
	form, _ := c.MultipartForm()
	for k, v := range form.File {
		fmt.Println("k:", k)
		for k1, v1 := range v {
			fmt.Println("k1:", k1, ", filename:", v1.Filename)
		}
	}
	/*form, err := c.MultipartForm()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	files := form.File["upload[]"]
	for k, v := range files {
		fmt.Println("k:", k)
		if err = c.SaveUploadedFile(v, "tmp/" + v.Filename); err != nil {
			fmt.Println(err.Error())
		}
	}*/
}

func uploadFilesToOSS () {
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
		os.Exit(-1)
	}

	// 上传本地文件。
	err = bucket.PutObjectFromFile("/", "")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
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