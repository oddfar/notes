package star

import "time"

/**
 * 仓库信息
 * @author oddfar
 */
type Repository struct {
	Platform string `json:"platform"`
	URL      string `json:"url"`
}

/**
 * 检查star请求
 * @author oddfar
 */
type CheckStarRequest struct {
	Repositories []Repository `json:"repositories"`
}

/**
 * 检查star响应
 * @author oddfar
 */
type CheckStarResponse struct {
	Success bool   `json:"success"`
	Token   string `json:"token,omitempty"`
	Message string `json:"message"`
}

/**
 * 验证token请求
 * @author oddfar
 */
type VerifyTokenRequest struct {
	Token string `json:"token"`
}

/**
 * 验证token响应
 * @author oddfar
 */
type VerifyTokenResponse struct {
	Valid bool `json:"valid"`
}

/**
 * Token存储结构
 * @author oddfar
 */
type TokenInfo struct {
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
}
