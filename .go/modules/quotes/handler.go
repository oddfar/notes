package quotes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/**
 * 语录处理器
 * @author oddfar
 */
type QuoteHandler struct {
	service *QuoteService
}

/**
 * 创建语录处理器实例
 * @author oddfar
 */
func NewQuoteHandler(service *QuoteService) *QuoteHandler {
	return &QuoteHandler{
		service: service,
	}
}

/**
 * 获取语录接口
 * @author oddfar
 */
func (h *QuoteHandler) GetQuote(c *gin.Context) {
	var req QuoteRequest

	// 绑定查询参数
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": "400",
			"msg":  "参数错误",
		})
		return
	}

	// 获取语录
	result, err := h.service.GetQuote(req.Type, req.Encode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": "500",
			"msg":  "获取语录失败: " + err.Error(),
		})
		return
	}

	// 根据编码格式返回结果
	if req.Encode == "json" {
		c.JSON(http.StatusOK, result)
	} else {
		c.String(http.StatusOK, result.(string))
	}
}
