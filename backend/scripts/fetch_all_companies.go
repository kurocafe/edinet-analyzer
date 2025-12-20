package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/kurocafe/edinet-analyzer/config"
	"github.com/kurocafe/edinet-analyzer/models"
	"github.com/kurocafe/edinet-analyzer/services"
)

func main() {
	// 環境変数読み込み
	godotenv.Load("../.env")

	// データベース接続
	config.ConnectDatabase()

	// 対照企業
	companies := []struct {
		ID      uint
		Name    string
		SecCode string
	}{
		{1, "ソニーグループ", "6758"},
		{2, "任天堂", "7974"},
		{3, "KDDI", "9433"},
		{4, "ソフトバンクグループ", "9984"},
		{5, "NTT", "9432"},
		{6, "楽天グループ", "4755"},
		{7, "サイバーエージェント", "4751"},
		{8, "DeNA", "2432"},
		{9, "メルカリ", "4385"},
		{10, "富士通", "6702"},
		{11, "NEC", "6701"},
		{12, "パナソニック", "6752"},
		{13, "キーエンス", "6861"},
		{14, "オムロン", "6645"},
		{15, "村田製作所", "6981"},
	}

	// 提出日候補（2024年6月）
	years := []struct {
		Year  int
		Dates []string
	}{
		{
			Year: 2023,
			Dates: []string{
				// 5月
				"2023-05-31", "2023-05-30", "2023-05-29", "2023-05-26", "2023-05-25",
				"2023-05-24", "2023-05-23", "2023-05-22", "2023-05-19", "2023-05-18",
				"2023-05-17", "2023-05-16", "2023-05-15", "2023-05-12", "2023-05-11",
				// 6月
				"2023-06-30", "2023-06-29", "2023-06-28", "2023-06-27", "2023-06-26",
				"2023-06-23", "2023-06-22", "2023-06-21", "2023-06-20", "2023-06-19",
				"2023-06-16", "2023-06-15", "2023-06-14", "2023-06-13", "2023-06-12",
				"2023-06-09", "2023-06-08", "2023-06-07", "2023-06-06", "2023-06-05",
				"2023-06-02", "2023-06-01",
				// 7月
				"2023-07-31", "2023-07-28", "2023-07-27", "2023-07-26", "2023-07-25",
				"2023-07-24", "2023-07-21", "2023-07-20", "2023-07-19", "2023-07-18",
				"2023-07-14", "2023-07-13", "2023-07-12", "2023-07-11", "2023-07-10",
				"2023-07-07", "2023-07-06", "2023-07-05", "2023-07-04", "2023-07-03",
			},
		},
		{
			Year: 2024,
			Dates: []string{
				"2024-05-31", "2024-05-30", "2024-05-29", "2024-05-28", "2024-05-27",
				"2024-06-28", "2024-06-27", "2024-06-26", "2024-06-25", "2024-06-24",
				"2024-06-21", "2024-06-20", "2024-06-19", "2024-06-18", "2024-06-17",
				"2024-06-14", "2024-06-13", "2024-06-12", "2024-06-11", "2024-06-10",
				"2024-07-31", "2024-07-30", "2024-07-29", "2024-07-26", "2024-07-25",
			},
		},
	}

	successCount := 0
	failCount := 0

	for _, yearInfo := range years {
		fmt.Printf("\n========== %d年度のデータ取得 ==========\n", yearInfo.Year)

		for _, company := range companies {
			fmt.Printf("\n=== %s (%s) %d年度 ===\n", company.Name, company.SecCode, yearInfo.Year)

			// 既に存在するかチェック
			var existing models.FinancialData
			result := config.DB.Where("company_id = ? AND fiscal_year = ?", company.ID, yearInfo.Year).First(&existing)
			if result.Error == nil {
				fmt.Printf("スキップ: 既に%d年度のデータがあります\n", yearInfo.Year)
				continue
			}

			var docID string
			var found bool
			var foundDate string

			// 複数の提出日を試す
			for _, date := range yearInfo.Dates {
				docs, err := services.GetDocuments(date, company.SecCode)
				if err != nil {
					continue
				}

				if len(docs) > 0 {
					docID = docs[0].DocID
					foundDate = date
					found = true
					break
				}
			}

			if !found {
				fmt.Printf("❌ 書類が見つかりませんでした\n")
				failCount++
				continue
			}

			fmt.Printf("📄 書類発見: %s (提出日: %s)\n", docID, foundDate)

			// 書類をダウンロード
			zipPath, err := services.DownloadDocument(docID)
			if err != nil {
				fmt.Printf("❌ ダウンロード失敗: %v\n", err)
				failCount++
				continue
			}
			fmt.Printf("⬇️ ダウンロード完了: %s\n", zipPath)

			// XBRL解析
			xbrlData, err := services.ParseXBRL(zipPath)
			if err != nil {
				fmt.Printf("❌ XBRL解析失敗: %v\n", err)
				failCount++
				continue
			}

			// DBに保存
			financialData := models.FinancialData{
				CompanyID:       company.ID,
				FiscalYear:      xbrlData.FiscalYear,
				Revenue:         xbrlData.Revenue,
				OperatingIncome: xbrlData.OperatingIncome,
				NetIncome:       xbrlData.NetIncome,
				Dividend:        xbrlData.Dividend,
			}
			config.DB.Create(&financialData)

			fmt.Printf("☑️ 保存完了: %d年度（売上: %.2f億円）\n",
				xbrlData.FiscalYear,
				float64(xbrlData.Revenue)/100000000)
			successCount++
		}
	}

	fmt.Println("\n========================================")
	fmt.Printf("完了: 成功 %d件 / 失敗 %d件\n", successCount, failCount)
	fmt.Println("========================================")
}
