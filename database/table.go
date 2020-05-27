/*
 * @Author: lin.zhenhui
 * @Date: 2020-03-12 15:45:36
 * @Last Modified by: lin.zhenhui
 * @Last Modified time: 2020-03-14 20:05:52
 */

package database

// User 用户
type User struct {
	Model
	Number     string     `gorm:"unique;not null;comment:'账号'"`
	Password   string     `gorm:"not null;comment:'密码'"`
	State      uint8      `gorm:"DEFAULT:1;comment:'1: 正常；2: 禁用'"`
}

// UserLogin 用户登录记录
type UserLogin struct {
	Model
	UserID     uint       `gorm:"not null;comment:'账号id'"`
	Signature  string     `gorm:"comment:'JWT的签名'"`
	Exp        uint       `gorm:"comment:'过期时间'"`
	IP         string     `gorm:"comment:'IP地址'"`
	Country    string     `gorm:"comment:'IP地址地区'"`
	WebsiteID  uint       `gorm:"comment:'网站id'"`
	Message    string     `gorm:"comment:'一些信息，如登录错误'"`
	State      uint8      `gorm:"DEFAULT:1;comment:'1: 正常；2: 退出'"`
}

// Website 网站
type Website struct {
	Model
	Host       string     `gorm:"not null;comment:'网站'"`
	Title      string     `gorm:"comment:'网站名'"`
}

// autoMigrate 迁移表
func autoMigrate() {
	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&UserLogin{})
	DB.AutoMigrate(&Website{})
}
