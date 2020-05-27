/*
 * @Author: lin.zhenhui
 * @Date: 2020-03-12 16:19:37
 * @Last Modified by: lin.zhenhui
 * @Last Modified time: 2020-03-12 16:32:20
 */

package model

import (
	"sso-api/database"
)

// User user
type User struct {
	DeletedAt
	ID              uint
	Number          string
	Password        string
	State           uint8
	Filter          database.User
}

// Create 新增
func (u *User) Create() {
	database.DB.Create(&u.Filter)
}

// Detail 详情
func (u *User) Detail() {
	database.DB.Where(&u.Filter).First(&u)
}

// Update 更新
func (u *User) Update() {
	database.DB.Model(&u.Filter).Updates(u.Filter)
}

// NumberExist 检查账号是否存在
func (u *User) NumberExist() bool {
	type user struct {
		ID     uint
		Number string
	}
	userInfo := user{ Number: u.Filter.Number }
	database.DB.Where(&userInfo).First(&userInfo)
	if userInfo.ID == 0 {
		return false
	}
	return true
}
