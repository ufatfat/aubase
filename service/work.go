package service

import (
	"aubase/model"
	"aubase/util"
	"gorm.io/gorm"
)

func GetWorkToVote (userID, activityID, turnID uint32) (workInfo model.WorkInfo, err error) {

	// 获取当前浏览作品ID
	var currentWorkID uint32
	if err = db.Table("votes").Select("current_work_id").Where("user_id =? and activity_id=? and turn_id=?", userID, activityID, turnID).Take(&currentWorkID).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return
		} else {
			db.Table("work").Select("work_id").Where("current_turn_index=(?)", db.Table("turns").Select("turn_index").Where("turn_id=?", turnID)).First(&currentWorkID)
			err = nil
		}
	}

	// 获取下一个作品的ID
	workRange := GetWorkRange(turnID)
	idx := util.GetIndexOfElem(&workRange, currentWorkID)

	// 当前轮次不存在此作品
	if idx == 0 {
		return model.WorkInfo{}, nil
	}
	// 当前轮次已浏览完
	if idx == len(workRange) - 1 {
		return model.WorkInfo{
			WorkGroup: "End",
		}, nil
	}
	nextWorkID := workRange[idx + 1]

	// 查询作品信息
	db.Table("work").Select("work_id", "work_group", "work_index").Where("work_id=? and activity_id=?", nextWorkID).First(&workInfo)

	workInfo.WorkImages = getWorkImages(currentWorkID + 1)
	return
}

func getWorkImages (workID uint32) (workImages []model.WorkImage) {
	db.Table("images").Select("image_id", "image_index", "image_url").Where("work_id=?", workID).Find(&workImages)
	return
}

func GetWorkRange (turnID uint32) (workRange []uint32) {
	db.Table("work").Select("work_id").Where("current_turn_index=(?)", db.Table("turns").Select("turn_index").Where("turn_id=?", turnID)).Scan(&workRange)
	return
}

func GetWorkNum (turnID uint32) (workNum uint32) {
	db.Table("work").Select("count(work_id) as num").Where("current_turn_index=(?)", db.Table("turns").Select("turn_index").Where("turn_id=?", turnID)).Take(&workNum)
	return
}