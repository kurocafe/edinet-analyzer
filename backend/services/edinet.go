package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// EDINET API レスポンス構造体
type EDINETDocument struct {
	DocID          string `json:"docID"`
	EDINETCode     string `json:"edinetCode"`
	SecCode        string `json:"secCode"`
	FilerName      string `json:"filerName"`
	DocTypeCode    string `json:"docTypeCode"`
	PeriodStart    string `json:"periodStart"`
	PeriodEnd      string `json:"periodEnd"`
	SubmitDateTime string `json:"submitDateTime"`
	XBRLFlag       string `json:"xbrlFlag"`
}

type EDINETResponse struct {
	Results []EDINETDocument `json:"results"`
}

// EDINET書類一覧API
// date: YYYY-MM-DD形式
// secCode: 証券コード4桁 → 5桁に変換して検索
func GetDocuments(date string, secCode string) ([]EDINETDocument, error) {
	// APIキー取得
	apiKey := os.Getenv("EDINET_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("EDINET_API_KEY が設定されていません")
	}

	// EDINET APIのURL
	url := fmt.Sprintf(
		"https://disclosure.edinet-fsa.go.jp/api/v2/documents.json?date=%s&type=2&Subscription-Key=%s",
		date,
		apiKey,
	)

	// HTTPリクエスト
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("EDINET APIリクエスト失敗: %v", err)
	}
	defer resp.Body.Close()

	// レスポンスボディの読み込み
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("レスポンス読み込み失敗: %v", err)
	}

	// JSONデコード
	var result EDINETResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("JSON解析失敗: %v", err)
	}

	// 証券コードでフィルタリング
	if secCode != "" {
		// 4桁 → 5桁に変換
		secCode5 := secCode + "0"

		filtered := []EDINETDocument{}
		for _, doc := range result.Results {
			// 有価証券報告書(docTypeCode: 120)かつコード一致
			if doc.DocTypeCode == "120" && doc.SecCode == secCode5 {
				filtered = append(filtered, doc)
			}
		}

		return filtered, nil
	}

	// 証券コード指定なしの場合は有価証券報告書をすべて返す
	filtered := []EDINETDocument{}
	for _, doc := range result.Results {
		if doc.DocTypeCode == "120" {
			filtered = append(filtered, doc)
		}
	}

	return filtered, nil
}

// EDINET書類取得API
// docID: 書類ID（例：S100TX1S）
// 戻り値: 保存したファイルパス
func DownloadDocument(docID string) (string, error) {
	apiKey := os.Getenv("EDINET_API_KEY")
	if apiKey == "" {
		return "", fmt.Errorf("EDINET_API_KEY が設定されていません")
	}

	// EDINET 書類取得APIのURL
	// type=1: 提出本文書及び監査報告書（ZIP形式）
	url := fmt.Sprintf(
		"https://disclosure.edinet-fsa.go.jp/api/v1/documents/%s?type=1&Subscription-Key=%s",
		docID,
		apiKey,
	)

	// HTTPリクエスト
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("EDINET APIリクエスト失敗: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("EDINET API エラー: ステータスコード %d", resp.StatusCode)
	}

	// 保存先ディレクトリ作成
	downloadDir := "/app/data/edinet"
	os.MkdirAll(downloadDir, 0755)

	// ファイル名
	filename := fmt.Sprintf("%s/%s.zip", downloadDir, docID)

	// ファイル作成
	file, err := os.Create(filename)
	if err != nil {
		return "", fmt.Errorf("ファイル作成失敗: %v", err)
	}
	defer file.Close()

	// レスポンスボディをファイルに書き込み
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return "", fmt.Errorf("ファイル書き込み失敗: %v", err)
	}

	return filename, nil
}
