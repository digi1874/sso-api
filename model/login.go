/*
 * @Author: lin.zhenhui
 * @Date: 2020-03-12 20:40:18
 * @Last Modified by: lin.zhenhui
 * @Last Modified time: 2020-03-18 13:00:34
 */

package model

import (
	"sso-api/database"
)

// UserLogin 登录记录
type UserLogin struct {
	DeletedAt
	ID          uint                 `json:"id"`
	UserID      uint                 `json:"-"`
	IP          string               `json:"ip"`
	Country     string               `json:"country"`
	Signature   string               `json:"signature"`
	Exp         uint                 `json:"exp"`
	WebsiteID   uint                 `json:"-"`
	Website     Website              `json:"website"`
	Message     string               `json:"message"`
	State       uint8                `json:"state"`
	CreatedTime uint                 `json:"createdTime"`
	Filter      database.UserLogin   `json:"-"`
}

// UserLoginList 登录记录列表
type UserLoginList struct {
	List
	Data      []UserLogin          `json:"data"`
	Filter    database.UserLogin   `json:"-"`
}

// Create 新增
func (ul *UserLogin) Create() {
	database.DB.Create(&ul.Filter)
}

// Detail 详情
func (ul *UserLogin) Detail() {
	database.DB.Where(&ul.Filter).First(&ul)
}

// Update 更新
func (ul *UserLogin) Update() {
	database.DB.Model(&ul.Filter).Updates(ul.Filter)
}

// Find 列表
func (ull *UserLoginList) Find() {
	// 不能查删除的
	ull.Filter.DeletedAt = nil
	db := database.DB.Where(&ull.Filter)

	if ull.Page < 1 {
		ull.Page = 1
	}
	if ull.Size < 1 {
		ull.Size = 20
	}
	offset := (ull.Page - 1) * ull.Size

	db.Order("Created_time desc").Limit(ull.Size).Offset(offset).Preload("Website").Find(&ull.Data)

	db.Model(&database.UserLogin{}).Count(&ull.Count)
}
