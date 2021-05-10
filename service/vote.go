package service

import (
	"aubase/model"
	"gorm.io/gorm"
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
	var curVotedWorkList string


	return nil
}