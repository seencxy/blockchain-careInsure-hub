package main

import (
	"backed/app/models"
	"backed/internal/pkg"
)

// 生成 gorm 基础mysql框架代码
func main() {
	generator := pkg.NewGenerator("database")
	generator.ApplyBasic(&models.UserInfo{})
	generator.ApplyBasic(&models.AlipayRecord{})
	generator.ApplyBasic(&models.TransferAccount{})
	generator.ApplyBasic(&models.ComboDescription{})
	generator.ApplyBasic(&models.Combo{})
	generator.Execute()
}
