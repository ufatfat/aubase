package service

import (
	"aubase/model"
	"aubase/util"
	"fmt"
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
			db.Table("work").Select("work_id").Where("current_turn_index=(?)", db.Table("turns").Select("turn_index").Where("turn_id=?", turnID)).First(&currentWorkID)
			err = nil
		}
	}
	fmt.Println("currentWorkID:", currentWorkID)

	// 获取下一个作品的ID
	workRange := GetWorkRange(turnID)
	idx := util.GetIndexOfElem(&workRange, currentWorkID)

	fmt.Println("idx:", idx)
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
	nextWorkID := workRange[idx + 1]

	// 查询作品信息
	db.Table("work").Select("work_id", "work_group", "work_index").Where("work_id=? and activity_id=?", nextWorkID, activityID).First(&workInfo)

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
	if err = db.Table("work").Select("work_id", "work_index", "work_group").Where("work_id=? and current_turn_index=(?)", workID, db.Table("turns").Select("turn_index").Where("turn_id=?", turnID)).First(&workInfo).Error; err != nil {
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