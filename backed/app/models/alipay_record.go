package models

import (
	"encoding/json"
	"gorm.io/gorm"
	"time"
)

type AlipayRecord struct {
	gorm.Model
	Username  string
	Timestamp time.Time
	Code      string
	Status    int64
	Choose    int64
	Account   string
}

// Serialize 序列化
func (a *AlipayRecord) Serialize() (string, error) {
	// 对数据序序列化
	marshal, err := json.Marshal(a)
	if err != nil {
		return "", err
	}
	return string(marshal), nil
}

// Deserialize 反序列化
func Deserialize(s string) (*AlipayRecord, error) {
	var a AlipayRecord
	err := json.Unmarshal([]byte(s), &a)
	if err != nil {
		return nil, err
	}

	return &a, nil
}
