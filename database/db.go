/*
 * @Author: lin.zhenhui
 * @Date: 2020-03-20 11:04:51
 * @Last Modified by: lin.zhenhui
 * @Last Modified time: 2020-03-21 14:03:27
 */

package database

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	// 导入数据库的驱动
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"sso-api/process"
)

// DBConfig 数据库配置
type DBConfig struct {
	User            string
	Password        string
	DatabaseName    string
	Localhost       string
}

// DB 数据库
var DB *gorm.DB

func init()  {
	err := run()
	if err != nil {
		fmt.Println("database ERROR:", err)
		os.Exit(1)
	}

	if process.IsDev {
		DB.LogMode(true)
	}

	autoMigrate()
}

func run() error {
	var byteConfig []byte
	var err error

	// 读取配置文件
	byteConfig, err = ioutil.ReadFile("./db.json")
	if err != nil {
		return err
	}

	var dbConfig DBConfig
	// 解析配置数据
	err = json.Unmarshal(byteConfig, &dbConfig)
	if err != nil {
		return err
	}

	// 连接数据库
	DB, err = gorm.Open(
		"mysql",
		dbConfig.User +
			":" +
			dbConfig.Password +
			"@(" +
			dbConfig.Localhost +
			")/" +
			dbConfig.DatabaseName +
			"?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		return err
	}

	DB.Callback().Create().Replace("gorm:update_time_stamp",updateTimeStampForCreateCallback)
	DB.Callback().Update().Replace("gorm:update_time_stamp",updateTimeStampForUpdateCallback)

	return nil
}

// 注册新建钩子在持久化之前
func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now().Unix()
		if createTimeField, ok := scope.FieldByName("CreatedTime"); ok {
			if createTimeField.IsBlank {
				createTimeField.Set(nowTime)
			}
		}

		if modifyTimeField, ok := scope.FieldByName("UpdatedTime"); ok {
			if modifyTimeField.IsBlank {
				modifyTimeField.Set(nowTime)
			}
		}
	}
}

// 注册更新钩子在持久化之前
func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		scope.SetColumn("UpdatedTime", time.Now().Unix())
	}
}
