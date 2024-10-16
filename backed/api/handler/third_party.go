package handler

import (
	"backed/app/core/consts"
	"backed/app/core/variable"
	"backed/app/models"
	"backed/app/utils"
	"backed/internal/database"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/smartwalle/alipay/v3"
	"github.com/smartwalle/xid"
	"log"
	"math/big"
	"net/http"
	"strconv"
	"time"
)

// BuyCombo 购买养老保险接口
func BuyCombo(ctx *gin.Context) {
	// 生产环境用1块钱代替一万
	var userInfo models.UserInfo

	err := ctx.ShouldBindJSON(&userInfo)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "绑定请求参数失败:" + err.Error(),
		})
		return
	}

	// 获取购买的套餐
	productDao := database.Combo
	first, err2 := productDao.Where(productDao.ID.Eq(userInfo.ID)).First()
	if err2 != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取套餐信息失败:" + err2.Error(),
		})
		return
	}
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
	// 检查用户是否是否认证
	userDao := database.UserInfo
	// 获取用户信息
	userInfos, err := userDao.WithContext(ctx).Where(userDao.Username.Eq(username)).First()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取用户信息失败:" + err.Error(),
		})
		return
	}

	if len(userInfos.IdCard) != 18 {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "用户没有实名认证，请先实名认证",
		})
		return
	}

	// 生成订单
	var tradeNo = fmt.Sprintf("%d", xid.Next())

	var p = alipay.TradePagePay{}
	p.NotifyURL = variable.Viper.GetString("alipay.k_server_domain") + "/third_party/notify"
	p.ReturnURL = variable.Viper.GetString("alipay.k_server_domain") + "/third_party/callback"
	p.Subject = "区块链养老平台订单:" + tradeNo
	p.OutTradeNo = tradeNo
	p.TotalAmount = strconv.Itoa(first.Price)
	p.ProductCode = "FAST_INSTANT_TRADE_PAY"
	p.Body = username
	query := database.AlipayRecord
	data := models.AlipayRecord{
		Username:  username,
		Code:      tradeNo,
		Status:    0,
		Choose:    int64(userInfo.ID),
		Timestamp: time.Now(),
		Account:   userInfo.Username,
	}
	err = query.WithContext(ctx).Create(&data)
	if err != nil {
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    400,
				"message": "存储记录失败: " + err.Error(),
			})
			return
		}
	}
	url, err := variable.AliClient.TradePagePay(p)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    400,
			"message": "生成支付链接失败: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "生成支付链接成功",
		"url":     url.String(),
	})
}

// AlipayCallback 支付宝成功回调函数
func AlipayCallback(ctx *gin.Context) {
	// 获取参数
	err := ctx.Request.ParseForm()
	if err != nil {
		logrus.Error("获取请求参数失败: ", err.Error())
		ctx.Redirect(http.StatusFound, consts.FrontAddress)
		return
	}
	var outTradeNo = ctx.Request.Form.Get("out_trade_no")
	var p = alipay.TradeQuery{}
	p.OutTradeNo = outTradeNo
	rsp, err := variable.AliClient.TradeQuery(p)
	if err != nil {
		logrus.Error("验证请求失败: ", err.Error())
		ctx.Redirect(http.StatusFound, consts.FrontAddress)
		return
	}
	if rsp.IsSuccess() == false {
		logrus.Error("验证订单信息失败: ", err.Error())
		ctx.Redirect(http.StatusFound, consts.FrontAddress)
		return
	}
	logrus.Info("支付成功: ", rsp)
	// 修改订单信息
	query := database.AlipayRecord
	_, err = query.WithContext(ctx).
		Where(query.Code.Eq(rsp.OutTradeNo)).
		Update(query.Status, 1)
	if rsp.IsSuccess() == false {
		logrus.Error("订单信息失败: ", err.Error())
		ctx.Redirect(http.StatusFound, consts.FrontAddress)
		return
	}

	ctx.Redirect(http.StatusFound, consts.FrontAddress)
}

// AlipayNotify 支付宝提示接口
func AlipayNotify(c *gin.Context) {
	c.Request.ParseForm()

	err := variable.AliClient.VerifySign(c.Request.Form)
	if err != nil {
		logrus.Error("异步通知验证签名发生错误", err)
		return
	}

	var outTradeNo = c.Request.Form.Get("out_trade_no")
	var p = alipay.TradeQuery{}
	p.OutTradeNo = outTradeNo
	rsp, err := variable.AliClient.TradeQuery(p)
	if err != nil {
		logrus.Error("异步通知验证订单 %s 信息发生错误: %s \n", outTradeNo, err.Error())
		return
	}
	if rsp.IsSuccess() == false {
		logrus.Error("异步通知验证订单 %s 信息发生错误: %s-%s \n", outTradeNo, rsp.Msg, rsp.SubMsg)
		return
	}

	logrus.Error("订单 %s 支付成功 \n", outTradeNo)
}

