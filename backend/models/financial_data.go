package models

import "gorm.io/gorm"

type FinancialData struct {
	gorm.Model
	CompanyID       uint    `gorm:"not null;index" json:"companyId"`               // 企業ID（外部キー）
	Company         Company `gorm:"foreignKey:CompanyID" json:"company,omitempty"` // 企業情報（リレーション）
	FiscalYear      int     `gorm:"not null;index" json:"fiscalYear"`              // 会計年度（例: 2023）
	Revenue         int64   `json:"revenue"`                                       // 売上高（円）
	OperatingIncome int64   `json:"operatingIncome"`                               // 営業利益（円）
	NetIncome       int64   `json:"netIncome"`                                     // 純利益（円）
	Dividend        int     `json:"dividend"`                                      // 配当金（円/株）
}
