package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// Ginのルーターを作成
	r := gin.Default()

	r.GET(
		"/", 									// パス
		func(c *gin.Context) 	// ハンドラー関数
			c.JSON(							// JSONレスポンスを返す
				http.StatusOK,		// ステータスコード200
				 gin.H{						// gin.Hはmap[string]interface{}のエイリアス
					"message": "Hello, World!",
				}
			)
		}
	)

	// サーバーをポート8080で起動
	r.Run(":8080")
}