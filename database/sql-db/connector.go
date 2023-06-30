package database

import (
	"log"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/Lukmanern/go-starter/config"
)

var (
	gormDatabase     *gorm.DB
	gormDatabaseOnce sync.Once
)

func LoadGormDatabase() *gorm.DB {
	gormDatabaseOnce.Do(func() {
		var err error
		gormDatabase, err = gorm.Open(mysql.Open(config.Configuration().GetDatabaseURI()), &gorm.Config{})
		if err != nil {
			log.Panicf("cannot connect to MySQL %s", err)
		}
		gormDatabase.Logger = logger.Default.LogMode(logger.Info)
	})
	return gormDatabase
}
