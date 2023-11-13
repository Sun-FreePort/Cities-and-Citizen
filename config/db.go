package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"sync"
)

var onceDB sync.Once

// type global
var (
	db *gorm.DB
)

func GetDB() *gorm.DB {
	getDB("")
	return db
}

func GetDBByPath(projectPath string) *gorm.DB {
	getDB(projectPath)
	return db
}

func getDB(projectPath string) {
	onceDB.Do(func() {
		// 生成配置
		configKV := GetConfig(projectPath)
		params := MySQLParams{
			Host:      configKV["DB_HOST"],
			Port:      configKV["DB_PORT"],
			Username:  configKV["DB_USERNAME"],
			Password:  configKV["DB_PASSWORD"],
			Database:  configKV["DB_DATABASE"],
			ParseTime: configKV["DB_PARSE_TIME"],
		}

		// 生成对象
		dbConn, err := gorm.Open(mysql.Open(fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?parseTime=%s",
			params.Username,
			params.Password,
			params.Host,
			params.Port,
			params.Database,
			params.ParseTime)), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}
		db = dbConn
	})
}

type MySQLParams struct {
	Host      string
	Port      string
	Username  string
	Password  string
	Database  string
	ParseTime string
}

func TestDB() *gorm.DB {
	return db
}

func DropTestDB() error {
	if err := os.Remove("./../database/realworld_test.db"); err != nil {
		return err
	}
	return nil
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(
	//&model.User{},
	//&model.Follow{},
	//&model.Article{},
	//&model.Comment{},
	//&model.Tag{},
	)
}
