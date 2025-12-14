package services

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
