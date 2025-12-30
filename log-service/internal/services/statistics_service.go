// Package services 统计业务逻辑服务层
// 实现基于 authorization 字段的用户请求统计数据
package services

import (
	"errors"
	"time"
	"zxm_ai_admin/log-service/internal/database"
	"zxm_ai_admin/log-service/internal/models"

	"gorm.io/gorm"
)

// StatisticsService 统计服务
type StatisticsService struct{}

// NewStatisticsService 创建统计服务实例
func NewStatisticsService() *StatisticsService {
	return &StatisticsService{}
}

// GetUserStatisticsRequest 获取用户统计数据请求
type GetUserStatisticsRequest struct {
	Authorization string `form:"authorization" binding:"required"`
	StartTime     string `form:"start_time"`
	EndTime       string `form:"end_time"`
}

// UserStatisticsResponse 用户统计数据响应
type UserStatisticsResponse struct {
	Authorization string             `json:"authorization"`
	TimeRange     TimeRange          `json:"time_range"`
	Summary       SummaryStatistics  `json:"summary"`
	Latency       LatencyStatistics  `json:"latency"`
	ByIP          []IPStatistics     `json:"by_ip"`
	ByPath        []PathStatistics   `json:"by_path"`
	ByDate        []DateStatistics   `json:"by_date"`
	ByTime        []TimeStatistics   `json:"by_time"`
}

// TimeRange 时间范围
type TimeRange struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

// SummaryStatistics 汇总统计
type SummaryStatistics struct {
	TotalRequests       int64 `json:"total_requests"`
	TotalRequestBytes   int64 `json:"total_request_bytes"`
	TotalResponseBytes  int64 `json:"total_response_bytes"`
	AvgRequestBytes     int64 `json:"avg_request_bytes"`
	AvgResponseBytes    int64 `json:"avg_response_bytes"`
}

// LatencyStatistics 延迟统计
type LatencyStatistics struct {
	TotalMs int64   `json:"total_ms"`
	AvgMs   float64 `json:"avg_ms"`
	MinMs   int64   `json:"min_ms"`
	MaxMs   int64   `json:"max_ms"`
}

// IPStatistics IP统计
type IPStatistics struct {
	IP    string `json:"ip"`
	Count int64  `json:"count"`
}

// PathStatistics 路径统计
type PathStatistics struct {
	Path  string `json:"path"`
	Count int64  `json:"count"`
}

// DateStatistics 日期统计
type DateStatistics struct {
	Date  string `json:"date"`
	Count int64  `json:"count"`
}

// TimeStatistics 小时统计
type TimeStatistics struct {
	Time  string `json:"time"`
	Count int64  `json:"count"`
}

// GetAuthorizationRankingRequest 获取排行榜请求
type GetAuthorizationRankingRequest struct {
	StartTime string `form:"start_time"`
	EndTime   string `form:"end_time"`
	Page      int    `form:"page"`
	PageSize  int    `form:"page_size"`
}

// AuthorizationRankingResponse 排行榜响应
type AuthorizationRankingResponse struct {
	Total     int64                       `json:"total"`
	TimeRange TimeRange                   `json:"time_range"`
	List      []AuthorizationRankingItem `json:"list"`
}

// AuthorizationRankingItem 排行榜单项
type AuthorizationRankingItem struct {
	Authorization string `json:"authorization"`
	Count         int64  `json:"count"`
}

