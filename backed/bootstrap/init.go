package bootstrap

import (
	"backed/app/core/consts"
	"backed/app/core/variable"
	"backed/app/models"
	RetirementInsurance "backed/internal/contract"
	"backed/internal/database"
	"fmt"
	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/conf"
	"github.com/ethereum/go-ethereum/common"
	"github.com/sirupsen/logrus"
	"github.com/smartwalle/alipay/v3"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// 启动函数 项目初始化

func init() {
	// 配置日志输出到文件
	logFile, err := os.OpenFile(consts.LogFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		// 设置日志输出到文件
		logrus.SetOutput(logFile)
	} else {
		logrus.Info("设置日志文件成功...")
	}

	// 设置日志级别 日志等级设置为最低
	logrus.SetLevel(logrus.TraceLevel)

	// 读取配置文件
	newViper := viper.New()
	// 设置配置文件存在目录
	newViper.SetConfigFile(consts.ConfigFilePath)
	// 读取配置文件内容
	err = newViper.ReadInConfig()
	if err != nil {
		log.Println("读取配置文件出错: ", err.Error())
	}
	// 初始化化viper
	variable.Viper = newViper

	// 连接mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		newViper.GetString("mysql.username"),
		newViper.GetString("mysql.password"),
		newViper.GetString("mysql.host"),
		newViper.GetString("mysql.port"),
		newViper.GetString("mysql.database"))

	variable.Gorm, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Info("连接mysql出错: ", err.Error())
		log.Println("连接mysql出错: ", err.Error())
		os.Exit(http.StatusBadRequest)
	}

	// 迁移数据库table
	variable.Gorm.AutoMigrate(&models.UserInfo{})
	variable.Gorm.AutoMigrate(&models.AlipayRecord{})
	variable.Gorm.AutoMigrate(&models.TransferAccount{})
	variable.Gorm.AutoMigrate(&models.ComboDescription{})
	variable.Gorm.AutoMigrate(&models.Combo{})

	// 设置默认操作数据库
	database.SetDefault(variable.Gorm)

	// 解析配置文件
	configs, err := conf.ParseConfigFile(consts.FiscoBcosFilePath)
	if err != nil {
		log.Fatal(err)
	}
	config := &configs[0]
	// 连接webase客户端
	client, err := client.Dial(config)
	if err != nil {
		log.Fatal("连接webase客户端失败: ", err.Error())
	}

	variable.FiscoBcos = client

	// 初始化alipay客户端
	aliClient, err := alipay.New(variable.Viper.GetString("alipay.app_id"), variable.Viper.GetString("alipay.k_private_key"), false)
	if err != nil {
		log.Fatal(err)
	}
	aliClient.LoadAliPayPublicKey(variable.Viper.GetString("alipay.alipay_public_key"))

	variable.AliClient = aliClient

	// 初始化合约对象
	ContractAddress := common.HexToAddress(variable.Viper.GetString("contract_address"))
	variable.ContractAddress = &ContractAddress

	// 创建合约实例
	instance, err := RetirementInsurance.NewRetirementInsurance(*variable.ContractAddress, variable.FiscoBcos)
	if err != nil {
		logrus.Error("创建合约实例失败: ", err.Error())
	}
	// 获取合约session
	session := &RetirementInsurance.RetirementInsuranceSession{Contract: instance, CallOpts: *variable.FiscoBcos.GetCallOpts(), TransactOpts: *variable.FiscoBcos.GetTransactOpts()}
	variable.ContractSession = session

	// 初始化chan
	variable.BuyChannel = make(chan string, 10000)
	variable.TransChannel = make(chan string, 10000)

	// 读取abi
	abi, err := ioutil.ReadFile(consts.ContractAbiPath)
	if err != nil {
		logrus.Error("读取abi文件失败: ", err.Error())
	}
	variable.ContractAbi = string(abi)
}
