package watermark

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/**
 * 水印解析处理器
 * @author zhaoshenchen
 */
type WatermarkHandler struct {
	service *WatermarkService
}

/**
 * 创建水印解析处理器实例
 * @param service 水印解析服务实例
 * @author zhaoshenchen
 */
func NewWatermarkHandler(service *WatermarkService) *WatermarkHandler {
	return &WatermarkHandler{
		service: service,
	}
}

/**
 * 水印解析接口
 * @author zhaoshenchen
 */
func (h *WatermarkHandler) ParseWatermark(c *gin.Context) {
	var req WatermarkRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, APIResponse{
			Code:    "400",
			Message: "请求参数错误: " + err.Error(),
		})
		return
	}

	result, err := h.service.ParseWatermark(req.URL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, APIResponse{
			Code:    "500",
			Message: "水印解析失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, APIResponse{
		Code:    "200",
		Message: "success",
		Data:    result,
	})
}