// GetUserStatistics 获取用户统计数据
func (s *StatisticsService) GetUserStatistics(req *GetUserStatisticsRequest) (*UserStatisticsResponse, error) {
	// 解析时间范围，默认最近7天
	startTime, endTime, err := s.parseTimeRange(req.StartTime, req.EndTime)
	if err != nil {
		return nil, err
	}

	// 构建基础查询
	baseQuery := database.DB.Model(&models.TokenUsageLog{}).
		Where("authorization = ?", req.Authorization).
		Where("time >= ?", startTime).
		Where("time <= ?", endTime)

	// 获取汇总统计
	summary, err := s.getSummary(baseQuery)
	if err != nil {
		return nil, errors.New("获取汇总统计失败: " + err.Error())
	}

	// 获取延迟统计
	latency, err := s.getLatency(baseQuery)
	if err != nil {
		return nil, errors.New("获取延迟统计失败: " + err.Error())
	}

	// 按IP分组统计
	byIP, err := s.groupByIP(baseQuery)
	if err != nil {
		return nil, errors.New("获取IP统计失败: " + err.Error())
	}

	// 按路径分组统计
	byPath, err := s.groupByPath(baseQuery)
	if err != nil {
		return nil, errors.New("获取路径统计失败: " + err.Error())
	}

	// 按日期分组统计
	byDate, err := s.groupByDate(baseQuery)
	if err != nil {
		return nil, errors.New("获取日期统计失败: " + err.Error())
	}

	// 按小时分组统计
	byTime, err := s.groupByTime(baseQuery)
	if err != nil {
		return nil, errors.New("获取小时统计失败: " + err.Error())
	}

	return &UserStatisticsResponse{
		Authorization: req.Authorization,
		TimeRange: TimeRange{
			Start: startTime.Format("2006-01-02 15:04:05"),
			End:   endTime.Format("2006-01-02 15:04:05"),
		},
		Summary: *summary,
		Latency: *latency,
		ByIP:    byIP,
		ByPath:  byPath,
		ByDate:  byDate,
		ByTime:  byTime,
	}, nil
}

// parseTimeRange 解析时间范围，默认最近7天
func (s *StatisticsService) parseTimeRange(startTimeStr, endTimeStr string) (time.Time, time.Time, error) {
	var startTime, endTime time.Time
	var err error

	// 解析开始时间
	if startTimeStr != "" {
		startTime, err = time.Parse("2006-01-02 15:04:05", startTimeStr)
		if err != nil {
			return time.Time{}, time.Time{}, errors.New("开始时间格式错误，正确格式为: 2006-01-02 15:04:05")
		}
	} else {
		// 默认7天前
		startTime = time.Now().AddDate(0, 0, -7)
	}

	// 解析结束时间
	if endTimeStr != "" {
		endTime, err = time.Parse("2006-01-02 15:04:05", endTimeStr)
		if err != nil {
			return time.Time{}, time.Time{}, errors.New("结束时间格式错误，正确格式为: 2006-01-02 15:04:05")
		}
	} else {
		endTime = time.Now()
	}

	if startTime.After(endTime) {
		return time.Time{}, time.Time{}, errors.New("开始时间不能晚于结束时间")
	}

	return startTime, endTime, nil
}

// getSummary 获取汇总统计
func (s *StatisticsService) getSummary(query *gorm.DB) (*SummaryStatistics, error) {
	type Result struct {
		TotalRequests       int64
		TotalRequestBytes   int64
		TotalResponseBytes  int64
	}

	var result Result
	err := query.Select(
		"COUNT(*) as total_requests",
		"COALESCE(SUM(request_size_bytes), 0) as total_request_bytes",
		"COALESCE(SUM(response_size_bytes), 0) as total_response_bytes",
	).Scan(&result).Error

	if err != nil {
		return nil, err
	}

	avgRequestBytes := int64(0)
	avgResponseBytes := int64(0)
	if result.TotalRequests > 0 {
		avgRequestBytes = result.TotalRequestBytes / result.TotalRequests
		avgResponseBytes = result.TotalResponseBytes / result.TotalRequests
	}

	return &SummaryStatistics{
		TotalRequests:      result.TotalRequests,
		TotalRequestBytes:  result.TotalRequestBytes,
		TotalResponseBytes: result.TotalResponseBytes,
		AvgRequestBytes:    avgRequestBytes,
		AvgResponseBytes:   avgResponseBytes,
	}, nil
}

// getLatency 获取延迟统计
func (s *StatisticsService) getLatency(query *gorm.DB) (*LatencyStatistics, error) {
	type Result struct {
		TotalMs int64
		AvgMs   float64
		MinMs   int64
		MaxMs   int64
	}

	var result Result
	err := query.Select(
		"COALESCE(SUM(latency_ms), 0) as total_ms",
		"COALESCE(AVG(latency_ms), 0) as avg_ms",
		"COALESCE(MIN(latency_ms), 0) as min_ms",
		"COALESCE(MAX(latency_ms), 0) as max_ms",
	).Scan(&result).Error

	if err != nil {
		return nil, err
	}

	return &LatencyStatistics{
		TotalMs: result.TotalMs,
		AvgMs:   result.AvgMs,
		MinMs:   result.MinMs,
		MaxMs:   result.MaxMs,
	}, nil
}

