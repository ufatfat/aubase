package service

import (
	"AUBase/model"
	"fmt"
	"gorm.io/gorm"
)

func GetWorkNum () (num uint64) {
	db.Table("work").Select("count(id) as num").Take(&num)
	return
}

func GetWorkToVoteByUserID (userID uint32) (work model.WorkToVote, err error) {
	var workInfo model.WorkInfo
	if err = db.Table("work").Select("id", "work_index").Where("id not in (?)", db.Table("results").Select("work_id").Where("user_id=?", userID)).First(&workInfo).Error; err != nil {
		fmt.Println(err.Error())
		return
	}
	work.Images, err = getWorkImages(workInfo.ID)
	work.ID = workInfo.ID
	work.WorkIndex = workInfo.WorkIndex
	return
}

func GetWorkInfo (workID uint64, userID uint32) (work model.WorkToVote, err error) {
	var workInfo model.WorkInfo
	if err = db.Table("work").Select("id", "work_index" ).Where("id=?", workID).First(&workInfo).Error; err != nil {
		fmt.Println(err.Error())
		return
	}
	if err = db.Table("results").Select("is_negative").Where("user_id=? and work_id=?", userID, workID).Take(&work.IsNegative).Error; err != nil && err != gorm.ErrRecordNotFound {
		fmt.Println(err.Error())
		return
	}
	work.Images, err = getWorkImages(workInfo.ID)
	work.ID = workInfo.ID
	work.WorkIndex = workInfo.WorkIndex
	return
}