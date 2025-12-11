package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	// 自分のパッケージ
	"github.com/kurocafe/edinet-analyzer/config"
	"github.com/kurocafe/edinet-analyzer/models"
)

func main() {
	// データベース接続
	config.ConnectDatabase()

	// マイグレーション（テーブル作成）
	// 存在すればスキップ
	config.DB.AutoMigrate(&models.Company{})

	// Ginのルーター作成
	// ロガーとリカバリー（パニック防止）ミドルウェア付き
	r := gin.Default()

	// CORS設定
	r.Use(cors.New(cors.Config{
		// ここから許可するオリジンを設定
		AllowOrigins:     []string{"http://localhost:3000"},

		// これらのメソッドを許可
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},

		// これらのヘッダーを許可
		AllowHeaders:     []string{"Origin", "Content-Type"},

		// クレデンシャル（Cookieなど）を許可
		AllowCredentials: true,
	}))

	// ルート
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello from EDINET Analyzer API",
			"status":  "ok",
		})
	})

	// 企業一覧API
	r.GET("/api/v1/companies", func(c *gin.Context) {
		var companies []models.Company
		config.DB.Find(&companies)
		c.JSON(http.StatusOK, companies)
	})

	// 企業追加API（テスト用）
	r.POST("/api/v1/companies", func(c *gin.Context) {
		var company models.Company
		if err := c.ShouldBindJSON(&company); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		config.DB.Create(&company)
		c.JSON(http.StatusCreated, company)
	})

	r.Run(":8080")
}