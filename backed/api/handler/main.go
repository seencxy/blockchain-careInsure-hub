package handler

import (
	"backed/app/core/variable"
	"backed/app/models"
	"backed/app/utils"
	RetirementInsurance "backed/internal/contract"
	"backed/internal/database"
	"context"
	"encoding/json"
	"fmt"
	"github.com/FISCO-BCOS/go-sdk/core/types"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/smartwalle/alipay/v3"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// CheckUserInsurance 检查用户是否购买过保险
func CheckUserInsurance(ctx *gin.Context) {
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

	// 查询用户是否购买过保险
	unpurchasedPackage, err := variable.ContractSession.GetUnpurchasedPackages(username)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "检查是否购买过保险失败: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "检查是否购买过保险成功",
		"list":    unpurchasedPackage,
	})
}

// CancelInsurance 取消购买保险
func CancelInsurance(ctx *gin.Context) {
	var userInfo models.UserInfo
	choose := []string{"10", "50"}
	// 获取购买保险的套餐 套餐一: 基础套餐; 套餐二: 高级套餐
	err := ctx.ShouldBindJSON(&userInfo)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取请求参数失败:" + err.Error(),
		})
		return
	}
	if userInfo.ID != 1 && userInfo.ID != 2 {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "保险套餐出错:" + string(userInfo.ID),
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

	// 查询用户是否购买过保险
	unpurchasedPackage, err := variable.ContractSession.GetUnpurchasedPackages(username)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "检查是否购买过保险失败: " + err.Error(),
		})
		return
	}

	// 查看是否包含此ID
	for _, v := range unpurchasedPackage {
		if v.Int64() == int64(userInfo.ID) {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    500,
				"message": "用户并未购买该套餐保险",
			})
			return
		}
	}

	// 取消套餐
	tx, rep, err := variable.ContractSession.CancelPackage(username, big.NewInt(int64(userInfo.ID)))
	if err != nil || rep.Status != 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "取消套餐失败: " + rep.GetErrorMessage(),
		})
		return
	}

	// 首先查询订单信息
	dao := database.AlipayRecord
	first, err := dao.Where(dao.Username.Eq(username), dao.Choose.Eq(int64(userInfo.ID))).First()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "查询订单失败: " + err.Error(),
		})
		return
	}
	// 修改订单信息
	_, err = dao.Where(dao.Username.Eq(username), dao.Choose.Eq(int64(userInfo.ID))).Delete()
	if err != nil {
		logrus.Error("修改订单信息失败: ", err.Error())
	}

	result, err := variable.AliClient.TradeRefund(alipay.TradeRefund{
		OutTradeNo:   first.Code,
		RefundAmount: choose[userInfo.ID-1],
		RefundReason: "用户取消购买保险",
	})

	variable.AliClient.FundTransToAccountTransfer(alipay.FundTransToAccountTransfer{})

	logrus.Error("退款信息: ", result)

	if err != nil {
		logrus.Error(first.Code+" 退款失败: ", err.Error())
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "取消购买保险成功",
		"hash":    tx.Hash().Hex(),
	})
}

// TransferMoney 发放养老金
func TransferMoney(ctx context.Context) {
	// 获取最新区块 高度
	blockNumber, err := variable.FiscoBcos.GetBlockNumber(ctx)
	if err != nil {
		logrus.Error("获取最新区块高度失败: ", err.Error())
	}
	// 获取当前区块高度
	currentBlockNumber := uint64(blockNumber)
	// 根据配置文件构建eventDemoSession
	_, err = variable.ContractSession.WatchAllPensionPaid(&currentBlockNumber, func(status int, logs []types.Log) {
		fmt.Println("Status Code:", status)
		for k, v := range logs {
			log, errs := variable.ContractSession.ParsePensionPaid(v) // 解析数据
			if errs != nil {
				logrus.Println(errs)
			}
			request := models.TransferAccount{
				Username: log.Username,
				Account:  log.Amount.String(),
			}
			// 将数据序列化
			marshal, errs := json.Marshal(request)
			if errs != nil {
				logrus.Error("序列化失败: ", err.Error())
			}

			variable.TransChannel <- string(marshal)
			logrus.Println("Log id:", k, "Log Content:", log)
		}
	})
	if err != nil {
		logrus.Error("监听合约事件失败: ", err.Error())
	}
}

