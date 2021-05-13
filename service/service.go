package service

import (
	"aubase/config"
	"aubase/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var db *gorm.DB

func init () {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,   // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,         // Disable color
		},
	)
	dsn := 	mysql.Open(config.DSN)
	db, _ = gorm.Open(dsn, &gorm.Config{
		Logger: newLogger,
	})
}

func CheckActivityOpen (activityID uint32) (ok bool) {
	db.Table("activities").Select("is_open").Where("activity_id=?", activityID).Take(&ok)
	return
}

func CheckTurnOpen (turnID uint32) (ok bool) {
	db.Table("turns").Select("is_open").Where("turn_id=?", turnID).Take(&ok)
	return
}

func GetStats (activityID uint32) (workInfos []model.CreateWork) {
	db.Table("work").Select("work.work_id", "groups.group_name as work_group", "work_name", "seq_id", "leader_name", "leader_org", "designers", "teacher", "phone", "email", "work_index", "class", "current_votes_num").Joins("left join `groups` on work.group_id=groups.group_id").Where("work.activity_id=?", activityID).Order("work.current_votes_num desc").Scan(&workInfos)
	return
}

func GenIndex (workID uint32, idx int) {
	db.Table("work").Where("work_id=?", workID).Update("work_index", idx+1)
}