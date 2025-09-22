package routes

import (
	"star-auth-api/modules/common"
	"star-auth-api/modules/quotes"
	"star-auth-api/modules/star"
	"star-auth-api/modules/watermark"

	"github.com/gin-gonic/gin"
)

/**
 * 设置路由
 * @author zhaoshenchen
 */
func SetupRoutes(starHandler *star.StarHandler, quoteHandler *quotes.QuoteHandler, watermarkHandler *watermark.WatermarkHandler, unifiedHandler *common.UnifiedAPIHandler) *gin.Engine {
	// 创建Gin引擎
	r := gin.Default()

	// 设置CORS中间件
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// API路由组
	api := r.Group("/api")
	{
		// 健康检查
		api.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"status":  "ok",
				"message": "Service is running",
			})
		})

		// Star检测相关路由
		api.POST("/check-star", starHandler.CheckStar)
		api.POST("/verify-token", starHandler.VerifyToken)

		// 语录相关路由
		api.GET("/quote", quoteHandler.GetQuote)

		// 水印解析路由
		api.POST("/watermark", watermarkHandler.ParseWatermark)

		// 统一API入口（除Star检测外的其他功能）
		api.POST("/unified", unifiedHandler.HandleUnifiedAPI)
		
		// 接口调用统计
		api.GET("/stats", unifiedHandler.GetAPIStats)
	}

	return r
}