// GetAllOrderUser 获取所有购买人的订单
func GetAllOrderUser(ctx *gin.Context) {
	var userInfo models.UserInfo
	// 获取所有信息

	id := ctx.Query("id")
	username := ctx.Query("username")
	userInfo.Username = username
	atoi, err2 := strconv.Atoi(id)
	if err2 != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取请求参数失败: " + err2.Error(),
		})
		return
	}
	userInfo.ID = uint(int64(atoi))
	// 获取所有购买人的信息
	allBuyersInfo, err := variable.ContractSession.GetAllBuyersInfo()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取页数失败:" + err.Error(),
		})
		return
	}

	type s struct {
		Id        int64     `json:"id"`
		Username  string    `json:"username"`
		Gender    string    `json:"gender"`
		Age       int64     `json:"age"`
		Kind      string    `json:"kind"`
		Timestamp time.Time `json:"timestamp"`
	}

	// 初始化数据切片
	var data []s

	if len(allBuyersInfo) == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "暂时没有数据",
			"data":    "",
		})
		return
	}

	var ppp []RetirementInsurance.RetirementInsuranceUserInfo
	// 在这里获取用户信息
	if len(userInfo.Username) != 0 {
		for _, v := range allBuyersInfo {
			contains := strings.Contains(v.Username, userInfo.Username)
			if contains {
				ppp = append(ppp, v)
			}
		}
	} else {
		ppp = allBuyersInfo
	}

	userDao := database.UserInfo
	recordDao := database.AlipayRecord

	// 分页逻辑
	startIndex, endIndex := utils.CalculatePaginationIndexes(len(ppp), int(userInfo.ID), 10)

	// 遍历用户信息
	for i := startIndex; i >= endIndex; i-- {
		buyer := ppp[i]

		// 获取数据库中的用户信息
		user, err := userDao.WithContext(ctx).
			Where(userDao.Username.Eq(ppp[i].Username)).
			First()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"code":    500,
				"message": "获取用户信息失败: " + err.Error(),
			})
			return
		}

		// 获取订单记录
		record, err := recordDao.WithContext(ctx).
			Where(recordDao.Username.Eq(ppp[i].Username), recordDao.Choose.Eq(ppp[i].PackageId.Int64())).
			First()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"code":    500,
				"message": "获取订单记录失败: " + err.Error(),
			})
			return
		}

		// 添加到用户详细信息切片
		data = append(data, s{
			Id:        int64(user.ID),
			Username:  user.Username,
			Gender:    user.Gender,
			Age:       user.Age,
			Kind:      buyer.PackageId.String(),
			Timestamp: record.Timestamp,
		})
	}

	// 返回结果
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取所有购买人信息成功",
		"data":    data,
	})
}

func updateAgeGroup(age int, group *AgeGroups) {
	switch {
	case age >= 0 && age <= 20:
		group.Age0to20++
	case age > 20 && age <= 30:
		group.Age20to30++
	case age > 30 && age <= 40:
		group.Age30to40++
	case age > 40 && age <= 50:
		group.Age40to50++
	case age > 50 && age <= 60:
		group.Age50to60++
	case age > 60 && age <= 70:
		group.Age60to70++
	case age > 70 && age <= 80:
		group.Age70to80++
	default:
		group.Age80Plus++
	}
}

// AgeGroups 用来表示不同的年龄段
type AgeGroups struct {
	Age0to20  int `json:"age0To20"`
	Age20to30 int `json:"age20To30"`
	Age30to40 int `json:"age30To40"`
	Age40to50 int `json:"age40To50"`
	Age50to60 int `json:"age50To60"`
	Age60to70 int `json:"age60To70"`
	Age70to80 int `json:"age70To80"`
	Age80Plus int `json:"age80Plus"`
}

func classifyAndCountAgeGroupsDynamically(records []*models.AlipayRecord, users []*models.UserInfo) map[string]AgeGroups {
	result := make(map[string]AgeGroups)
	userMap := make(map[string]*models.UserInfo)

	// 将用户信息放入map中，以便快速查找
	for _, user := range users {
		userMap[user.Username] = user
	}

	for _, record := range records {
		if user, exists := userMap[record.Username]; exists {
			chooseStr := strconv.Itoa(int(record.Choose)) // 正确地将int转为string

			// 直接更新result中的AgeGroups实例
			group := result[chooseStr]
			updateAgeGroup(int(user.Age), &group)
			result[chooseStr] = group // 因为group是一个副本，所以这一步是必须的
		}
	}

	return result
}

