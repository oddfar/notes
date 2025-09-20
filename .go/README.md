# Star Auth API 服务

这是一个用于检测用户是否已Star指定GitHub/Gitee仓库的API服务。

## 功能特性

- 支持GitHub和Gitee平台
- Token验证机制
- 自动清理过期token
- CORS支持
- Mock数据支持（开发环境）

## 快速开始

### 1. 安装依赖

```bash
cd .go
go mod tidy
```

### 2. 启动服务

```bash
go run main.go
```

服务将在 `http://localhost:8081` 启动。

### 3. API端点

#### 检查Star状态
```
POST /api/check-star
Content-Type: application/json

{
  "repositories": [
    {
      "platform": "github",
      "url": "https://github.com/username/repo"
    },
    {
      "platform": "gitee",
      "url": "https://gitee.com/username/repo"
    }
  ]
}
```

响应：
```json
{
  "success": true,
  "token": "star-auth-1234567890",
  "message": "验证成功！感谢您的支持！"
}
```

#### 验证Token
```
POST /api/check-star/verify
Content-Type: application/json

{
  "token": "star-auth-1234567890"
}
```

响应：
```json
{
  "valid": true
}
```

#### 健康检查
```
GET /api/health
```

响应：
```json
{
  "status": "ok",
  "time": "2024-01-01T12:00:00Z"
}
```

## 开发说明

### Mock模式

当前实现使用Mock数据，70%概率返回验证成功。在生产环境中，需要：

1. 实现GitHub API调用
2. 实现Gitee API调用
3. 解析仓库URL获取用户名和仓库名
4. 调用相应平台的API检查Star状态

### GitHub API示例

```go
// 检查GitHub Star状态
func checkGitHubStar(owner, repo, username string) (bool, error) {
    url := fmt.Sprintf("https://api.github.com/user/starred/%s/%s", owner, repo)
    req, _ := http.NewRequest("GET", url, nil)
    req.Header.Set("Authorization", "token YOUR_GITHUB_TOKEN")
    
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return false, err
    }
    defer resp.Body.Close()
    
    return resp.StatusCode == 204, nil
}
```

### Gitee API示例

```go
// 检查Gitee Star状态
func checkGiteeStar(owner, repo, username string) (bool, error) {
    url := fmt.Sprintf("https://gitee.com/api/v5/user/starred/%s/%s", owner, repo)
    req, _ := http.NewRequest("GET", url, nil)
    req.Header.Set("Authorization", "token YOUR_GITEE_TOKEN")
    
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return false, err
    }
    defer resp.Body.Close()
    
    return resp.StatusCode == 200, nil
}
```

## 配置说明

- **端口**: 默认8081，可通过环境变量修改
- **Token有效期**: 24小时
- **清理间隔**: 1小时
- **CORS**: 允许所有域名（生产环境需要限制）

## 部署建议

1. 使用环境变量管理敏感配置
2. 配置反向代理（Nginx）
3. 启用HTTPS
4. 限制CORS域名
5. 添加日志记录
6. 配置监控和告警