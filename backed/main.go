package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	functions := []string{
		"buyPackageUsers(uint256)",
		"cancelPackage(string,uint256)",
		"getAllBuyersInfo()",
		"getBuyerInfo(string)",
		"getUnpurchasedPackage(string)",
		"owner()",
		"packages(uint256)",
		"payPensions(uint256,uint256)",
		"purchasePackage(string,uint256,uint256,uint256)",
	}

	for _, f := range functions {
		hash := crypto.Keccak256Hash([]byte(f))
		fmt.Printf("Function: %s, Method ID: %s\n", f, hash.Hex()[:10])
	}
}
