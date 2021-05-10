package service

import "aubase/model"

func GetVotedInfos (turnID uint32) (votedInfos []model.VotedInfo) {
	db.Table("votes").Select("votes.voted_work_ids", "users.name").Joins("left join users on votes.user_id=users.user_id").Where("turn_id=?", turnID).Scan(&votedInfos)
	return
}

func GetOrder (turnID uint32) (workInfos []model.WorkOrder) {
	db.Table("work").Select("work_id", "work_index", "work_name", "current_votes_num").Where("current_turn_index=(?)", db.Table("turns").Select("turn_index").Where("turn_id=?", turnID)).Order("current_votes_num desc").Scan(&workInfos)
	return
}