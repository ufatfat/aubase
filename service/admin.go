package service

import (
	"aubase/model"
	"errors"
	"gorm.io/gorm"
	"strconv"
)

func GetVotedInfos (turnID uint32) (votedInfos []model.VotedInfo) {
	db.Table("votes").Select("votes.voted_work_ids", "users.name").Joins("left join users on votes.user_id=users.user_id").Where("turn_id=?", turnID).Scan(&votedInfos)
	return
}

func GetOrder (turnID uint32) (workInfos []model.WorkOrder) {
	db.Table("work").Select("work_id", "work_index", "work_name", "current_votes_num").Where("current_turn_index=(?)", db.Table("turns").Select("turn_index").Where("turn_id=?", turnID)).Order("current_votes_num desc").Scan(&workInfos)
	return
}

func AdminVote (userID, turnID, activityID uint32, workIdxs []string) (err error) {
	var ids []uint32
	db.Table("work").Select("work_id").Where("work_index in (?)", workIdxs).Scan(&ids)
	s := ""
	tx := db.Begin()
	for k := range ids {
		s += strconv.FormatUint(uint64(ids[k]), 10) + ";"
		if err = tx.Table("work").Where("work_id=?", ids[k]).Update("current_votes_num", gorm.Expr("current_votes_num+?", 1)).Error; err != nil {
			tx.Rollback()
			return
		}
	}
	var vid uint32
	if err = db.Table("votes").Select("vote_id").Where("user_id=? and turn_id=? and activity_id=?", userID, turnID, activityID).Take(&vid).Error; err == nil {
		return errors.New("record exists")
	} else {
		if err == gorm.ErrRecordNotFound {
			err = nil
		} else {
			return
		}
	}

	if err = tx.Table("votes").Create(map[string]interface{}{
		"user_id": userID,
		"turn_id": turnID,
		"activity_id": activityID,
		"voted_work_ids": s,
	}).Error; err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
	return
}