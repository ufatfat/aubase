package service

import (
	"aubase/model"
	"fmt"
)

func SignIn (username, password string) (userInfo model.UserInfo, err error) {
	if err = db.Table("users").Select("user_id", "name").Where("username=? and password=?", username, password).First(&userInfo).Error; err != nil {
		fmt.Println(err.Error())
	}
	return
}

/*func MyVotedWork (id uint32) (work []model.VotedWork, err error) {
	if err = db.Table("work").Select("work_id").Where("id not in (?) ", db.Table("results").Select("work_id").Where("is_negative=1 and user_id=?", id)).Find(&work).Error; err != nil {
		fmt.Println(err.Error())
		return
	}
	var temp []model.VotedWork
	if err = db.Table("results").Select("work_id", "is_negative").Where("is_negative=1 and user_id=?", id).Find(&temp).Error; err != nil {
		fmt.Println(err.Error())
		return
	}
	for i := 0; i < len(temp); i ++ {
		work = append(work, temp[i])
	}
	return
}*/