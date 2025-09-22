package star

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"star-auth-api/config"
	"strings"
	"time"
)

/**
 * Star检测服务
 * @author oddfar
 */
type StarService struct {
	config      *config.Config
	validTokens map[string]time.Time
}

/**
 * 创建Star服务实例
 * @author oddfar
 */
func NewStarService(cfg *config.Config) *StarService {
	service := &StarService{
		config:      cfg,
		validTokens: make(map[string]time.Time),
	}

	// 启动清理任务
	service.startCleanupTask()

	return service
}

/**
 * 检查仓库star状态
 * @author oddfar
 */
func (s *StarService) CheckRepositoriesStarStatus(repositories []Repository) bool {
	for _, repo := range repositories {
		isStarred, err := s.checkSingleRepository(repo)
		if err != nil {
			log.Printf("检查仓库 %s 失败: %v", repo.URL, err)
			continue
		}
		if isStarred {
			return true // 只要有一个仓库被star就返回true
		}
	}
	return false
}

/**
 * 检查单个仓库的star状态
 * @author oddfar
 */
func (s *StarService) checkSingleRepository(repo Repository) (bool, error) {
	owner, repoName, err := s.parseRepositoryURL(repo.URL)
	if err != nil {
		return false, err
	}

	switch strings.ToLower(repo.Platform) {
	case "github":
		return s.checkGitHubStar(owner, repoName)
	case "gitee":
		return s.checkGiteeStar(owner, repoName)
	default:
		return false, fmt.Errorf("不支持的平台: %s", repo.Platform)
	}
}

/**
 * 解析仓库URL获取owner和repo名称
 * @author oddfar
 */
func (s *StarService) parseRepositoryURL(url string) (string, string, error) {
	// GitHub: https://github.com/owner/repo
	// Gitee: https://gitee.com/owner/repo

	githubRegex := regexp.MustCompile(`https://github\.com/([^/]+)/([^/]+)`)
	giteeRegex := regexp.MustCompile(`https://gitee\.com/([^/]+)/([^/]+)`)

	if matches := githubRegex.FindStringSubmatch(url); len(matches) == 3 {
		return matches[1], matches[2], nil
	}

	if matches := giteeRegex.FindStringSubmatch(url); len(matches) == 3 {
		return matches[1], matches[2], nil
	}

	return "", "", fmt.Errorf("无效的仓库URL: %s", url)
}

/**
 * 检查GitHub star状态
 * @author oddfar
 */
func (s *StarService) checkGitHubStar(owner, repo string) (bool, error) {
	if s.config.GitHub.Token == "" {
		log.Println("GitHub token未配置，使用mock数据")
		return s.mockStarCheck(), nil
	}

	url := fmt.Sprintf("https://api.github.com/user/starred/%s/%s", owner, repo)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return false, err
	}

	req.Header.Set("Authorization", "token "+s.config.GitHub.Token)
	req.Header.Set("Accept", "application/vnd.github.v3+json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	// 204表示已star，404表示未star
	return resp.StatusCode == 204, nil
}

/**
 * 检查Gitee star状态
 * @author oddfar
 */
func (s *StarService) checkGiteeStar(owner, repo string) (bool, error) {
	if s.config.Gitee.Token == "" {
		log.Println("Gitee token未配置，使用mock数据")
		return s.mockStarCheck(), nil
	}

	url := fmt.Sprintf("https://gitee.com/api/v5/user/starred/%s/%s", owner, repo)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return false, err
	}

	req.Header.Set("Authorization", "token "+s.config.Gitee.Token)

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	// 200表示已star，404表示未star
	return resp.StatusCode == 200, nil
}

/**
 * Mock star检查（用于测试）
 * @author oddfar
 */
func (s *StarService) mockStarCheck() bool {
	// 70%概率返回true
	return time.Now().UnixNano()%10 < 7
}

/**
 * 生成token
 * @author oddfar
 */
func (s *StarService) GenerateToken() string {
	token := fmt.Sprintf("star-auth-%d", time.Now().UnixNano())
	s.validTokens[token] = time.Now().Add(24 * time.Hour)
	return token
}

/**
 * 验证token
 * @author oddfar
 */
func (s *StarService) VerifyToken(token string) bool {
	expireTime, exists := s.validTokens[token]
	valid := exists && time.Now().Before(expireTime)

	// 清理过期token
	if exists && !valid {
		delete(s.validTokens, token)
	}

	return valid
}

/**
 * 启动清理过期token的定时任务
 * @author oddfar
 */
func (s *StarService) startCleanupTask() {
	ticker := time.NewTicker(1 * time.Hour)
	go func() {
		for range ticker.C {
			now := time.Now()
			for token, expireTime := range s.validTokens {
				if now.After(expireTime) {
					delete(s.validTokens, token)
					log.Printf("清理过期token: %s", token)
				}
			}
		}
	}()
}
