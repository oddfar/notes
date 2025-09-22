package common

import (
	"encoding/json"
	"fmt"
	"star-auth-api/modules/quotes"
	"star-auth-api/modules/watermark"
)

/**
 * 统一API服务
 * @author zhaoshenchen
 */
type UnifiedAPIService struct {
	quoteService     *quotes.QuoteService
	watermarkService *watermark.WatermarkService
	// TODO: 添加接口调用统计服务
}

/**
 * 创建统一API服务实例
 * @param quoteService 语录服务
 * @param watermarkService 水印解析服务
 * @author zhaoshenchen
 */
func NewUnifiedAPIService(quoteService *quotes.QuoteService, watermarkService *watermark.WatermarkService) *UnifiedAPIService {
	return &UnifiedAPIService{
		quoteService:     quoteService,
		watermarkService: watermarkService,
	}
}

/**
 * 处理统一API请求
 * @param request 统一API请求
 * @param clientIP 客户端IP
 * @param userAgent 用户代理
 * @return 处理结果
 * @author zhaoshenchen
 */
func (s *UnifiedAPIService) ProcessRequest(request *UnifiedAPIRequest, clientIP, userAgent string) (*UnifiedAPIResponse, error) {
	// TODO: 记录接口调用开始
	// startTime := time.Now()
	// s.recordAPICall(request.Method, clientIP, userAgent, startTime)
	
	var result interface{}
	var err error
	
	// 根据method参数路由到不同的服务
	switch request.Method {
	case "oddfar.api.quote":
		result, err = s.handleQuoteRequest(request.Body)
	case "oddfar.api.watermark":
		result, err = s.handleWatermarkRequest(request.Body)
	default:
		return &UnifiedAPIResponse{
			Code:    "400",
			Message: fmt.Sprintf("不支持的接口方法: %s", request.Method),
			Method:  request.Method,
		}, nil
	}
	
	// TODO: 记录接口调用结束
	// duration := time.Since(startTime).Milliseconds()
	// s.recordAPICallEnd(request.Method, duration, err == nil)
	
	if err != nil {
		return &UnifiedAPIResponse{
			Code:    "500",
			Message: err.Error(),
			Method:  request.Method,
		}, nil
	}
	
	return &UnifiedAPIResponse{
		Code:    "200",
		Message: "success",
		Data:    result,
		Method:  request.Method,
	}, nil
}

/**
 * 处理语录请求
 * @param body 请求体
 * @return 语录结果
 * @author zhaoshenchen
 */
func (s *UnifiedAPIService) handleQuoteRequest(body interface{}) (interface{}, error) {
	// 解析请求参数
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("解析请求参数失败: %v", err)
	}
	
	var req quotes.QuoteRequest
	if err := json.Unmarshal(bodyBytes, &req); err != nil {
		return nil, fmt.Errorf("解析语录请求参数失败: %v", err)
	}
	
	// 调用语录服务
	return s.quoteService.GetQuote(req.Type, req.Encode)
}

/**
 * 处理水印解析请求
 * @param body 请求体
 * @return 水印解析结果
 * @author zhaoshenchen
 */
func (s *UnifiedAPIService) handleWatermarkRequest(body interface{}) (interface{}, error) {
	// 解析请求参数
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("解析请求参数失败: %v", err)
	}
	
	var req watermark.WatermarkRequest
	if err := json.Unmarshal(bodyBytes, &req); err != nil {
		return nil, fmt.Errorf("解析水印请求参数失败: %v", err)
	}
	
	// 调用水印解析服务
	return s.watermarkService.ParseWatermark(req.URL)
}

/**
 * 获取接口调用统计
 * @param request 统计查询请求
 * @return 统计结果
 * @author zhaoshenchen
 */
func (s *UnifiedAPIService) GetAPIStats(request *APIStatsRequest) (*APIStatsResponse, error) {
	// TODO: 实现接口调用统计查询逻辑
	// 这里可以从数据库查询指定时间段内的接口调用统计
	
	return &APIStatsResponse{
		Method:      request.Method,
		TotalCalls:  0,
		StartTime:   request.StartTime,
		EndTime:     request.EndTime,
		AvgDuration: 0.0,
	}, nil
}