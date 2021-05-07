package service

import (
	"aubase/config"
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