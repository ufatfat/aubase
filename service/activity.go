package service

import "aubase/model"

func GetTurnInfo (activityID uint32) (turnInfo model.TurnInfo, err error) {
	err = db.Table("turns").Select("turn_id", "turn_index", "is_positive", "votes_num").Where("activity_id=? and is_open=1", activityID).Find(&turnInfo).Error
	return
}