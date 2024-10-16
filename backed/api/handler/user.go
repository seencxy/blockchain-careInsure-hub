package handler

import (
	"backed/app/core/consts"
	"backed/app/models"
	"backed/app/utils"
	"backed/internal/database"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// UserLoginHandler 用户登录接口层
func UserLoginHandler(ctx *gin.Context) {
	// 从请求中获取参数
	var requestData models.UserInfo
	// 绑定请求参数
	err := ctx.ShouldBindJSON(&requestData)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取请求参数失败:" + err.Error(),
		})
		return
	}

	// 进行逻辑操作
	userDao := database.UserInfo
	// 判断用户登录账号和密码是否正确
	users, err := userDao.WithContext(ctx).
		Where(userDao.Username.Eq(requestData.Username)).
		Where(userDao.Password.Eq(requestData.Password)).
		First()
	if err != nil && err.Error() != "record not found" {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取用户信息失败:" + err.Error(),
		})
		return
	}

	// 如果count不等0则代表用户登录成功
	if users == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "用户名或密码错误",
		})
		return
	}

	// 生成token
	token, err := utils.GenToken(requestData.Username, consts.TokenExpireDuration, []byte(consts.TokenSecret))
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "生成token失败: " + err.Error(),
		})
		return
	}

	// 判断用户是否认证
	flag := true
	if len(users.IdCard) != 18 || len(users.Name) == 0 {
		flag = false
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "登录成功! ",
		"token":   token, // 用户登录auth_token
		"flag":    flag,  // 用户账号是否认证
		"avatar":  users.Avatar,
	})
}

// UserRegisterHandler 用户注册接口层
func UserRegisterHandler(ctx *gin.Context) {
	// 从请求中获取参数
	var requestData models.UserInfo
	// 绑定请求参数
	err := ctx.ShouldBindJSON(&requestData)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取请求参数失败:" + err.Error(),
		})
		return
	}

	requestData.Status = true

	// 实现注册逻辑
	userDao := database.UserInfo

	// 先判断用户名是否存在
	count, err := userDao.WithContext(ctx).Where(userDao.Username.Eq(requestData.Username)).Count()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取用户信息失败:" + err.Error(),
		})
		return
	}

	if count != 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "用户名已存在",
		})
		return
	}

	// 设置默认头像
	requestData.Avatar = consts.DefaultAvatar

	// 注册用户
	err = userDao.WithContext(ctx).Create(&requestData)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "用户注册失败:" + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "正常",
	})
}

// UserFindPasswordHandler 找回密码接口层
func UserFindPasswordHandler(ctx *gin.Context) {
	// 从请求中获取参数
	var requestData models.UserInfo
	// 绑定请求参数
	err := ctx.ShouldBindJSON(&requestData)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取请求参数失败:" + err.Error(),
		})
		return
	}

	// 通过用户的身份信息获取用户的密码
	userDao := database.UserInfo
	userInfo, err := userDao.WithContext(ctx).
		Where(userDao.Name.Eq(requestData.Name)).
		Where(userDao.IdCard.Eq(requestData.IdCard)).
		First()
	if err != nil && err.Error() != "record not found" {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取用户信息失败:" + err.Error(),
		})
		return
	}

	if userInfo == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "没有账号绑定用户信息，请重新注册",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":     200,
		"message":  "获取密码成功",
		"username": userInfo.Username,
		"password": userInfo.Password,
	})
}

// UserRealNameHandler 实名认证接口层
func UserRealNameHandler(ctx *gin.Context) {
	// 从上下文中获取token
	username, err := utils.GetUserInfoFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取用户信息失败:" + err.Error(),
		})
		return
	}

	if len(username) == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "用户信息出错，请重新登录",
		})
		return
	}

	// 获取该用户的信息
	userDao := database.UserInfo
	userInfo, err := userDao.WithContext(ctx).Where(userDao.Username.Eq(username)).First()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取用户信息出错:" + err.Error(),
		})
		return
	}

	if len(userInfo.IdCard) == 18 || len(userInfo.Name) != 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "用户已经绑定了身份信息",
		})
		return
	}

	// 从请求中获取参数
	var requestData models.UserInfo
	// 绑定请求参数
	err = ctx.ShouldBindJSON(&requestData)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取请求参数失败:" + err.Error(),
		})
		return
	}

	// 检验身份证是否合法
	flag := utils.CheckIDCard(requestData.IdCard)
	if !flag {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "身份证不合法!",
		})
		return
	}

	// 检查有没有其他用户绑定了该身份证
	count, err := userDao.WithContext(ctx).
		Where(userDao.IdCard.Eq(requestData.IdCard)).
		Count()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "查询身份证信息出错:" + err.Error(),
		})
		return
	}

	if count != 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "身份证已被绑定，请利用身份证找回",
		})
		return
	}

	// 根据身份证号来获取用户的性别以及年龄
	// 获取身份证的倒数一位数
	last := requestData.IdCard[len(requestData.IdCard)-2 : len(requestData.IdCard)-1]
	s, err := strconv.Atoi(last)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "判断性别出错: " + err.Error(),
		})
		return
	}

	if s%2 == 0 {
		requestData.Gender = "女"
	} else {
		requestData.Gender = "男"
	}

	// 获取年龄
	age, err := utils.CalculateAge(requestData.IdCard)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取年龄出错: " + err.Error(),
		})
		return
	}

	requestData.Age = int64(age)

	// 更新用户信息
	_, err = userDao.WithContext(ctx).
		Where(userDao.Username.Eq(username)).
		Updates(&requestData)

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "更新用户信息失败:" + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "实名认证成功",
	})
}

// UserInfoHandler 获取用户信息接口层
func UserInfoHandler(ctx *gin.Context) {
	// 从上下文中获取token
	username, err := utils.GetUserInfoFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取用户信息失败:" + err.Error(),
		})
		return
	}

	if len(username) == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "用户信息出错，请重新登录",
		})
		return
	}

	// 查询用户信息
	userDao := database.UserInfo
	userInfo, err := userDao.WithContext(ctx).
		Where(userDao.Username.Eq(username)).
		First()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "查询用户信息出错: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":     200,
		"message":  "获取用户信息成功",
		"userInfo": userInfo,
	})
}

// UserChangePasswordHandler 修改用户密码以及头像
func UserChangePasswordHandler(ctx *gin.Context) {
	// 从上下文中获取token
	username, err := utils.GetUserInfoFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取用户信息失败:" + err.Error(),
		})
		return
	}

	if len(username) == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "用户信息出错，请重新登录",
		})
		return
	}

	// 从请求中获取参数
	var requestData models.UserInfo
	// 绑定请求参数
	err = ctx.ShouldBindJSON(&requestData)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取请求参数失败:" + err.Error(),
		})
		return
	}

	// 更新用户信息
	userDao := database.UserInfo
	_, err = userDao.WithContext(ctx).
		Where(userDao.Username.Eq(username)).
		Updates(&requestData)

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "更新用户信息失败:" + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "修改用户信息成功",
	})
}

// UserLogoutHandler 用户注销接口层
func UserLogoutHandler(ctx *gin.Context) {
	// 从上下文中获取token
	username, err := utils.GetUserInfoFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取用户信息失败:" + err.Error(),
		})
		return
	}

	if len(username) == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "用户信息出错，请重新登录",
		})
		return
	}

	// 删除用户信息
	userDao := database.UserInfo
	_, err = userDao.WithContext(ctx).Where(userDao.Username.Eq(username)).Delete()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "注销用户失败:" + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "注销成功",
	})
}
