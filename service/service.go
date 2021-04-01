package service

import (
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
	dsn := 	mysql.Open("aubase:PivotStudio@2020@tcp(rm-2zeok8s8lj4322wc3do.mysql.rds.aliyuncs.com)/aubase?parseTime=true")
	db, _ = gorm.Open(dsn, &gorm.Config{
		Logger: newLogger,
	})
}