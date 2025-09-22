package quotes

import (
	"sync"
	"time"
)

/**
 * 缓存项
 * @author zhaoshenchen
 */
type CacheItem struct {
	Data      []FileQuote `json:"data"`
	Timestamp int64       `json:"timestamp"`
}

/**
 * 语录缓存服务
 * @author zhaoshenchen
 */
type QuoteCache struct {
	fileCache    map[string]*CacheItem // 文件缓存
	dbCache      map[string][]DBQuote  // 数据库缓存
	mutex        sync.RWMutex          // 读写锁
	cacheTTL     time.Duration         // 缓存过期时间
	lastDBUpdate map[string]int64      // 数据库缓存最后更新时间
}

/**
 * 创建语录缓存实例
 * @author zhaoshenchen
 */
func NewQuoteCache() *QuoteCache {
	return &QuoteCache{
		fileCache:    make(map[string]*CacheItem),
		dbCache:      make(map[string][]DBQuote),
		cacheTTL:     30 * time.Minute, // 缓存30分钟
		lastDBUpdate: make(map[string]int64),
	}
}

/**
 * 获取文件缓存
 * @param fileName 文件名
 * @return 缓存的语录数据
 * @author zhaoshenchen
 */
func (c *QuoteCache) GetFileCache(fileName string) ([]FileQuote, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	
	item, exists := c.fileCache[fileName]
	if !exists {
		return nil, false
	}
	
	// 检查缓存是否过期
	if time.Now().Unix()-item.Timestamp > int64(c.cacheTTL.Seconds()) {
		return nil, false
	}
	
	return item.Data, true
}

/**
 * 设置文件缓存
 * @param fileName 文件名
 * @param data 语录数据
 * @author zhaoshenchen
 */
func (c *QuoteCache) SetFileCache(fileName string, data []FileQuote) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	
	c.fileCache[fileName] = &CacheItem{
		Data:      data,
		Timestamp: time.Now().Unix(),
	}
}

/**
 * 获取数据库缓存
 * @param tableName 表名
 * @return 缓存的语录数据
 * @author zhaoshenchen
 */
func (c *QuoteCache) GetDBCache(tableName string) ([]DBQuote, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	
	data, exists := c.dbCache[tableName]
	if !exists {
		return nil, false
	}
	
	// 检查缓存是否过期（数据库缓存时间较短，5分钟）
	lastUpdate, updateExists := c.lastDBUpdate[tableName]
	if !updateExists || time.Now().Unix()-lastUpdate > 300 { // 5分钟
		return nil, false
	}
	
	return data, true
}

/**
 * 设置数据库缓存
 * @param tableName 表名
 * @param data 语录数据
 * @author zhaoshenchen
 */
func (c *QuoteCache) SetDBCache(tableName string, data []DBQuote) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	
	c.dbCache[tableName] = data
	c.lastDBUpdate[tableName] = time.Now().Unix()
}

/**
 * 清理过期缓存
 * @author zhaoshenchen
 */
func (c *QuoteCache) CleanExpiredCache() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	
	now := time.Now().Unix()
	
	// 清理文件缓存
	for fileName, item := range c.fileCache {
		if now-item.Timestamp > int64(c.cacheTTL.Seconds()) {
			delete(c.fileCache, fileName)
		}
	}
	
	// 清理数据库缓存
	for tableName, lastUpdate := range c.lastDBUpdate {
		if now-lastUpdate > 300 { // 5分钟
			delete(c.dbCache, tableName)
			delete(c.lastDBUpdate, tableName)
		}
	}
}