package models

import "gorm.io/gorm"

type Combo struct {
	gorm.Model
	Name                string             `json:"name"`
	Description         []ComboDescription `json:"description"`
	Price               int                `json:"price"`
	StartYear           int                `json:"start_year"`
	EndYear             int                `json:"end_year"`
	MonthFee            int                `json:"month_fee"`
	HighMedicalCoverage bool               `json:"highMedicalCoverage"`
	RefundPeriod        int                `json:"refundPeriod"`
}

type ComboDescription struct {
	gorm.Model
	ComboID     uint   // 外键，指向Combo的ID
	Description string `json:"description"` // 实际的描述内容
}
