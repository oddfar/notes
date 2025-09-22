package main

import (
	"log"
	"star-auth-api/config"
	"star-auth-api/database"
	"star-auth-api/modules/common"
	"star-auth-api/modules/quotes"
	"star-auth-api/modules/star"
	"star-auth-api/modules/watermark"
	"star-auth-api/routes"
)

/**
 * 主函数
 * @author oddfar
 */
func main() {
	// 加载配置
	cfg := config.GetConfig()

	// 初始化数据库连接
	if err := database.InitDB(cfg); err != nil {
		log.Printf("数据库连接失败: %v", err)
		log.Println("将在没有数据库的情况下运行")
	}
	defer database.CloseDB()

	// 创建服务实例
	starService := star.NewStarService(cfg)
	quoteService := quotes.NewQuoteService("doc/sentences")
	watermarkService := watermark.NewWatermarkService()
	unifiedService := common.NewUnifiedAPIService(quoteService, watermarkService)

	// 创建处理器实例
	starHandler := star.NewStarHandler(starService)
	quoteHandler := quotes.NewQuoteHandler(quoteService)
	watermarkHandler := watermark.NewWatermarkHandler(watermarkService)
	unifiedHandler := common.NewUnifiedAPIHandler(unifiedService)

	// 设置路由
	r := routes.SetupRoutes(starHandler, quoteHandler, watermarkHandler, unifiedHandler)

	// 启动服务器
	log.Printf("服务器启动在端口 %s", cfg.Server.Port)
	if err := r.Run(":" + cfg.Server.Port); err != nil {
		log.Fatal("启动服务器失败:", err)
	}
}
