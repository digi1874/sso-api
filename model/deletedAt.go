/*
 * @Author: lin.zhenhui
 * @Date: 2020-03-21 11:51:19
 * @Last Modified by: lin.zhenhui
 * @Last Modified time: 2020-03-21 14:04:57
 */

package model

import (
	"time"
)

// DeletedAt DeletedAt
type DeletedAt struct {
	DeletedAt   *time.Time     `sql:"index" json:"-"`
}
