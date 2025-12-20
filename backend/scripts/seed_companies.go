package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kurocafe/edinet-analyzer/config"
	"github.com/kurocafe/edinet-analyzer/models"
)

// このスクリプトは、data/companies.csv から企業データを読み込み、データベースに挿入します。
func main() {
	// 1. .env読み込み
	godotenv.Load("../.env")

	// 2. データベース接続
	config.ConnectDatabase()

	// 3. CSVファイルを開く
	file, err := os.Open("../data/companies.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 4. CSVを読み込む
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// 5. 1行ずつデータベースに挿入
	for i, record := range records {
		if i == 0 {
			continue // ヘッダー行をスキップ
		}

		// CSVの1行からCompany構造体を作る
		company := models.Company{
			Name:    record[0], // companies.csv の1列目
			SecCode: record[1], // companies.csv の2列目
		}

		// 既に存在するかチェック
		var existing models.Company
		result := config.DB.Where("sec_code = ?", company.SecCode).First(&existing)
		if result.Error == nil {
			fmt.Printf("スキップ: %s (既に存在)\n", company.Name)
			continue
		}

		// データベースに挿入
		config.DB.Create(&company)
		fmt.Printf("追加: %s (%s)\n", company.Name, company.SecCode)
	}

	fmt.Println("完了！")
}
