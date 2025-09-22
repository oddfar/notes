package common

/**
 * 统一API请求参数
 * @author zhaoshenchen
 */
type UnifiedAPIRequest struct {
	Method string      `json:"method" form:"method" binding:"required"` // 接口标识，格式：oddfar.api.xxx
	Param  interface{} `json:"param,omitempty" form:"param"`            // 公共请求参数
	Body   interface{} `json:"body,omitempty"`                          // 业务请求参数
}

/**
 * 统一API响应结构
 * @author zhaoshenchen
 */
type UnifiedAPIResponse struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Method  string      `json:"method,omitempty"`
}

/**
 * 接口调用统计记录
 * @author zhaoshenchen
 */
type APICallRecord struct {
	ID        int64  `json:"id"`
	Method    string `json:"method"`
	IP        string `json:"ip"`
	UserAgent string `json:"user_agent"`
	CallTime  string `json:"call_time"`
	Duration  int64  `json:"duration"` // 毫秒
	Status    string `json:"status"`   // success, error
}

/**
 * 接口调用统计查询参数
 * @author zhaoshenchen
 */
type APIStatsRequest struct {
	Method    string `json:"method" form:"method"`
	StartTime string `json:"start_time" form:"start_time"`
	EndTime   string `json:"end_time" form:"end_time"`
	IP        string `json:"ip" form:"ip"`
}

/**
 * 接口调用统计响应
 * @author zhaoshenchen
 */
type APIStatsResponse struct {
	Method     string `json:"method"`
	TotalCalls int64  `json:"total_calls"`
	StartTime  string `json:"start_time"`
	EndTime    string `json:"end_time"`
	AvgDuration float64 `json:"avg_duration"`
}