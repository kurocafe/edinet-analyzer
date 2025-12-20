package main

import (
	"fmt"
	"log"

	"github.com/kurocafe/edinet-analyzer/services"
)

func main() {
	// zipファイルを解析
	zipPath := "/app/data/edinet/S100TX1S.zip"

	fmt.Println("XBRL解析開始...")
	fmt.Println("対照ファイル: ", zipPath)
	fmt.Println()

	data, err := services.ParseXBRL(zipPath)
	if err != nil {
		log.Fatal("エラー: ", err)
	}

	fmt.Println("=== 解析結果 ===")
	fmt.Printf("会計年度: %d年\n", data.FiscalYear)
	fmt.Printf("売上高: %d円 (%.2f億円)\n", data.Revenue, float64(data.Revenue)/100000000)
	fmt.Printf("営業利益: %d円 (%.2f億円)\n", data.OperatingIncome, float64(data.OperatingIncome)/100000000)
	fmt.Printf("純利益: %d円 (%.2f億円)\n", data.NetIncome, float64(data.NetIncome)/100000000)
	fmt.Printf("配当金: %d円/株\n", data.Dividend)
}
