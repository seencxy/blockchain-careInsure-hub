package main

import (
	"backed/api/handler"
	"backed/api/router"
	"backed/app/core/variable"
	_ "backed/bootstrap" // 引入启动包
	"context"
	"github.com/gin-gonic/gin"
)

// 项目启动文件
func main() {
	// 初始化gin引擎
	engine := gin.Default()

	// 总路由
	router.ApplicationRouter(engine)

	// 等待用户购买订单
	go handler.WaitUserBuyCombo(context.TODO())
	// 监听养老金发放
	go handler.TransferMoney(context.TODO())

	// 监听路由
	engine.Run(":" + variable.Viper.GetString("application_port"))
}
