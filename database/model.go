/*
 * @Author: lin.zhenhui
 * @Date: 2020-03-21 11:46:01
 * @Last Modified by: lin.zhenhui
 * @Last Modified time: 2020-03-21 14:00:37
 */

package database

import "time"

// Model base model definition
type Model struct {
	ID          uint    `gorm:"primary_key"`
	CreatedTime uint
	UpdatedTime uint
	DeletedAt   *time.Time `sql:"index"`
}
