package models

import (
	"gorm.io/gorm"
	"time"
)

type TransferAccount struct {
	gorm.Model
	Username  string
	Account   string
	Mount     string
	Timestamp time.Time
	OutBizNo  string
}
