package service

import (
	"AUBase/model"
	"fmt"
	"gorm.io/gorm"
)



func getWorkImages (workID uint64) (images []model.WorkImages, err error) {
	if err = db.Table("images").Select("image_index", "image_url").Where("work_id=?", workID).Find(&images).Error; err != nil {
		fmt.Println(err.Error())
	}
	return
}

func VoteForWork (workID uint64, userID uint32, negative bool) (err error) {
	var id uint64
	if err = db.Table("results").Select("id").Where("user_id=? and work_id=?", userID, workID).Take(&id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			db.Table("results").Create(map[string]interface{}{
				"work_id": workID,
				"user_id": userID,
				"is_negative": negative,
			})
		} else {
			fmt.Println(err.Error())
			return
		}
	} else {
		db.Table("results").Where("user_id=? and work_id=?", userID, workID).Update("is_negative", negative)
	}
	return nil
}

func GetVotedNum (userID, activityID, turnID uint32) (votedNum uint32, err error) {
	if err = db.Table("votes").Select("voted_num").Where("userID=? and activityID=? and turnID=?", userID, activityID, turnID).Take(&votedNum).Error; err == gorm.ErrRecordNotFound{
		err = nil
	}
	return
}