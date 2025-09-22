# Star Auth API 服务

这是一个基于Gin框架的完整系统级别项目，提供Star检测和语录API功能。

## 功能特性

- 🌟 **Star检测**: 支持GitHub和Gitee平台的Star状态检测（独立模块）
- 📝 **语录API**: 支持多种类型的语录获取，包括文件和数据库来源
  - 文件类型：a-l（动画、漫画、游戏等12种分类）
  - 数据库类型：1003（网易云热评）、1006（舔狗日记）
  - 🚀 **智能缓存**: 文件缓存30分钟，数据库缓存5分钟，显著提升性能
- 🔐 **Token验证**: 完整的Token生成和验证机制
- 🗄️ **数据库支持**: MySQL数据库集成，支持多表查询
- 🌐 **CORS支持**: 跨域请求支持
- 📁 **模块化架构**: 严格的模块化设计，低耦合高内聚
- ⚙️ **配置管理**: 环境变量配置支持
- 🔧 **水印解析**: 支持多平台视频/图片水印去除（预留接口）
- 🎯 **统一API入口**: 采用method参数路由的统一接口设计
- 📊 **调用统计**: 预留接口调用统计功能架构

## 项目结构

```
.go/
├── config/          # 配置管理
├── database/        # 数据库连接
├── modules/         # 业务模块
│   ├── star/        # Star检测模块（独立模块）
│   ├── quotes/      # 语录模块
│   ├── watermark/   # 水印解析模块
│   └── common/      # 通用模块（统一API入口）
├── routes/          # 路由配置
├── doc/             # 文档和数据文件
│   └── sentences/   # 语录JSON文件
├── main.go          # 主程序入口
├── go.mod           # Go模块文件
└── .env.example     # 环境配置示例
```

## 快速开始

### 1. 环境配置

复制环境配置文件并修改：
```bash
cp .env.example .env
# 编辑 .env 文件，配置数据库和API Token
```

### 2. 安装依赖

```bash
go mod tidy
```

### 3. 启动服务

```bash
go run main.go
```

服务将在 `http://localhost:8080` 启动。

## API文档

### Star检测相关

#### 检查Star状态
```
POST /api/check-star
Content-Type: application/json

{
  "username": "用户名",
  "repositories": [
    {
      "platform": "github",
      "url": "https://github.com/username/repo"
    }
  ]
}
```

响应：
```json
{
  "code": "200",
  "message": "Success",
  "data": {
    "https://github.com/username/repo": {
      "starred": true,
      "message": "Repository star status"
    }
  },
  "token": "generated-token-if-all-starred"
}
```

#### 验证Token
```
POST /api/verify-token
Content-Type: application/json

{
  "token": "your-token"
}
```

响应：
```json
{
  "code": "200",
  "message": "Success",
  "valid": true
}
```

### 语录API

#### 获取语录
```
GET /api/quote?c=a&encode=json
```

参数说明：
- `c`: 语录类型，支持 a-l（对应不同分类）、1003（网易云热评）、其他（默认为动画类型）
- `encode`: 返回格式，`text`（纯文本）或 `json`（JSON格式）

响应（JSON格式）：
```json
{
  "code": "200",
  "msg": "success",
  "type": "a",
  "data": {
    "text": "语录内容"
  },
  "from": "xxxapi"
}
```

响应（文本格式）：
```
语录内容
```

#### 健康检查
```
GET /api/health
```

响应：
```json
{
  "status": "ok",
  "message": "Service is running"
}
```

### 统一API入口

#### 统一接口调用
```
POST /api/unified
Content-Type: application/json

{
  "method": "oddfar.api.quote",
  "param": {},
  "body": {
    "c": "a",
    "encode": "json"
  }
}
```

支持的method参数：
- `oddfar.api.quote`: 语录接口
- `oddfar.api.watermark`: 水印解析接口

#### 接口调用统计
```
GET /api/stats?method=oddfar.api.quote&start_time=2024-01-01&end_time=2024-01-02
```

响应：
```json
{
  "code": "200",
  "message": "success",
  "data": {
    "method": "oddfar.api.quote",
    "total_calls": 1000,
    "start_time": "2024-01-01",
    "end_time": "2024-01-02",
    "avg_duration": 25.5
  }
}
```

### 水印解析API

#### 水印解析
```
POST /api/watermark
Content-Type: application/json

{
  "url": "https://example.com/video"
}
```

支持平台：
- 抖音 (douyin.com)
- 快手 (kuaishou.com) 
- 小红书 (xiaohongshu.com)
- 微博 (weibo.com)
- 其他平台（通用处理）

响应：
```json
{
  "code": "200",
  "message": "success",
  "data": {
    "original_url": "原始URL",
    "clean_url": "去水印后的URL",
    "title": "内容标题",
    "author": "作者",
    "platform": "检测到的平台"
  }
}
```

## 开发说明

### 缓存机制

- **文件缓存**: 30分钟过期，避免频繁读取JSON文件
- **数据库缓存**: 5分钟过期，减少数据库查询压力
- **自动清理**: 每10分钟自动清理过期缓存
- **线程安全**: 使用读写锁保证并发安全

### 统一API设计

采用method参数路由设计，便于接口管理和扩展：
```json
{
  "method": "oddfar.api.接口标识",
  "param": "公共请求参数",
  "body": "业务请求参数"
}
```

### 接口调用统计

预留统计功能架构，支持：
- 接口调用次数统计
- 响应时间统计
- 客户端IP记录
- 时间段查询

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