// StatisticsInfo 获取统计信息
func StatisticsInfo(ctx *gin.Context) {

	dao := database.AlipayRecord
	userDao := database.UserInfo
	// 获取所有订单信息
	find, err := dao.WithContext(ctx).Where(dao.Status.Eq(2)).Find()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取订单信息失败: " + err.Error(),
		})
		return
	}

	var userIds []string
	for _, v := range find {
		userIds = append(userIds, v.Username)
	}

	// 获取所有用户信息
	find2, err := userDao.WithContext(ctx).Where(userDao.Username.In(userIds...)).Find()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取用户信息失败: " + err.Error(),
		})
		return
	}

	// 动态分类并计算年龄分组统计信息
	result := classifyAndCountAgeGroupsDynamically(find, find2)

	comboDao := database.Combo
	// 使用临时map存储更新后的结果，避免在遍历过程中直接修改原map
	updatedResult := make(map[string]AgeGroups)
	for k, v := range result {
		atoi, err := strconv.Atoi(k)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    500,
				"message": "套餐ID转换失败: " + err.Error(),
			})
			return
		}
		combo, err := comboDao.Where(comboDao.ID.Eq(uint(atoi))).First()
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    500,
				"message": "查找套餐名称失败: " + err.Error(),
			})
			return
		}
		updatedResult[combo.Name] = v
	}
	// 使用更新后的结果
	result = updatedResult

	// 返回结果
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取统计信息成功",
		"data":    result,
	})

}

// GetTransactionInfo 获取交易信息
func GetTransactionInfo(ctx *gin.Context) {
	methods := map[string]string{
		"0x4367edd9": "buyPackageUsers(uint256)",
		"0x32ce40b1": "cancelPackage(string,uint256)",
		"0x3b3b57de": "getAllBuyersInfo()",
		"0x36945753": "getBuyerInfo(string)",
		"0xcc2590ea": "getUnpurchasedPackage(string)",
		"0x8da5cb5b": "owner()",
		"0xc216212a": "packages(uint256)",
		"0xb853f19a": "payPensions(uint256,uint256)",
		"0x9c96de78": "purchasePackage(string,uint256,uint256,uint256)",
	}
	id := ctx.Query("id")
	page, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取请求参数失败: " + err.Error(),
		})
		return
	}
	// 获取最新区块
	blockNumber, err := variable.FiscoBcos.GetBlockNumber(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取最新区块失败: " + err.Error(),
		})
		return
	}
	type result struct {
		Hash        string `json:"hash"`
		Method      string `json:"method"`
		BlockNumber string `json:"blockNumber"`
		Timestamp   string `json:"timestamp"`
	}

	var results []result
	// 分页逻辑
	startIndex, endIndex := utils.CalculatePaginationIndexes(int(blockNumber), int(page), 10)
	for i := startIndex; i >= endIndex; i-- {
		var res result
		block, err := variable.FiscoBcos.GetBlockByNumber(ctx, int64(i), true)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    500,
				"message": "获取区块信息失败: " + err.Error(),
			})
			return
		}
		res.Timestamp = block.Timestamp
		for _, tx := range block.Transactions {
			da := tx.(map[string]interface{})
			// 获取input前个字符
			method := da["input"].(string)[:10]
			res.Method = methods[method]
			res.Hash = da["hash"].(string)
			res.BlockNumber = da["blockNumber"].(string)
		}
		results = append(results, res)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取区块信息成功",
		"data":    results,
	})
}

// GetNodeInfo 获取节点信息
func GetNodeInfo(ctx *gin.Context) {
	// 定义响应参数
	var res models.GetNodeInfoResponse
	// 获取区块数量
	number, err := variable.FiscoBcos.GetBlockNumber(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取区块数量失败: " + err.Error(),
		})
		return
	}
	res.BlockNumber = strconv.Itoa(int(number))

	// 获取节点数量
	infoJson, err := variable.FiscoBcos.GetNodeIDList(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取节点数量失败: " + err.Error(),
		})
		return
	}

	var list []string
	// 解析json
	err = json.Unmarshal(infoJson, &list)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "解析节点信息失败: " + err.Error(),
		})
		return
	}

	res.NodeNumber = len(list)

	// 获取交易数量
	transactions, err := variable.FiscoBcos.GetTotalTransactionCount(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取交易数量失败: " + err.Error(),
		})
		return
	}

	res.TransactionNumber = transactions.BlockNumber

	// 获取交易失败的数量
	transactionsFail, err := variable.FiscoBcos.GetPendingTransactions(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取交易失败数量失败: " + err.Error(),
		})
		return
	}

	res.PendingTransactionNumber = len(*transactionsFail)

	res.NodeInfoList = list

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取节点信息成功",
		"data":    res,
	})
}

// GetAllUserInfo 获取所有用户信息
func GetAllUserInfo(ctx *gin.Context) {
	userDao := database.UserInfo
	// 获取所有用户信息
	users, err := userDao.WithContext(ctx).Find()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取用户信息失败: " + err.Error(),
		})
		return
	}

	// 将管理员删除
	var result []*models.UserInfo
	for _, element := range users {
		if element.Username != "admin123" {
			result = append(result, element)
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取用户信息成功",
		"data":    result,
	})
}

