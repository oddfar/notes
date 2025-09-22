package watermark

/**
 * 水印解析请求参数
 * @author zhaoshenchen
 */
type WatermarkRequest struct {
	URL string `json:"url" form:"url" binding:"required"`
}

/**
 * 水印解析响应数据
 * @author zhaoshenchen
 */
type WatermarkResponse struct {
	OriginalURL string `json:"original_url"`
	CleanURL    string `json:"clean_url"`
	Title       string `json:"title,omitempty"`
	Author      string `json:"author,omitempty"`
	Platform    string `json:"platform,omitempty"`
}

/**
 * 通用API响应结构
 * @author zhaoshenchen
 */
type APIResponse struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}