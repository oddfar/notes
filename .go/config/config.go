package config

import (
	"os"
)

/**
 * 应用配置
 * @author oddfar
 */
type Config struct {
	Server   ServerConfig   `json:"server"`
	Database DatabaseConfig `json:"database"`
	GitHub   GitHubConfig   `json:"github"`
	Gitee    GiteeConfig    `json:"gitee"`
}

/**
 * 服务器配置
 * @author oddfar
 */
type ServerConfig struct {
	Port string `json:"port"`
	Mode string `json:"mode"`
}

/**
 * 数据库配置
 * @author oddfar
 */
type DatabaseConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
}

/**
 * GitHub配置
 * @author oddfar
 */
type GitHubConfig struct {
	Token string `json:"token"`
}

/**
 * Gitee配置
 * @author oddfar
 */
type GiteeConfig struct {
	Token string `json:"token"`
}

/**
 * 获取配置
 * @author oddfar
 */
func GetConfig() *Config {
	return &Config{
		Server: ServerConfig{
			Port: getEnv("SERVER_PORT", "8081"),
			Mode: getEnv("GIN_MODE", "debug"),
		},
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "3306"),
			Username: getEnv("DB_USERNAME", "root"),
			Password: getEnv("DB_PASSWORD", ""),
			Database: getEnv("DB_DATABASE", "test"),
		},
		GitHub: GitHubConfig{
			Token: getEnv("GITHUB_TOKEN", ""),
		},
		Gitee: GiteeConfig{
			Token: getEnv("GITEE_TOKEN", ""),
		},
	}
}

/**
 * 获取环境变量，如果不存在则返回默认值
 * @author oddfar
 */
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
