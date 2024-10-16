package router

import (
	"backed/api/handler"
	"backed/api/middleware"
	"github.com/gin-gonic/gin"
)

// ApplicationRouter 项目总路由
func ApplicationRouter(router *gin.Engine) {
	// 使用跨局中间件
	router.Use(middleware.CorsHandler)

	// 注册用户模块子路由
	userRouter(router.Group("/user"))

	// 注册主程序路由
	mainRouter(router.Group("/main"))

	// 注册第三方路由
	thirdPartyRouter(router.Group("/third_party"))

}

// userRouter 用户模块子路由
func userRouter(router *gin.RouterGroup) {
	// 登录接口
	router.POST("/login", handler.UserLoginHandler)
	// 注册接口
	router.POST("/register", handler.UserRegisterHandler)
	// 找回密码
	router.POST("/findPassword", handler.UserFindPasswordHandler)
	// 实名认证
	router.POST("/UserRealNameHandler", middleware.JwtTokenValid, handler.UserRealNameHandler)
	// 获取用户信息
	router.GET("/getUserInfo", middleware.JwtTokenValid, handler.UserInfoHandler)
	// 修改密码
	router.POST("/updatePassword", middleware.JwtTokenValid, handler.UserChangePasswordHandler)
	// 账号注销
	router.POST("/logout", middleware.JwtTokenValid, handler.UserLogoutHandler)
}

// mainRouter 主程序路由
func mainRouter(router *gin.RouterGroup) {
	// 检查用户是否购买过保险
	router.GET("/checkUserInsurance", middleware.JwtTokenValid, handler.CheckUserInsurance)
	// 取消购买保险
	router.POST("/cancelInsurance", middleware.JwtTokenValid, handler.CancelInsurance)
	// 获取购买人的信息
	router.GET("/getBuyerInfo", middleware.JwtTokenValid, handler.GetAllOrderUser)
	// 获取统计信息
	router.GET("/getStatistics", middleware.JwtTokenValid, handler.StatisticsInfo)
	// 获取交易信息
	router.GET("/getTransactionInfo", middleware.JwtTokenValid, handler.GetTransactionInfo)
	// 获取节点信息
	router.GET("/getNodeInfo", middleware.JwtTokenValid, handler.GetNodeInfo)
	// 获取所有用户信息
	router.GET("/getAllUserInfo", middleware.JwtTokenValid, handler.GetAllUserInfo)
	// 修改用户状态
	router.POST("/updateUserStatus", middleware.JwtTokenValid, handler.UpdateUserStatus)
	// 获取付款记录
	router.GET("/getPayRecord", middleware.JwtTokenValid, handler.GetPayRecord)
	// 创建套餐
	router.POST("/createCombo", middleware.JwtTokenValid, handler.CreateCombo)
	// 获取所有套餐
	router.GET("/getAllCombo", middleware.JwtTokenValid, handler.GetAllCombo)
	// 更新养老金状态
	router.POST("/updateComboStatus", middleware.JwtTokenValid, handler.UpdateComboStatus)
	// 获取所有套餐(用户)
	router.GET("/getAllComboUser", middleware.JwtTokenValid, handler.GetAllComboUser)
	// 获取套餐详情
	router.GET("/getComboDetail", middleware.JwtTokenValid, handler.GetComboDetail)
}

// third_party 第三方路由
func thirdPartyRouter(router *gin.RouterGroup) {
	// 购买养老保险接口
	router.POST("/buyCombo", middleware.JwtTokenValid, handler.BuyCombo)
	// 支付成功反馈
	router.GET("/callback", handler.AlipayCallback)
	// 支付提示
	router.POST("/notify", middleware.JwtTokenValid, handler.AlipayNotify)
}
