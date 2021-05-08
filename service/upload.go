package service

import "aubase/model"

func CreateWork (workInfo *model.CreateWork) (workID uint32, err error) {
	if err = db.Table("work").Create(workInfo).Error; err != nil {
		return
	}
	workID = workInfo.WorkID
	return
}

func GetGroups (activityID uint32) (groups []model.WorkGroup) {
	db.Table("groups").Select("group_id", "group_name", "group_desc").Where("activity_id=?", activityID).Scan(&groups)
	return
}