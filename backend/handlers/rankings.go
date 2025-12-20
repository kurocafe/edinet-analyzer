package handlers

import (
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
	"github.com/kurocafe/edinet-analyzer/config"
	"github.com/kurocafe/edinet-analyzer/models"
)

type RankingItem struct {
	Rank          int     `json:"rank"`
	CompanyID     uint    `json:"companyId"`
	CompanyName   string  `json:"companyName"`
	SecCode       string  `json:"secCode"`
	CurrentValue  int64   `json:"currentValue"`
	PreviousValue int64   `json:"previousValue"`
	GrowthRate    float64 `json:"growthRate"`
	CurrentYear   int     `json:"currentYear"`
	PreviousYear  int     `json:"previousYear"`
}

// 売上成長率ランキング
func GetRevenueGrowthRanking(ctx *gin.Context) {
	var companies []models.Company
	config.DB.Find(&companies)

	var rankings []RankingItem

	for _, company := range companies {
		// 2024年のデータ
		var current models.FinancialData
		result2024 := config.DB.Where("company_id = ? AND fiscal_year = 2024", company.ID).First(&current)

		// 2023年のデータ
		var previous models.FinancialData
		result2023 := config.DB.Where("company_id = ? AND fiscal_year = 2023", company.ID).First(&previous)

		// 両年のデータが存在する場合に計算
		if result2024.Error == nil && result2023.Error == nil {
			growthRate := float64(current.Revenue-previous.Revenue) / float64(previous.Revenue) * 100

			rankings = append(rankings, RankingItem{
				CompanyID:     company.ID,
				CompanyName:   company.Name,
				SecCode:       company.SecCode,
				CurrentValue:  current.Revenue,
				PreviousValue: previous.Revenue,
				GrowthRate:    growthRate,
				CurrentYear:   2024,
				PreviousYear:  2023,
			})
		}
	}

	// 成長率でソート（降順）
	// バブルソート(使わない)
	// for i := 0; i < len(rankings); i++ {
	// 	for j := i + 1; j < len(rankings); j++ {
	// 		if rankings[j].GrowthRate > rankings[i].GrowthRate {
	// 			rankings[i], rankings[j] = rankings[j], rankings[i]
	// 		}
	// 	}
	// }

	sort.Slice(rankings, func(i, j int) bool {
		return rankings[i].GrowthRate > rankings[j].GrowthRate
	})

	// ランク付け
	for i := range rankings {
		rankings[i].Rank = i + 1
	}

	ctx.JSON(http.StatusOK, rankings)
}

// 営業利益成長率ランキング
func GetProfitGrowthRanking(ctx *gin.Context) {
	var companies []models.Company
	config.DB.Find(&companies)

	var rankings []RankingItem

	for _, company := range companies {
		var current models.FinancialData
		result2024 := config.DB.Where("company_id = ? AND fiscal_year = 2024", company.ID).First(&current)

		var previous models.FinancialData
		result2023 := config.DB.Where("company_id = ? AND fiscal_year = 2023", company.ID).First(&previous)

		if result2024.Error == nil && result2023.Error == nil && previous.OperatingIncome > 0 {
			growthRate := float64(current.OperatingIncome-previous.OperatingIncome) / float64(previous.OperatingIncome) * 100

			rankings = append(rankings, RankingItem{
				CompanyID:     company.ID,
				CompanyName:   company.Name,
				SecCode:       company.SecCode,
				CurrentValue:  current.OperatingIncome,
				PreviousValue: previous.OperatingIncome,
				GrowthRate:    growthRate,
				CurrentYear:   2024,
				PreviousYear:  2023,
			})
		}
	}

	// 成長率でソート（降順）
	sort.Slice(rankings, func(i, j int) bool {
		return rankings[i].GrowthRate > rankings[j].GrowthRate
	})

	// ランク付け
	for i := range rankings {
		rankings[i].Rank = i + 1
	}

	ctx.JSON(http.StatusOK, rankings)
}
