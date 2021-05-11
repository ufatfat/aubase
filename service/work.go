package service

import (
	"aubase/model"
	"aubase/util"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

func GetWorkToVote (userID, activityID, turnID uint32) (workInfo model.WorkInfo, err error) {

	// 获取当前浏览作品ID
	var currentWorkID uint32
	if err = db.Table("votes").Select("current_work_id").Where("user_id =? and activity_id=? and turn_id=?", userID, activityID, turnID).Take(&currentWorkID).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return
		} else {
			err = nil
		}
	}

	var nextWorkID uint32
	// 获取下一个作品的ID
	workRange := GetWorkRange(turnID)
	if currentWorkID == 0 {
		nextWorkID = workRange[0]
	} else {
		idx := util.GetIndexOfElem(&workRange, currentWorkID)
		// 当前轮次不存在此作品
		if idx == -1 {
			return model.WorkInfo{
				WorkGroup: "No",
			}, nil
		}
		// 当前轮次已浏览完
		if idx == len(workRange) - 1 {
			return model.WorkInfo{
				WorkGroup: "End",
			}, nil
		}
		nextWorkID = workRange[idx + 1]
	}

	// 查询作品信息
	db.Table("work").Select("work.work_id", "groups.group_name as work_group", "work_index").Joins("left join groups on work.group_id=groups.group_id").Where("work_id=? and work.activity_id=?", nextWorkID, activityID).First(&workInfo)

	workInfo.WorkImages = getWorkImages(nextWorkID)
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

func GetWorkToVoteByID (userID, turnID, workID uint32) (workInfo model.WorkInfo, err error) {
	if err = db.Table("work").Select("work.work_id", "groups.group_name as work_group", "work_index").Joins("left join groups on work.group_id=groups.group_id").Where("work_id=? and current_turn_index=(?)", workID, db.Table("turns").Select("turn_index").Where("turn_id=?", turnID)).First(&workInfo).Error; err != nil {
		return
	}
	workInfo.WorkImages = getWorkImages(workID)
	workInfo.IsVoted = checkIsVoted(userID, turnID, workID)
	return
}

func getVotedWorkList (userID, turnID uint32) map[string]bool {
	var v string
	if err := db.Table("votes").Select("voted_work_ids").Where("user_id=? and turn_id=?", userID, turnID).Take(&v).Error; err == gorm.ErrRecordNotFound {
		return nil
	}
	b := strings.Split(v, ";")
	t := make(map[string]bool)
	for i := 0; i < len(b); i++ {
		t[b[i]] = true
	}
	delete(t, "")
	return t
}

func checkIsVoted (userID, turnID, workID uint32) bool {
	votedWorkList := getVotedWorkList(userID, turnID)
	w := strconv.FormatUint(uint64(workID), 10)
	return votedWorkList[w]
}

func GetWorkByGroup (groupID, userID, turnID uint32, get string) (workInfos []model.WorkInfo, err error) {
	switch get {
	case "all":
		if groupID == 0 {
			db.Table("work").Select("work.work_id", "groups.group_name as work_group", "work_index").Joins("left join groups on work.group_id=groups.group_id").Where("current_turn_index=(?)", db.Table("turns").Select("turn_index").Where("turn_id=?", turnID)).Scan(&workInfos)
		} else {
			db.Table("work").Select("work.work_id", "groups.group_name as work_group", "work_index").Joins("left join groups on work.group_id=groups.group_id").Where("current_turn_index=(?) and work.group_id=?", db.Table("turns").Select("turn_index").Where("turn_id=?", turnID), groupID).Scan(&workInfos)
		}
		for k := range workInfos {
			workInfos[k].WorkImages = getWorkImages(workInfos[k].WorkID)
		}
	case "voted":
		var v string
		db.Table("votes").Select("voted_work_ids").Where("user_id=? and turn_id=?", userID, turnID).Take(&v)
		b := strings.Split(v, ";")
		if groupID == 0 {
			db.Table("work").Select("work.work_id", "groups.group_name as work_group", "work_index").Joins("left join groups on work.group_id=groups.group_id").Where("current_turn_index=(?) and work.work_id in (?)", db.Table("turns").Select("turn_index").Where("turn_id=?", turnID), b).Scan(&workInfos)
		} else {
			db.Table("work").Select("work.work_id", "groups.group_name as work_group", "work_index").Joins("left join groups on work.group_id=groups.group_id").Where("current_turn_index=(?) and work.group_id=? and work.work_id in (?)", db.Table("turns").Select("turn_index").Where("turn_id=?", turnID), groupID, b).Scan(&workInfos)
		}
		for k := range workInfos {
			workInfos[k].WorkImages = getWorkImages(workInfos[k].WorkID)
		}
	case "not_voted":
		var v string
		db.Table("votes").Select("voted_work_ids").Where("user_id=? and turn_id=?", userID, turnID).Take(&v)
		b := strings.Split(v, ";")
		if groupID == 0 {
			db.Table("work").Select("work.work_id", "groups.group_name as work_group", "work_index").Joins("left join groups on work.group_id=groups.group_id").Where("current_turn_index=(?) and work.work_id not in (?)", db.Table("turns").Select("turn_index").Where("turn_id=?", turnID), b).Scan(&workInfos)
		} else {
			db.Table("work").Select("work.work_id", "groups.group_name as work_group", "work_index").Joins("left join groups on work.group_id=groups.group_id").Where("current_turn_index=(?) and work.group_id=? and work.work_id not in (?)", db.Table("turns").Select("turn_index").Where("turn_id=?", turnID), groupID, b).Scan(&workInfos)
		}
		for k := range workInfos {
			workInfos[k].WorkImages = getWorkImages(workInfos[k].WorkID)
		}
	}
	return
}

func GetWorkIDByWorkIndex (idx uint16) (workID uint32, err error) {
	err = db.Table("work").Select("work_id").Where("work_index=?", idx).Take(&workID).Error
	return
}