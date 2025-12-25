// Package services 业务逻辑服务层
// 实现代理服务管理相关的业务逻辑，包括增删改查操作
package services

import (
	"errors"
	"net"
	"zxm_ai_admin/server/internal/database"
	"zxm_ai_admin/server/internal/models"
)

type ProxyServiceService struct{}

// NewProxyServiceService 创建代理服务业务逻辑实例
func NewProxyServiceService() *ProxyServiceService {
	return &ProxyServiceService{}
}

// CreateProxyServiceRequest 创建代理服务请求
type CreateProxyServiceRequest struct {
	ServiceID string `json:"service_id" binding:"required"` // 服务标识
	ServerIP  string `json:"server_ip" binding:"required"`  // 服务器IP
	Status    int    `json:"status"`                        // 状态：1=启用，0=未启用
	Remark    string `json:"remark"`                        // 备注
}

// UpdateProxyServiceRequest 更新代理服务请求
type UpdateProxyServiceRequest struct {
	ServiceID string  `json:"service_id"` // 服务标识
	ServerIP  string  `json:"server_ip"`  // 服务器IP
	Status    *int    `json:"status"`      // 状态：1=启用，0=未启用
	Remark    *string `json:"remark"`      // 备注
}

// ListProxyServicesRequest 列表查询请求
type ListProxyServicesRequest struct {
	Page     int `form:"page"`      // 页码，从1开始
	PageSize int `form:"page_size"` // 每页数量
}

// ListProxyServicesResponse 列表查询响应
type ListProxyServicesResponse struct {
	Total int64              `json:"total"` // 总数量
	List  []models.ProxyService `json:"list"`  // 列表数据
}

// CreateProxyService 创建代理服务
func (s *ProxyServiceService) CreateProxyService(req *CreateProxyServiceRequest) (*models.ProxyService, error) {
	// 验证IP格式
	if err := validateIP(req.ServerIP); err != nil {
		return nil, err
	}

	// 检查服务标识是否已存在
	var count int64
	if err := database.DB.Model(&models.ProxyService{}).
		Where("service_id = ?", req.ServiceID).
		Count(&count).Error; err != nil {
		return nil, errors.New("检查服务标识失败")
	}
	if count > 0 {
		return nil, errors.New("服务标识已存在")
	}

	// 设置默认状态
	status := req.Status
	if status != 0 && status != 1 {
		status = 1 // 默认启用
	}

	// 创建代理服务
	proxyService := &models.ProxyService{
		ServiceID: req.ServiceID,
		ServerIP:  req.ServerIP,
		Status:    status,
		Remark:    req.Remark,
	}

	if err := database.DB.Create(proxyService).Error; err != nil {
		return nil, errors.New("创建代理服务失败")
	}

	return proxyService, nil
}

// GetProxyService 根据ID获取代理服务
func (s *ProxyServiceService) GetProxyService(id uint) (*models.ProxyService, error) {
	var proxyService models.ProxyService
	if err := database.DB.First(&proxyService, id).Error; err != nil {
		return nil, errors.New("代理服务不存在")
	}
	return &proxyService, nil
}

// ListProxyServices 获取代理服务列表
func (s *ProxyServiceService) ListProxyServices(req *ListProxyServicesRequest) (*ListProxyServicesResponse, error) {
	// 设置默认分页参数
	page := req.Page
	if page < 1 {
		page = 1
	}
	pageSize := req.PageSize
	if pageSize < 1 {
		pageSize = 10
	}
	if pageSize > 100 {
		pageSize = 100
	}

	var total int64
	var list []models.ProxyService

	// 查询总数
	if err := database.DB.Model(&models.ProxyService{}).Count(&total).Error; err != nil {
		return nil, errors.New("查询代理服务列表失败")
	}

	// 查询列表
	offset := (page - 1) * pageSize
	if err := database.DB.
		Order("created_at DESC").
		Offset(offset).
		Limit(pageSize).
		Find(&list).Error; err != nil {
		return nil, errors.New("查询代理服务列表失败")
	}

	return &ListProxyServicesResponse{
		Total: total,
		List:  list,
	}, nil
}

// UpdateProxyService 更新代理服务
func (s *ProxyServiceService) UpdateProxyService(id uint, req *UpdateProxyServiceRequest) (*models.ProxyService, error) {
	// 查询代理服务是否存在
	var proxyService models.ProxyService
	if err := database.DB.First(&proxyService, id).Error; err != nil {
		return nil, errors.New("代理服务不存在")
	}

	// 如果更新服务标识，检查是否重复
	if req.ServiceID != "" && req.ServiceID != proxyService.ServiceID {
		var count int64
		if err := database.DB.Model(&models.ProxyService{}).
			Where("service_id = ? AND id != ?", req.ServiceID, id).
			Count(&count).Error; err != nil {
			return nil, errors.New("检查服务标识失败")
		}
		if count > 0 {
			return nil, errors.New("服务标识已存在")
		}
		proxyService.ServiceID = req.ServiceID
	}

	// 如果更新IP，验证IP格式
	if req.ServerIP != "" {
		if err := validateIP(req.ServerIP); err != nil {
			return nil, err
		}
		proxyService.ServerIP = req.ServerIP
	}

	// 更新状态
	if req.Status != nil {
		if *req.Status != 0 && *req.Status != 1 {
			return nil, errors.New("状态值无效，只能为0或1")
		}
		proxyService.Status = *req.Status
	}

	// 更新备注
	if req.Remark != nil {
		proxyService.Remark = *req.Remark
	}

	// 保存更新
	if err := database.DB.Save(&proxyService).Error; err != nil {
		return nil, errors.New("更新代理服务失败")
	}

	return &proxyService, nil
}

// DeleteProxyService 删除代理服务
func (s *ProxyServiceService) DeleteProxyService(id uint) error {
	// 检查是否存在
	var proxyService models.ProxyService
	if err := database.DB.First(&proxyService, id).Error; err != nil {
		return errors.New("代理服务不存在")
	}

	// 软删除
	if err := database.DB.Delete(&proxyService).Error; err != nil {
		return errors.New("删除代理服务失败")
	}

	return nil
}

// validateIP 验证IP地址格式
func validateIP(ip string) error {
	if ip == "" {
		return errors.New("IP地址不能为空")
	}
	if net.ParseIP(ip) == nil {
		return errors.New("IP地址格式无效")
	}
	return nil
}

