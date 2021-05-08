package service

import (
	"aubase/model"
	"strconv"
	"strings"
)

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

func AddImageToDB (workID uint32, filename string) (ok bool) {
	begin := strings.LastIndex(filename, "/")
	end := strings.LastIndex(filename, ".")
	idx, _ := strconv.ParseUint(filename[begin:end], 10, 8)
	if err := db.Table("images").Create(&model.ImageInfo{
		WorkID: workID,
		ImageIndex: uint8(idx),
		ImageUrl: "http://static.aubase.cn" + filename,
	}).Error; err == nil {
		return true
	}
	return
}