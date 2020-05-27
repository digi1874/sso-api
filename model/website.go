/*
 * @Author: lin.zhenhui
 * @Date: 2020-03-13 14:13:16
 * @Last Modified by: lin.zhenhui
 * @Last Modified time: 2020-03-18 12:46:54
 */

package model

import (
	"sso-api/database"
)

// Website website
type Website struct {
	DeletedAt
	ID              uint                 `json:"id"`
	Host            string               `json:"host"`
	Title           string               `json:"title"`
	Filter          database.Website     `json:"-"`
}

// WebsiteList website list
type WebsiteList struct {
	List
	Data            []Website            `json:"data"`
	Filter          database.Website     `json:"-"`
}

// FirstOrCreate Website detail or Create
func (w *Website) FirstOrCreate() {
	database.DB.Where(&w.Filter).FirstOrCreate(&w.Filter)
}

// Find 列表
func (wl *WebsiteList) Find() {
	db := database.DB

	if (wl.Filter.Host != "") {
		db = db.Where("`host` LIKE ?", "%" + wl.Filter.Host + "%")
		wl.Filter.Host = ""
	}

	wl.Filter.DeletedAt = nil
	db = db.Where(&wl.Filter)

	if wl.Page < 1 {
		wl.Page = 1
	}
	if wl.Size < 1 {
		wl.Size = 20
	}
	offset := (wl.Page - 1) * wl.Size

	db.Limit(wl.Size).Offset(offset).Find(&wl.Data)

	db.Model(&database.Website{}).Count(&wl.Count)
}
