/*
 * @Author: lin.zhenhui
 * @Date: 2020-03-14 15:54:43
 * @Last Modified by: lin.zhenhui
 * @Last Modified time: 2020-03-14 16:03:38
 */

package model

// List 列表
type List struct {
  Count           int        `json:"count"`
  Page            int        `json:"page"`
  Size            int        `json:"size"`
}