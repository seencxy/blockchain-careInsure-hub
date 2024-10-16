package main

import (
	"backed/app/core/consts"
	"backed/app/core/variable"
	"backed/app/models"
	"backed/app/utils"
	_ "backed/bootstrap"
	"backed/internal/database"
	"context"
	"fmt"
	"math/big"
	"math/rand"
	"strconv"
	"time"
)

// 生成随机数据
func main() {

	id, _ := variable.ContractSession.NextPackageId()

	var economicCombo = models.Combo{
		Name:                "经济版养老保险套餐",
		Price:               50000,
		StartYear:           55,
		EndYear:             75,
		MonthFee:            800,
		HighMedicalCoverage: false,
		RefundPeriod:        2, // 假设这里的2表示2个月
		Description: []models.ComboDescription{
			{
				Description: "保障期限：覆盖至被保险人75岁。",
			},
			{
				Description: "医疗保险：包含基本的医疗保障。",
			},
			{
				Description: "养老金：每月支付800元基本生活费用。",
			},
			{
				Description: "取消: 支持2个月内取消(3个月内保险费用扣除，比如医疗，养老金等费用)",
			},
		},
	}
	economicCombo.ID = uint(id.Uint64())

	var fullCoverageCombo = models.Combo{
		Name:                "全能版养老保险套餐",
		Price:               350000,
		StartYear:           50,
		EndYear:             80,
		MonthFee:            2500,
		HighMedicalCoverage: true,
		RefundPeriod:        6, // 假设这里的6表示6个月
		Description: []models.ComboDescription{
			{
				Description: "保障期限：覆盖至被保险人80岁。",
			},
			{
				Description: "医疗保险：全面覆盖，包括专家会诊、国内外治疗等高级医疗服务。",
			},
			{
				Description: "养老金：每月支付2500元生活费用。",
			},
			{
				Description: "额外福利：包括长期护理、重大疾病保障、意外伤害保障等。",
			},
			{
				Description: "其他: 包含重疾、寿险、意外、医疗、教育金等多种保险。",
			},
			{
				Description: "取消: 支持半年内取消(半年内保险费用扣除，比如医疗，养老金等费用)",
			},
		},
	}
	fullCoverageCombo.ID = uint(id.Uint64() + 1)

	// 创建保险
	variable.ContractSession.AddPackage(
		big.NewInt(int64(economicCombo.Price)),
		big.NewInt(int64(economicCombo.MonthFee)),
		big.NewInt(int64(economicCombo.StartYear)),
		big.NewInt(int64(economicCombo.EndYear)),
		economicCombo.HighMedicalCoverage,
		big.NewInt(int64(economicCombo.RefundPeriod)))

	// 创建保险
	variable.ContractSession.AddPackage(
		big.NewInt(int64(fullCoverageCombo.Price)),
		big.NewInt(int64(fullCoverageCombo.MonthFee)),
		big.NewInt(int64(fullCoverageCombo.StartYear)),
		big.NewInt(int64(fullCoverageCombo.EndYear)),
		fullCoverageCombo.HighMedicalCoverage,
		big.NewInt(int64(fullCoverageCombo.RefundPeriod)))

	// 插入数据
	comboDao := database.Combo
	comboDao.WithContext(context.TODO()).Create(&economicCombo)
	comboDao.WithContext(context.TODO()).Create(&fullCoverageCombo)

	var users []*models.UserInfo
	var records []*models.AlipayRecord
	for i := 0; i < 100; i++ {
		var user models.UserInfo
		var record models.AlipayRecord

		user.Username = randomString(6)
		record.Username = user.Username
		user.Status = true
		record.Timestamp = time.Now()
		record.Code = randomString(10)
		record.Status = 2
		record.Choose = int64(randomOneOrTwo())
		record.Account = "pedggj8513@sandbox.com"

		user.Password = randomString(6)
		user.Name = randomString(6)
		user.Avatar = consts.DefaultAvatar
		user.IdCard = randomIdCard()
		last := user.IdCard[len(user.IdCard)-2 : len(user.IdCard)-1]
		s, _ := strconv.Atoi(last)

		if s%2 == 0 {
			user.Gender = "女"
		} else {
			user.Gender = "男"
		}

		//// 从身份证上提取出生年月
		//year, mouth, _ := utils.ExtractBirthYearMonth(user.IdCard)

		//// 从当前时间中提取年份和月份
		//currentYear := big.NewInt(int64(year))
		//currentMonth := big.NewInt(int64(mouth))

		//// 购买保险
		//_, _, _ = variable.ContractSession.PurchasePackage(record.Username, currentYear, currentMonth, big.NewInt(record.Choose))
		//// 获取年龄
		age, _ := utils.CalculateAge(user.IdCard)
		user.Age = int64(age)
		users = append(users, &user)
		records = append(records, &record)
	}

	// 插入数据
	userDao := database.UserInfo
	recordDao := database.AlipayRecord
	userDao.WithContext(context.TODO()).CreateInBatches(users, 100)
	recordDao.WithContext(context.TODO()).CreateInBatches(records, 100)
}

// 随机生成字符串
func randomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	rand.Seed(time.Now().UnixNano())
	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

// 随机生成出生年月日 (1900年到2024年之间)
func randomBirthDate() string {
	rand.Seed(time.Now().UnixNano())
	year := rand.Intn(2023-1900+1) + 1900 // 1900到2024年
	month := rand.Intn(12) + 1            // 1到12月
	day := rand.Intn(28) + 1              // 1到28日，简化处理，不考虑每月实际天数
	return fmt.Sprintf("%04d%02d%02d", year, month, day)
}

// 随机生成身份证号码（简化版，仅用于示例）
func randomIdCard() string {
	birthDate := randomBirthDate()
	seqCode := fmt.Sprintf("%03d", rand.Intn(1000)) // 随机生成三位顺序码
	// 简化处理，不生成真实校验码
	checkCode := fmt.Sprintf("%d", rand.Intn(10))

	// 假设地区码为 110000（北京市），仅作为示例
	return "110000" + birthDate + seqCode + checkCode
}

func randomOneOrTwo() int {
	rand.Seed(time.Now().UnixNano()) // 初始化随机数种子
	return rand.Intn(2) + 1          // 生成 0 或 1，然后加 1 得到 1 或 2
}
