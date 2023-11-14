package config

import (
	"fmt"
	"github.com/Sun-FreePort/Cities-and-Citizen/model"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
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

		var err error

		if configKV["ENV"] == "test" {
			dsn := "./../database/test_sqlite.db"
			db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
		} else {
			params := MySQLParams{
				Host:      configKV["DB_HOST"],
				Port:      configKV["DB_PORT"],
				Username:  configKV["DB_USERNAME"],
				Password:  configKV["DB_PASSWORD"],
				Database:  configKV["DB_DATABASE"],
				ParseTime: configKV["DB_PARSE_TIME"],
			}

			// 生成对象
			db, err = gorm.Open(mysql.Open(fmt.Sprintf(
				"%s:%s@tcp(%s:%s)/%s?parseTime=%s",
				params.Username,
				params.Password,
				params.Host,
				params.Port,
				params.Database,
				params.ParseTime)), &gorm.Config{})
		}
		if err != nil {
			panic(fmt.Sprintf("failed to connect database: %s", err))
		}
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

func DropTestDB() error {
	if err := os.Remove("./../database/test_sqlite.db"); err != nil {
		return err
	}
	return nil
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&model.UserModel{},
		//&model.Follow{},
		//&model.Article{},
		//&model.Comment{},
		//&model.Tag{},
	)
}