// groupByIP 按IP分组统计
func (s *StatisticsService) groupByIP(query *gorm.DB) ([]IPStatistics, error) {
	type Result struct {
		IP    string
		Count int64
	}

	var results []Result
	err := query.Select("x_forwarded_for as ip, COUNT(*) as count").
		Where("x_forwarded_for != ''").
		Group("x_forwarded_for").
		Order("count DESC").
		Scan(&results).Error

	if err != nil {
		return nil, err
	}

	var stats []IPStatistics
	for _, r := range results {
		stats = append(stats, IPStatistics{
			IP:    r.IP,
			Count: r.Count,
		})
	}

	return stats, nil
}

// groupByPath 按路径分组统计
func (s *StatisticsService) groupByPath(query *gorm.DB) ([]PathStatistics, error) {
	type Result struct {
		Path  string
		Count int64
	}

	var results []Result
	err := query.Select("path, COUNT(*) as count").
		Group("path").
		Order("count DESC").
		Scan(&results).Error

	if err != nil {
		return nil, err
	}

	var stats []PathStatistics
	for _, r := range results {
		stats = append(stats, PathStatistics{
			Path:  r.Path,
			Count: r.Count,
		})
	}

	return stats, nil
}

// groupByDate 按日期分组统计
func (s *StatisticsService) groupByDate(query *gorm.DB) ([]DateStatistics, error) {
	type Result struct {
		Date  string
		Count int64
	}

	var results []Result
	err := query.Select("date(time) as date, COUNT(*) as count").
		Group("date(time)").
		Order("date(time) ASC").
		Scan(&results).Error

	if err != nil {
		return nil, err
	}

	var stats []DateStatistics
	for _, r := range results {
		stats = append(stats, DateStatistics{
			Date:  r.Date,
			Count: r.Count,
		})
	}

	return stats, nil
}

// groupByTime 按小时分组统计
func (s *StatisticsService) groupByTime(query *gorm.DB) ([]TimeStatistics, error) {
	type Result struct {
		Time  string
		Count int64
	}

	var results []Result
	err := query.Select("strftime('%Y-%m-%d %H:00:00', time) as time, COUNT(*) as count").
		Group("strftime('%Y-%m-%d %H:00:00', time)").
		Order("time ASC").
		Scan(&results).Error

	if err != nil {
		return nil, err
	}

	var stats []TimeStatistics
	for _, r := range results {
		stats = append(stats, TimeStatistics{
			Time:  r.Time,
			Count: r.Count,
		})
	}

	return stats, nil
}

// GetAuthorizationRanking 获取 authorization 使用次数排行
func (s *StatisticsService) GetAuthorizationRanking(req *GetAuthorizationRankingRequest) (*AuthorizationRankingResponse, error) {
	// 解析时间范围，默认最近7天
	startTime, endTime, err := s.parseTimeRange(req.StartTime, req.EndTime)
	if err != nil {
		return nil, err
	}

	// 分页参数
	page := req.Page
	if page < 1 {
		page = 1
	}
	pageSize := req.PageSize
	if pageSize < 1 {
		pageSize = 20
	}
	if pageSize > 100 {
		pageSize = 100
	}

	// 构建查询
	baseQuery := database.DB.Model(&models.TokenUsageLog{}).
		Where("time >= ?", startTime).
		Where("time <= ?", endTime).
		Where("authorization != ''")

	// 获取总数（不同 authorization 的数量）
	var total int64
	if err := baseQuery.Distinct("authorization").Count(&total).Error; err != nil {
		return nil, errors.New("获取排行榜总数失败")
	}

	// 获取分页数据
	type Result struct {
		Authorization string
		Count         int64
	}

	var results []Result
	offset := (page - 1) * pageSize
	err = baseQuery.Select("authorization, COUNT(*) as count").
		Group("authorization").
		Order("count DESC").
		Offset(offset).
		Limit(pageSize).
		Scan(&results).Error

	if err != nil {
		return nil, errors.New("获取排行榜数据失败")
	}

	// 转换结果
	list := make([]AuthorizationRankingItem, 0, len(results))
	for _, r := range results {
		list = append(list, AuthorizationRankingItem{
			Authorization: r.Authorization,
			Count:         r.Count,
		})
	}

	return &AuthorizationRankingResponse{
		Total: total,
		TimeRange: TimeRange{
			Start: startTime.Format("2006-01-02 15:04:05"),
			End:   endTime.Format("2006-01-02 15:04:05"),
		},
		List: list,
	}, nil
}
