package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	// 自分のパッケージ
	"github.com/kurocafe/edinet-analyzer/config"
	"github.com/kurocafe/edinet-analyzer/models"
	"github.com/kurocafe/edinet-analyzer/services"
)

func main() {
	// データベース接続
	config.ConnectDatabase()

	// マイグレーション（テーブル作成）
	// 存在すればスキップ
	config.DB.AutoMigrate(&models.Company{})
	config.DB.AutoMigrate(&models.FinancialData{})

	// Ginのルーター作成
	// ロガーとリカバリー（パニック防止）ミドルウェア付き
	r := gin.Default()

	// CORS設定
	r.Use(cors.New(cors.Config{
		// ここから許可するオリジンを設定
		AllowOrigins: []string{"http://localhost:3000"},

		// これらのメソッドを許可
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},

		// これらのヘッダーを許可
		AllowHeaders: []string{"Origin", "Content-Type"},

		// クレデンシャル（Cookieなど）を許可
		AllowCredentials: true,
	}))

	// ルート
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello from EDINET Analyzer API",
			"status":  "ok",
		})
	})

	// 企業一覧API
	r.GET("/api/v1/companies", func(ctx *gin.Context) {
		var companies []models.Company
		config.DB.Find(&companies)
		ctx.JSON(http.StatusOK, companies)
	})

	// 企業追加API（テスト用）
	r.POST("/api/v1/companies", func(ctx *gin.Context) {
		var company models.Company
		if err := ctx.ShouldBindJSON(&company); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		config.DB.Create(&company)
		ctx.JSON(http.StatusCreated, company)
	})

	// 財務データ追加API
	r.POST("/api/v1/financial-data", func(ctx *gin.Context) {
		var data models.FinancialData
		if err := ctx.ShouldBindJSON(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		config.DB.Create(&data)
		ctx.JSON(http.StatusCreated, data)
	})

	// 財務データ一覧API
	r.GET("/api/v1/financial-data", func(ctx *gin.Context) {
		var financialData []models.FinancialData

		// クエリパラメータ取得
		fiscalYear := ctx.Query("fiscalYear")
		companyID := ctx.Query("companyId")

		query := config.DB

		if fiscalYear != "" {
			query = query.Where("fiscal_year = ?", fiscalYear)
		}
		if companyID != "" {
			query = query.Where("company_id = ?", companyID)
		}

		query.Find(&financialData)
		ctx.JSON(http.StatusOK, financialData)
	})

	// 特定企業の財務データAPI
	r.GET("/api/v1/companies/:id/financial-data", func(ctx *gin.Context) {
		companyID := ctx.Param("id")
		var financialData []models.FinancialData

		config.DB.Where("company_id = ?", companyID).Find(&financialData)
		ctx.JSON(http.StatusOK, financialData)
	})

	// EDINET書類一覧API
	r.GET("/api/v1/edinet/documents", func(ctx *gin.Context) {
		// クエリパラメータ取得
		date := ctx.Query("date")       // 形式: YYYY-MM-DD
		secCode := ctx.Query("secCode") // 証券コード4桁

		// dateは必須
		if date == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "date パラメータは必須です（例: 2024-06-28）",
			})
			return
		}

		// EDINET APIを呼び出し
		documents, err := services.GetDocuments(date, secCode)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, documents)
	})

	// EDINET書類ダウンロードAPI（ZIP形式）
	r.GET("/api/v1/edinet/documents/:docID/download", func(ctx *gin.Context) {
		docID := ctx.Param("docID")

		// EDINET APIを呼び出してZIPダウンロード
		filename, err := services.DownloadDocument(docID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message":  "ダウンロード完了",
			"filename": filename,
			"docID":    docID,
		})
	})

	// サーバー起動
	r.Run(":8080")
}
