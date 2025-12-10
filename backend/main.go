package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kurocafe/edinet-analyzer/config"
	"github.com/kurocafe/edinet-analyzer/models"
)

func main() {
	// データベース接続
	config.ConnectDatabase()

	// マイグレーション（テーブル作成）
	config.DB.AutoMigrate(&models.Company{})

	// Ginのルーター作成
	r := gin.Default()

	// CORS設定
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
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