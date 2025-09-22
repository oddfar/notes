package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/**
 * 统一API处理器
 * @author zhaoshenchen
 */
type UnifiedAPIHandler struct {
	service *UnifiedAPIService
}

/**
 * 创建统一API处理器实例
 * @param service 统一API服务实例
 * @author zhaoshenchen
 */
func NewUnifiedAPIHandler(service *UnifiedAPIService) *UnifiedAPIHandler {
	return &UnifiedAPIHandler{
		service: service,
	}
}

/**
 * 统一API入口
 * @author zhaoshenchen
 */
func (h *UnifiedAPIHandler) HandleUnifiedAPI(c *gin.Context) {
	var req UnifiedAPIRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, UnifiedAPIResponse{
			Code:    "400",
			Message: "请求参数错误: " + err.Error(),
		})
		return
	}

	// 获取客户端信息
	clientIP := c.ClientIP()
	userAgent := c.GetHeader("User-Agent")

	result, err := h.service.ProcessRequest(&req, clientIP, userAgent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, UnifiedAPIResponse{
			Code:    "500",
			Message: "处理请求失败: " + err.Error(),
			Method:  req.Method,
		})
		return
	}

	c.JSON(http.StatusOK, result)
}

/**
 * 获取接口调用统计
 * @author zhaoshenchen
 */
func (h *UnifiedAPIHandler) GetAPIStats(c *gin.Context) {
	var req APIStatsRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, UnifiedAPIResponse{
			Code:    "400",
			Message: "请求参数错误: " + err.Error(),
		})
		return
	}

	result, err := h.service.GetAPIStats(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, UnifiedAPIResponse{
			Code:    "500",
			Message: "获取统计失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, UnifiedAPIResponse{
		Code:    "200",
		Message: "success",
		Data:    result,
	})
}