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
	OparatingIncome int64 // 営業利益
	NetIncome       int64 // 純利益
	Dividend        int   // 配当金（円/株）
}

// ZIPファイルからXBRLを解析
func ParaseXBRL(zipPath string) (*XBRLData, error) {
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

	// XMLをテキストとして解析
	data := &XBRLData{}

	// 会計年度
	fiscalYearMatch := regexp.MustCompile(`(\d{4})-\d{2}-\d{2}`).FindStringSubmatch(xbrlFile.Name)
	// 会計年度を取得
	if len(fiscalYearMatch) > 1 {
		data.FiscalYear, _ = strconv.Atoi(fiscalYearMatch[1])
	}

	contentStr := string(content)

	// 売上高を取得
	data.Revenue = extractValue(contentStr, "NetSalesSummaryOfBusinessResults", "CurrentYearDuration")

	// 営業利益を取得
	data.OparatingIncome = extractValue(contentStr, "OperatingIncomeLossSummaryOfBusinessResults", "CurrentYearDuration")

	// 純利益を取得
	data.NetIncome = extractValue(contentStr, "ProfitLossAttributableToOwnersOfParentSummaryOfBusinessResults", "CurrentYearDuration")

	// 配当金を取得
	dividendValue := extractValue(contentStr, "DividendPaidPerShareSummaryOfBusinessResults", "CurrentYearDuration")
	data.Dividend = int(dividendValue)

	return data, nil
}

// タグから値を抽出
func extractValue(content string, tagName string, contextRef string) int64 {
	// 正規表現でタグを探す
	pattern := fmt.Sprintf(`<%s[^>]*contextRef="%s"[^>]*>([0-9]+)<`, tagName, contextRef)
	re := regexp.MustCompile(pattern)
	matches := re.FindStringSubmatch(content)

	if len(matches) > 1 {
		value, err := strconv.ParseInt(matches[1], 10, 64)
		if err == nil {
			return value
		}
	}

	return 0
}
