package main

import (
	"backed/app/core/consts"
	RetirementInsurance "backed/internal/contract"
	"fmt"
	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/conf"
	"log"
)

// 部署合约写入参数
func main() {
	configs, err := conf.ParseConfigFile(consts.FiscoBcosFilePath)
	if err != nil {
		log.Fatal(err)
	}
	config := &configs[0]

	client, err := client.Dial(config)
	if err != nil {
		log.Fatal(err)
	}
	address, tx, instance, err := RetirementInsurance.DeployRetirementInsurance(client.GetTransactOpts(), client)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("contract address: ", address.Hex()) // the address should be saved
	fmt.Println("transaction hash: ", tx.Hash())
	_ = instance

}
