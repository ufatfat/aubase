package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UploadImage (c *gin.Context) {
	form, err := c.MultipartForm()
	type test struct {
		Test	string `json:"test"`
	}
	var t test
	if err = c.BindJSON(&t); err != nil {
		fmt.Println("bind:", err.Error())
	}
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
	}
}

func CreateWork (c *gin.Context) {

}

func uploadFilesToOSS () {

}