// UpdateUserStatus 修改用户状态
func UpdateUserStatus(ctx *gin.Context) {
	var userInfo models.UserInfo
	err := ctx.ShouldBindJSON(&userInfo)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取请求参数失败: " + err.Error(),
		})
		return
	}

	userDao := database.UserInfo
	// 修改用户状态
	_, err = userDao.Where(userDao.Username.Eq(userInfo.Username)).Update(userDao.Status, userInfo.Status)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "修改用户状态失败: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "修改用户状态成功",
	})
}

// GetPayRecord 获取付款记录
func GetPayRecord(ctx *gin.Context) {
	// 获取所有付款记录
	dao := database.AlipayRecord
	records, err := dao.WithContext(ctx).Find()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取付款记录失败: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取付款记录成功",
		"data":    records,
	})
}

// CreateCombo 创建套餐
func CreateCombo(ctx *gin.Context) {
	// 绑定请求参数
	var combo models.Combo
	err := ctx.ShouldBindJSON(&combo)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取请求参数失败: " + err.Error(),
		})
		return
	}

	_, receipt, err := variable.ContractSession.AddPackage(
		big.NewInt(int64(combo.Price)),
		big.NewInt(int64(combo.MonthFee)),
		big.NewInt(int64(combo.StartYear)),
		big.NewInt(int64(combo.EndYear)),
		combo.HighMedicalCoverage,
		big.NewInt(int64(combo.RefundPeriod)))
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "创建套餐失败: " + receipt.GetErrorMessage(),
		})
		return
	}

	// 获取下一个此pageID
	pageID, err := variable.ContractSession.NextPackageId()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取套餐ID失败: " + err.Error(),
		})
		return
	}
	combo.ID = uint(pageID.Int64() - 1)

	// 创建套餐
	dao := database.Combo
	err = dao.WithContext(ctx).Create(&combo)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "创建套餐失败: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":            200,
		"message":         "创建套餐成功",
		"transactionHash": receipt.TransactionHash,
	})
}

// GetAllCombo 获取所有套餐
func GetAllCombo(ctx *gin.Context) {
	packages, err := variable.ContractSession.GetAllPackages()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取全部套餐失败: " + err.Error(),
		})
		return
	}

	comboDao := database.Combo

	var result []models.Combo
	for _, element := range packages {
		first, err := comboDao.Select().Where(comboDao.ID.Eq(uint(element.Id.Int64()))).First()
		if err != nil {
			continue
		}
		// 创建一个长度
		first.Description = make([]models.ComboDescription, 1)
		if element.IsActive {
			first.Description[0].ComboID = 1
		} else {
			first.Description[0].ComboID = 0
		}
		result = append(result, *first)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取套餐成功",
		"data":    result,
	})
}

// UpdateComboStatus 更新套餐状态
func UpdateComboStatus(ctx *gin.Context) {
	var combo models.Combo
	err := ctx.ShouldBindJSON(&combo)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取请求参数失败: " + err.Error(),
		})
		return
	}

	_, receipt, err := variable.ContractSession.UpdatePackageStatus(big.NewInt(int64(combo.ID)))
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "更新套餐状态失败: " + receipt.GetErrorMessage(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":            200,
		"message":         "更新套餐状态成功",
		"transactionHash": receipt.TransactionHash,
	})
}

// GetAllComboUser  获取所有套餐
func GetAllComboUser(ctx *gin.Context) {
	// 获取所有用户
	users, err := variable.ContractSession.GetAllPackages()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取所有用户失败: " + err.Error(),
		})
		return
	}

	comboDao := database.Combo
	descDao := database.ComboDescription
	var packages []*models.Combo
	for _, element := range users {
		if !element.IsActive {
			return
		}

		first, err := comboDao.Select().Where(comboDao.ID.Eq(uint(element.Id.Int64()))).First()
		if err != nil {
			continue
		}
		find, err := descDao.Where(descDao.ComboID.Eq(first.ID)).Find()
		if err != nil {
			continue
		}

		for _, v := range find {
			first.Description = append(first.Description, *v)
		}
		packages = append(packages, first)
		// 查询简介
		// 查询数据库中的套餐信息
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取所有产品成功",
		"data":    packages,
	})
}

// GetComboDetail 获取套餐详情
func GetComboDetail(ctx *gin.Context) {
	id := ctx.Query("id")
	atoi, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取请求参数失败: " + err.Error(),
		})
		return
	}

	comboDao := database.Combo
	descDao := database.ComboDescription
	first, err := comboDao.Select().Where(comboDao.ID.Eq(uint(atoi))).First()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取套餐失败: " + err.Error(),
		})
		return
	}

	find, err := descDao.Where(descDao.ComboID.Eq(first.ID)).Find()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取套餐详情失败: " + err.Error(),
		})
		return
	}

	for _, v := range find {
		first.Description = append(first.Description, *v)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取套餐详情成功",
		"data":    first,
	})
}
