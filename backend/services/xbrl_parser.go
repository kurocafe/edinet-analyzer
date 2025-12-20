package services

import (
	"archive/zip"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
)

// XBRL解析結果
type XBRLData struct {
	FiscalYear      int   // 会計年度
	Revenue         int64 // 売上高
	OperatingIncome int64 // 営業利益
	NetIncome       int64 // 純利益
	Dividend        int   // 配当金（円/株）
}

// ZIPファイルからXBRLを解析
func ParseXBRL(zipPath string) (*XBRLData, error) {
	// ZIPファイルをひらく
	r, err := zip.OpenReader(zipPath)
	if err != nil {
		return nil, fmt.Errorf("ZIP読み込み失敗: %v", err)
	}
	defer r.Close()

	// PublicDocのXBRLファイルを探す
	var xbrlFile *zip.File
	for _, f := range r.File {
		// XBRLファイルをみつける
		if strings.Contains(f.Name, "PublicDoc") && strings.HasSuffix(f.Name, ".xbrl") {
			xbrlFile = f

			// debug1
			fmt.Println("XBRLファイル発見: ", f.Name)
			break
		}
	}

	if xbrlFile == nil {
		return nil, fmt.Errorf("XBRLファイルが見つかりません")
	}

	// XBRLファイルを開く
	rc, err := xbrlFile.Open()
	if err != nil {
		return nil, fmt.Errorf("XBRLファイル読み込み失敗: %v", err)
	}
	defer rc.Close()

	// 内容を読み込む
	content, err := io.ReadAll(rc)
	if err != nil {
		return nil, fmt.Errorf("XBRL内容読み込み失敗: %v", err)
	}

	// debug2
	fmt.Println("ファイルサイズ:", len(content), "バイト")

	// XMLをテキストとして解析
	data := &XBRLData{}

	// 会計年度
	fiscalYearMatch := regexp.MustCompile(`(\d{4})-\d{2}-\d{2}`).FindStringSubmatch(xbrlFile.Name)
	// 会計年度を取得
	if len(fiscalYearMatch) > 1 {
		data.FiscalYear, _ = strconv.Atoi(fiscalYearMatch[1])
	}

	contentStr := string(content)

	// debug3
	// NetSalesを含む行を探す
	fmt.Println("\n=== NetSales を検索 ===")
	netSalesPattern := regexp.MustCompile(`NetSalesSummaryOfBusinessResults[^>]*>([0-9]+)<`)
	if matches := netSalesPattern.FindStringSubmatch(contentStr); len(matches) > 0 {
		fmt.Println("見つかった:", matches[0])
	} else {
		fmt.Println("見つからなかった")
		// 別のパターンを試す
		fmt.Println("\nNetSalesを含む行を表示:")
		lines := strings.Split(contentStr, "\n")
		for i, line := range lines {
			if strings.Contains(line, "NetSales") && strings.Contains(line, "CurrentYear") {
				fmt.Printf("行%d: %s\n", i, line)
				if i+1 < len(lines) {
					fmt.Printf("行%d: %s\n", i+1, lines[i+1])
				}
			}
		}
	}

	// 売上高を取得
	revenueTags := []string{
		"NetSalesSummaryOfBusinessResults",
		"NetSales",
		"Sales",
		"OperatingRevenue",
		"OperatingRevenue1",
		"OperatingRevenue2",
		"RevenuesSummaryOfBusinessResults",
	}
	data.Revenue = extractValue(contentStr, revenueTags, "CurrentYearDuration")

	// 営業利益を取得
	operatingIncomeTags := []string{
		"OperatingIncome",
		"OperatingIncomeLoss",
		"OperatingIncomeLossSummaryOfBusinessResults",
		"OperatingProfit",
	}
	data.OperatingIncome = extractValue(contentStr, operatingIncomeTags, "CurrentYearDuration")

	// 純利益を取得
	netIncomeTags := []string{
		"ProfitLossAttributableToOwnersOfParent",
		"ProfitLossAttributableToOwnersOfParentSummaryOfBusinessResults",
		"NetIncome",
		"ProfitLoss",
		"NetIncomeLoss",
	}
	data.NetIncome = extractValue(contentStr, netIncomeTags, "CurrentYearDuration")

	// 配当金を取得
	dividendTags := []string{
		"DividendPaidPerShareSummaryOfBusinessResults",
		"DividendPaidPerShare",
		"CashDividendsPaidPerShare",
	}
	dividendValue := extractValue(contentStr, dividendTags, "CurrentYearDuration_NonConsolidatedMember")
	if dividendValue == 0 {
		// NonConsolidatedMember がない場合も試す
		dividendValue = extractValue(contentStr, dividendTags, "CurrentYearDuration")
	}
	data.Dividend = int(dividendValue)

	return data, nil
}

// タグから値を抽出
func extractValue(content string, tagNames []string, contextRef string) int64 {
	for _, tagName := range tagNames {
		// パターン1: contextRef が先
		pattern1 := fmt.Sprintf(`<%s[^>]*contextRef="%s"[^>]*>([0-9]+\.?[0-9]*)<`, tagName, contextRef)
		re1 := regexp.MustCompile(pattern1)
		if matches := re1.FindStringSubmatch(content); len(matches) > 1 {
			floatValue, err := strconv.ParseFloat(matches[1], 64)
			if err == nil && floatValue > 0 {
				return int64(floatValue)
			}
		}

		// パターン2: より柔軟なパターン（属性の順序を問わない）
		// contextRef="CurrentYearDuration" を含む行を探す
		lines := strings.Split(content, "\n")
		for _, line := range lines {
			if strings.Contains(line, tagName) && strings.Contains(line, fmt.Sprintf(`contextRef="%s"`, contextRef)) {
				// 数値を抽出
				numPattern := regexp.MustCompile(`>([0-9]+\.?[0-9]*)<`)
				if matches := numPattern.FindStringSubmatch(line); len(matches) > 1 {
					floatValue, err := strconv.ParseFloat(matches[1], 64)
					if err == nil {
						return int64(floatValue)
					}
				}
			}
		}
	}

	return 0
}
