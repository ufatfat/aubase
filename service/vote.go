package service

import (
	"aubase/model"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

func VoteForWork (activityID, userID, turnID uint32, voteInfo model.VoteInfo) (err error) {
	var vID uint32
	var votesNum uint16
	if err = db.Table("turns").Select("votes_num").Where("turn_id=?", turnID).Take(&votesNum).Error; err != nil {
		return
	}
	createVote := model.CreateVote{
		CurrentWorkID: voteInfo.CurrentWorkID,
		UserID:        userID,
		ActivityID:    activityID,
		TurnID:        turnID,
		VotesNum:      votesNum,
	}
	if err = db.Table("votes").Select("vote_id").Where("activity_id=? and user_id=? and turn_id=?", activityID, userID, turnID).Take(&vID).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return
		} else {
			db.Table("votes").Create(&createVote)
		}
	}
	tx := db.Begin()
	var curVoteInfo model.CurVoteInfo
	db.Table("votes").Select("voted_work_ids, current_work_id").Where("user_id=? and turn_id=?", userID, turnID).Take(&curVoteInfo)
	m := strToMap(curVoteInfo.VotedWorkIDs)
	k := strconv.FormatUint(uint64(voteInfo.CurrentWorkID), 10)
	if voteInfo.IsVoted {
		if !m[k] {
			m[k] = true
			if err = tx.Table("work").Where("work_id=?", voteInfo.CurrentWorkID).Update("current_votes_num", gorm.Expr("current_votes_num+?", 1)).Error; err != nil {
				tx.Rollback()
				return
			}
			if err = tx.Table("votes").Where("user_id=? and turn_id=?", userID, turnID).Update("voted_num", gorm.Expr("voted_num+?", 1)).Error; err != nil {
				tx.Rollback()
				return
			}
		}
	} else {
		if m[k] {
			delete(m, strconv.FormatUint(uint64(voteInfo.CurrentWorkID), 10))
			if err = tx.Table("work").Where("work_id=?", voteInfo.CurrentWorkID).Update("current_votes_num", gorm.Expr("current_votes_num-?", 1)).Error; err != nil {
				tx.Rollback()
				return
			}
			if err = tx.Table("votes").Where("user_id=? and turn_id=?", userID, turnID).Update("voted_num", gorm.Expr("voted_num-?", 1)).Error; err != nil {
				tx.Rollback()
				return
			}
		}
	}
	curVoteInfo.VotedWorkIDs = mapToStr(m)
	if err = tx.Table("votes").Where("user_id=? and turn_id=?", userID, turnID).Update("voted_work_ids", curVoteInfo.VotedWorkIDs).Error; err != nil {
		tx.Rollback()
		return
	}
	if voteInfo.CurrentWorkID > curVoteInfo.CurrentVoteID {
		if err = tx.Table("votes").Where("user_id=? and turn_id=?", userID, turnID).Update("current_work_id", voteInfo.CurrentWorkID).Error; err != nil {
			tx.Rollback()
			return
		}
	}
	tx.Commit()
	return nil
}

func strToMap (votedWorkList string) map[string]bool {
	t := strings.Split(votedWorkList, ";")
	l := make(map[string]bool)
	for k := range t {
		l[t[k]] = true
	}
	delete(l, "")
	return l
}

func mapToStr (votedWorkList map[string]bool) string {
	s := ""
	for k := range votedWorkList {
		s += k + ";"
	}
	return s
}

func VoteDone (userID, turnID uint32) (err error) {
	err = db.Table("votes").Where("user_id=? and turn_id=?", userID, turnID).Update("is_done", 1).Error
	return
}

func CheckIsDone (userID, turnID uint32) (ok bool) {
	db.Table("votes").Select("is_done").Where("user_id=? and turn_id=?", userID, turnID).Take(&ok)
	return
}

func GetVotedNum (userID, turnID uint32) (num uint16) {
	db.Table("votes").Select("voted_num").Where("user_id=? and turn_id=?", userID, turnID).Take(&num)
	return
}