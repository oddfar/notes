package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

// Repository 仓库信息
type Repository struct {
	Platform string `json:"platform"`
	URL      string `json:"url"`
}

// CheckStarRequest 检查star请求
type CheckStarRequest struct {
	Repositories []Repository `json:"repositories"`
}

// CheckStarResponse 检查star响应
type CheckStarResponse struct {
	Success bool   `json:"success"`
	Token   string `json:"token,omitempty"`
	Message string `json:"message"`
}

// VerifyTokenRequest 验证token请求
type VerifyTokenRequest struct {
	Token string `json:"token"`
}

// VerifyTokenResponse 验证token响应
type VerifyTokenResponse struct {
	Valid bool `json:"valid"`
}

// 简单的内存存储，生产环境应使用数据库
var validTokens = make(map[string]time.Time)

/**
 * 检查用户是否已star指定仓库
 * @author zhaoshenchen
 */
func checkStarHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req CheckStarRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// 这里应该调用GitHub/Gitee API检查star状态
	// 目前使用mock数据进行演示
	isStarred := checkRepositoriesStarStatus(req.Repositories)

	var response CheckStarResponse
	if isStarred {
		// 生成token
		token := generateToken()
		validTokens[token] = time.Now().Add(24 * time.Hour) // token有效期24小时

		response = CheckStarResponse{
			Success: true,
			Token:   token,
			Message: "验证成功！感谢您的支持！",
		}
	} else {
		response = CheckStarResponse{
			Success: false,
			Message: "请先Star任意一个项目后再进行验证",
		}
	}

	json.NewEncoder(w).Encode(response)
}

/**
 * 验证token有效性
 * @author zhaoshenchen
 */
func verifyTokenHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req VerifyTokenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// 检查token是否存在且未过期
	expireTime, exists := validTokens[req.Token]
	valid := exists && time.Now().Before(expireTime)

	// 清理过期token
	if exists && !valid {
		delete(validTokens, req.Token)
	}

	response := VerifyTokenResponse{
		Valid: valid,
	}

	json.NewEncoder(w).Encode(response)
}

/**
 * 检查仓库star状态（Mock实现）
 * 生产环境中应该调用实际的GitHub/Gitee API
 * @author zhaoshenchen
 */
func checkRepositoriesStarStatus(repositories []Repository) bool {
	// Mock实现：70%概率返回true
	// 在实际实现中，这里应该：
	// 1. 解析仓库URL获取用户名和仓库名
	// 2. 调用GitHub/Gitee API检查star状态
	// 3. 返回任意一个仓库已star的结果（支持单个仓库验证）

	log.Printf("检查仓库star状态，仓库数量: %d", len(repositories))
	for _, repo := range repositories {
		log.Printf("仓库: %s - %s", repo.Platform, repo.URL)
	}

	// Mock逻辑：随机返回结果
	rand.Seed(time.Now().UnixNano()) // 程序里只需要调用一次
	return rand.Intn(10) < 7         // 0~9 随机数，有70%概率 <7
}

/**
 * 生成简单的token
 * @author zhaoshenchen
 */
func generateToken() string {
	return fmt.Sprintf("star-auth-%d", time.Now().UnixNano())
}

/**
 * 健康检查接口
 * @author zhaoshenchen
 */
func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{
		"status": "ok",
		"time":   time.Now().Format(time.RFC3339),
	}
	json.NewEncoder(w).Encode(response)
}

/**
 * 添加CORS头
 * @author zhaoshenchen
 */
func enableCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
}

/**
 * 路由处理器
 * @author zhaoshenchen
 */
func routeHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)

	// 处理预检请求
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	path := r.URL.Path
	method := r.Method

	switch {
	case path == "/api/check-star" && method == "POST":
		checkStarHandler(w, r)
	case path == "/api/check-star/verify" && method == "POST":
		verifyTokenHandler(w, r)
	case path == "/api/health" && method == "GET":
		healthHandler(w, r)
	default:
		http.NotFound(w, r)
	}
}

/**
 * 清理过期token的定时任务
 * @author zhaoshenchen
 */
func cleanupExpiredTokens() {
	ticker := time.NewTicker(1 * time.Hour)
	go func() {
		for range ticker.C {
			now := time.Now()
			for token, expireTime := range validTokens {
				if now.After(expireTime) {
					delete(validTokens, token)
					log.Printf("清理过期token: %s", token)
				}
			}
		}
	}()
}

func main() {
	// 启动清理任务
	cleanupExpiredTokens()

	// 设置路由
	http.HandleFunc("/", routeHandler)

	// 启动服务器
	port := ":8081"
	log.Printf("Star Auth API服务启动，监听端口: %s", port)
	log.Printf("API端点:")
	log.Printf("  POST /api/check-star - 检查star状态")
	log.Printf("  POST /api/check-star/verify - 验证token")
	log.Printf("  GET  /api/health - 健康检查")

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal("服务器启动失败:", err)
	}
}
