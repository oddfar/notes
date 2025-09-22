package quotes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"path/filepath"
	"star-auth-api/database"
	"time"
)

/**
 * 语录服务
 * @author zhaoshenchen
 */
type QuoteService struct {
	sentencesPath string
	cache         *QuoteCache
}

/**
 * 创建语录服务实例
 * @param sentencesPath 语录文件路径
 * @author zhaoshenchen
 */
func NewQuoteService(sentencesPath string) *QuoteService {
	service := &QuoteService{
		sentencesPath: sentencesPath,
		cache:         NewQuoteCache(),
	}
	
	// 启动缓存清理协程
	go service.startCacheCleanup()
	
	return service
}

/**
 * 获取语录
 * @author oddfar
 */
func (s *QuoteService) GetQuote(quoteType, encode string) (interface{}, error) {
	// 设置默认值
	if quoteType == "" {
		quoteType = "2001"
	}
	if encode == "" {
		encode = "text"
	}

	var text string
	var err error

	// 根据类型获取语录
	if quoteType == "1003" {
		text, err = s.getQuoteFromDB(quoteType)
	} else {
		text, err = s.getQuoteFromFile(quoteType)
	}

	if err != nil {
		return nil, err
	}

	// 根据编码格式返回结果
	if encode == "json" {
		return QuoteResponse{
			Code: "200",
			Msg:  "success",
			Type: quoteType,
			Data: QuoteData{Text: text},
			From: "xxxapi",
		}, nil
	}

	return text, nil
}

/**
 * 从数据库获取语录
 * @param tableName 表名
 * @author zhaoshenchen
 */
func (s *QuoteService) getQuoteFromDB(tableName string) (string, error) {
	if database.DB == nil {
		return "", fmt.Errorf("数据库连接未初始化")
	}

	// 先尝试从缓存获取
	if quotes, found := s.cache.GetDBCache(tableName); found {
		if len(quotes) > 0 {
			rand.Seed(time.Now().UnixNano())
			randomIndex := rand.Intn(len(quotes))
			return quotes[randomIndex].Text, nil
		}
	}

	// 缓存未命中，从数据库查询所有数据并缓存
	querySQL := fmt.Sprintf("SELECT Id, text FROM %s WHERE zt = 0", tableName)
	rows, err := database.DB.Query(querySQL)
	if err != nil {
		return "", fmt.Errorf("查询语录失败: %v", err)
	}
	defer rows.Close()

	var quotes []DBQuote
	for rows.Next() {
		var quote DBQuote
		if err := rows.Scan(&quote.ID, &quote.Text); err != nil {
			continue // 跳过错误行
		}
		quotes = append(quotes, quote)
	}

	if len(quotes) == 0 {
		return "暂无语录数据", nil
	}

	// 将数据存入缓存
	s.cache.SetDBCache(tableName, quotes)

	// 随机选择一条
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(quotes))
	
	return quotes[randomIndex].Text, nil
}

/**
 * 从文件获取语录
 * @author oddfar
 */
func (s *QuoteService) getQuoteFromFile(quoteType string) (string, error) {
	// 映射类型到文件名
	fileName := s.mapTypeToFileName(quoteType)
	
	// 先尝试从缓存获取
	if quotes, found := s.cache.GetFileCache(fileName); found {
		if len(quotes) > 0 {
			rand.Seed(time.Now().UnixNano())
			randomIndex := rand.Intn(len(quotes))
			return quotes[randomIndex].Hitokoto, nil
		}
	}
	
	// 缓存未命中，从文件读取
	filePath := filepath.Join(s.sentencesPath, fileName+".json")

	// 读取文件
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("读取语录文件失败: %v", err)
	}

	// 解析JSON
	var quotes []FileQuote
	if err := json.Unmarshal(data, &quotes); err != nil {
		return "", fmt.Errorf("解析语录文件失败: %v", err)
	}

	if len(quotes) == 0 {
		return "暂无语录数据", nil
	}

	// 将数据存入缓存
	s.cache.SetFileCache(fileName, quotes)

	// 随机选择一条
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(quotes))

	return quotes[randomIndex].Hitokoto, nil
}

/**
 * 映射类型到文件名
 * @author oddfar
 */
func (s *QuoteService) mapTypeToFileName(quoteType string) string {
	typeMap := map[string]string{
		"a": "a", // 动画
		"b": "b", // 漫画
		"c": "c", // 游戏
		"d": "d", // 文学
		"e": "e", // 原创
		"f": "f", // 来自网络
		"g": "g", // 其他
		"h": "h", // 影视
		"i": "i", // 诗词
		"j": "j", // 网易云
		"k": "k", // 哲学
		"l": "l", // 抖机灵
	}

	if fileName, exists := typeMap[quoteType]; exists {
		return fileName
	}

	// 其他类型作为动画类型处理
	return "a"
}

/**
 * 启动缓存清理协程
 * @author zhaoshenchen
 */
func (s *QuoteService) startCacheCleanup() {
	ticker := time.NewTicker(10 * time.Minute) // 每10分钟清理一次
	defer ticker.Stop()
	
	for {
		select {
		case <-ticker.C:
			s.cache.CleanExpiredCache()
		}
	}
}
