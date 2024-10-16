package models

import "gorm.io/gorm"

// UserInfo 存储用户信息
type UserInfo struct {
	gorm.Model
	Username string `json:"username"` // 用户名
	Password string `json:"password"` // 密码
	Name     string `json:"name"`     // 用户姓名
	Gender   string `json:"gender"`   // 用户性别
	Age      int64  `json:"age"`      // 用户年龄
	IdCard   string `json:"idCard"`   // 用户身份证号
	Avatar   string `json:"avatar"`   // 用户头像
	Status   bool   `json:"status"`   // 用户状态
}
