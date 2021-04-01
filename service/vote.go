package service

import (
	"AUBase/model"
	"fmt"
	"gorm.io/gorm"
)

func GetWorkToVote (userID uint32) (work model.WorkToVote, err error) {
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

func getWorkImages (workID uint64) (images []model.WorkImages, err error) {
	if err = db.Table("images").Select("image_index", "image_url").Where("work_id=?", workID).Find(&images).Error; err != nil {
		fmt.Println(err.Error())
	}
	return
}

func VoteForWork (workID uint64, userID uint32, positive bool) (err error) {
	var id uint64
	if err = db.Table("results").Select("id").Where("user_id=? and work_id=?", userID, workID).Take(&id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			db.Table("results").Create(map[string]interface{}{
				"work_id": workID,
				"user_id": userID,
				"is_positive": positive,
			})
		} else {
			fmt.Println(err.Error())
			return
		}
	} else {
		err = db.Table("results").Where("user_id=? and work_id=?", userID, workID).Update("is_positive", positive).Error
	}
	return
}