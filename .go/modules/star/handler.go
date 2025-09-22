package star

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/**
 * Star检测处理器
 * @author oddfar
 */
type StarHandler struct {
	service *StarService
}

/**
 * 创建Star处理器实例
 * @author oddfar
 */
func NewStarHandler(service *StarService) *StarHandler {
	return &StarHandler{
		service: service,
	}
}

/**
 * 检查star状态接口
 * @author oddfar
 */
func (h *StarHandler) CheckStar(c *gin.Context) {
	var req CheckStarRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "请求参数错误",
		})
		return
	}

	isStarred := h.service.CheckRepositoriesStarStatus(req.Repositories)

	if isStarred {
		token := h.service.GenerateToken()
		c.JSON(http.StatusOK, CheckStarResponse{
			Success: true,
			Token:   token,
			Message: "验证成功！感谢您的支持！",
		})
	} else {
		c.JSON(http.StatusOK, CheckStarResponse{
			Success: false,
			Message: "请先Star任意一个项目后再进行验证",
		})
	}
}

/**
 * 验证token接口
 * @author oddfar
 */
func (h *StarHandler) VerifyToken(c *gin.Context) {
	var req VerifyTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"valid": false,
		})
		return
	}

	valid := h.service.VerifyToken(req.Token)
	c.JSON(http.StatusOK, VerifyTokenResponse{
		Valid: valid,
	})
}
