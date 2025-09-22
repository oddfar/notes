package database

import (
	"database/sql"
	"fmt"
	"log"
	"star-auth-api/config"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

/**
 * 初始化数据库连接
 * @author oddfar
 */
func InitDB(cfg *config.Config) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Database,
	)

	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("连接数据库失败: %v", err)
	}

	// 测试连接
	if err = DB.Ping(); err != nil {
		return fmt.Errorf("数据库连接测试失败: %v", err)
	}

	log.Println("数据库连接成功")
	return nil
}

/**
 * 关闭数据库连接
 * @author oddfar
 */
func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}
