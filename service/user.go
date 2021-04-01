package service

import (
	"AUBase/model"
	"fmt"
)

func SignIn (username, password string) (userInfo model.UserInfo, err error) {
	if err = db.Table("users").Select("id", "name").Where("username=? and password=?", username, password).First(&userInfo).Error; err != nil {
		fmt.Println(err.Error())
	}
	return
}

func MyVotedWork (id uint32, positive bool) (work []model.VotedWork, err error) {
	if err = db.Table("images").Select("work_id", "image_url").Where("image_index=1 and work_id in (?) ", db.Table("results").Select("work_id").Where("is_positive=?", positive)).Find(&work).Error; err != nil {
		fmt.Println(err.Error())
	}
	return
}