// WaitUserBuyCombo 购买商品实际接口
func WaitUserBuyCombo(ctx context.Context) {
	// 生成定时器 每五s读取一次数据库 也就每5s进行一次购买查询
	ticker := time.NewTicker(5 * time.Second)
	// 定时器 每次间隔一个月调用一次
	tickerTranfer := time.NewTicker(5 * time.Second)
	query := database.AlipayRecord
	userDao := database.UserInfo
	dao := database.TransferAccount
	moneyMap := make(map[string]int64)
	moneyMap["1000"] = 1
	moneyMap["3000"] = 2
	// 从数据库中获取所有未支付的订单
	for {
		select {
		// 每五s自动执行
		case <-ticker.C:
			// 执行数据库操作
			result, err := query.WithContext(ctx).
				Where(query.Status.Eq(1)).
				Find()
			if err != nil {
				logrus.Error("数据库操作失败: ", err.Error())
			}
			// 进行保险购买
			for _, v := range result {
				// 将数据序列化
				serialize, err := v.Serialize()
				if err != nil {
					logrus.Error("序列化数据失败: ", err.Error())
				}
				variable.BuyChannel <- serialize
			}
		// 购买保险channel读取数据	// 从channel中读取数据
		case data := <-variable.BuyChannel:
			// 将数据反序列化
			record, err := models.Deserialize(data)
			if err != nil {
				logrus.Error("反序列化数据失败: ", err.Error())
			}
			// 根据用户名获取用户信息
			userInfo, err := userDao.WithContext(ctx).Where(userDao.Username.Eq(record.Username)).First()
			if err != nil {
				logrus.Error("获取用户信息失败: ", err.Error())
			}
			// 从身份证上提取出生年月
			year, mouth, err := utils.ExtractBirthYearMonth(userInfo.IdCard)
			if err != nil {
				logrus.Error("获取出生年月失败: ", err.Error())
			}

			// 从当前时间中提取年份和月份
			currentYear := big.NewInt(int64(year))
			currentMonth := big.NewInt(int64(mouth))

			// 购买保险
			_, _, err = variable.ContractSession.PurchasePackage(record.Username, currentYear, currentMonth, big.NewInt(record.Choose))
			if err != nil {
				logrus.Error("购买保险失败: ", err.Error())
			}
			// 修改数据状态
			record.Status = 2
			_, err = query.WithContext(ctx).Where(query.ID.Eq(record.ID)).Updates(&record)
			if err != nil {
				logrus.Error("修改数据状态失败: ", err.Error())
			}
		case <-tickerTranfer.C:
			// 调用合约的里面的方法
			// 获取当前年和月
			year := time.Now().Year()
			month := time.Now().Month()
			_, _, err := variable.ContractSession.PayPensions(big.NewInt(int64(year)), big.NewInt(int64(month)))
			if err != nil {
				logrus.Error("调用合约事件错误：", err.Error())
			}
		case data := <-variable.TransChannel:
			var request models.TransferAccount
			// 反序列化
			err := json.Unmarshal([]byte(data), &request)
			if err != nil {
				logrus.Error("将数据反序列化失败: ", err.Error())
			}
			// 获取当前时间
			now := time.Now()

			// 获取当前月份的第一天
			firstDayOfMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
			// 获取下个月的第一天
			firstDayOfNextMonth := firstDayOfMonth.AddDate(0, 1, 0)
			// 获取当前月份的最后一天
			lastDayOfMonth := firstDayOfNextMonth.Add(-time.Nanosecond)

			// 首先查询数据这个月有没有发过养老金 发过则不发
			count, err := dao.WithContext(ctx).
				Where(dao.Timestamp.Between(firstDayOfMonth, lastDayOfMonth), dao.Username.Eq(request.Username)).
				Count()
			if err != nil {
				logrus.Error("查询数据失败: ", err.Error())
			}
			if count == 0 {
				// 获取用户信息
				user, err := query.WithContext(ctx).Where(query.Username.Eq(request.Username), query.Choose.Eq(moneyMap[request.Mount])).First()
				if err != nil {
					logrus.Error("查询数据失败: ", err.Error())
					continue
				}

				// 生成订单
				var tradeNo = fmt.Sprintf("%d", xid.Next())
				a := models.TransferAccount{
					Username:  request.Username,
					Account:   user.Account,
					Mount:     request.Mount,
					Timestamp: time.Now(),
					OutBizNo:  tradeNo,
				}
				err = dao.WithContext(ctx).Create(&a)
				if err != nil {
					logrus.Error("创建对象失败: ", err.Error())
				}
				// 调用支付宝转账接口
				result, err := variable.AliClient.FundTransToAccountTransfer(alipay.FundTransToAccountTransfer{
					PayeeType:    "ALIPAY_LOGONID",
					OutBizNo:     tradeNo,
					PayeeAccount: user.Account,
					Amount:       request.Mount,
				})
				if err != nil {
					log.Println(err)
				}

				logrus.Info("转账接口: ", result)
			}
		}
	}
}
