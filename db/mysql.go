package db

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/cyansilver/go-lib/config"
)

// InitDB base on the config
func InitDB(cf *config.AppConfig) *gorm.DB {
	db, err := gorm.Open(mysql.Open(cf.DBUrl), &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(25)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(25)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(10 * time.Minute)

	return db
}
