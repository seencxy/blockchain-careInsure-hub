package variable

import (
	RetirementInsurance "backed/internal/contract"
	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartwalle/alipay/v3"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

/* 存储全局变量 */

var (
	Viper           *viper.Viper    // 读取配置文件包
	Gorm            *gorm.DB        // 存储mysql连接
	FiscoBcos       *client.Client  // fisco bcos客户端
	AliClient       *alipay.Client  // 支付宝客户端
	ContractAddress *common.Address // 合约地址
	ContractSession *RetirementInsurance.RetirementInsuranceSession
	ContractAbi     string

	BuyChannel   chan string // 购买保险channel
	TransChannel chan string // 转账channel
)
