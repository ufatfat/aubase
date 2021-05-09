package service

import (
	"aubase/model"
	"strconv"
	"strings"
)

func CreateWork (workInfo *model.CreateWork) (workID uint32, err error) {
	if err = db.Table("work").Create(workInfo).Omit("work_group").Error; err != nil {
		return
	}
	workID = workInfo.WorkID
	return
}

func GetGroups (activityID uint32) (groups []model.WorkGroup) {
	db.Table("groups").Select("group_id", "group_name", "group_desc").Where("activity_id=?", activityID).Scan(&groups)
	return
}

func AddImageToDB (fileInfo *model.FileInfo) (ok bool) {
	end := strings.LastIndex(fileInfo.ImageName, ".")
	idx, _ := strconv.ParseUint(fileInfo.ImageName[:end], 10, 8)
	workID, _ := strconv.ParseUint(fileInfo.WorkID, 10, 32)
	if err := db.Table("images").Create(&model.ImageInfo{
		WorkID: uint32(workID),
		ImageIndex: uint8(idx),
		ImageUrl: "http://static.aubase.cn/" + fileInfo.ActivityID + "/" + fileInfo.WorkID + "/" + fileInfo.ImageName,
	}).Error; err == nil {
		return true
	}
	return
}