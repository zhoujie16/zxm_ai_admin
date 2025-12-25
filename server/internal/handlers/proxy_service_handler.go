// Package handlers HTTP请求处理器
// 处理代理服务管理相关的HTTP请求，包括增删改查操作
package handlers

import (
	"strconv"
	"zxm_ai_admin/server/internal/services"
	"zxm_ai_admin/server/internal/utils"

	"github.com/gin-gonic/gin"
)

type ProxyServiceHandler struct {
	proxyServiceService *services.ProxyServiceService
}

// NewProxyServiceHandler 创建代理服务处理器实例
func NewProxyServiceHandler() *ProxyServiceHandler {
	return &ProxyServiceHandler{
		proxyServiceService: services.NewProxyServiceService(),
	}
}

// CreateProxyService 创建代理服务
// @Summary 创建代理服务
// @Description 创建新的代理服务
// @Tags 代理服务
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param body body services.CreateProxyServiceRequest true "代理服务信息"
// @Success 200 {object} utils.Response
// @Router /api/proxy-services [post]
func (h *ProxyServiceHandler) CreateProxyService(c *gin.Context) {
	var req services.CreateProxyServiceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, "参数错误: "+err.Error())
		return
	}

	proxyService, err := h.proxyServiceService.CreateProxyService(&req)
	if err != nil {
		utils.Error(c, 400, err.Error())
		return
	}

	utils.Success(c, proxyService)
}

// GetProxyService 获取代理服务详情
// @Summary 获取代理服务详情
// @Description 根据ID获取代理服务详情
// @Tags 代理服务
// @Security BearerAuth
// @Produce json
// @Param id path int true "代理服务ID"
// @Success 200 {object} utils.Response
// @Router /api/proxy-services/{id} [get]
func (h *ProxyServiceHandler) GetProxyService(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.BadRequest(c, "无效的ID")
		return
	}

	proxyService, err := h.proxyServiceService.GetProxyService(uint(id))
	if err != nil {
		utils.NotFound(c, err.Error())
		return
	}

	utils.Success(c, proxyService)
}

// ListProxyServices 获取代理服务列表
// @Summary 获取代理服务列表
// @Description 分页获取代理服务列表
// @Tags 代理服务
// @Security BearerAuth
// @Produce json
// @Param page query int false "页码，从1开始" default(1)
// @Param page_size query int false "每页数量" default(10)
// @Success 200 {object} utils.Response
// @Router /api/proxy-services [get]
func (h *ProxyServiceHandler) ListProxyServices(c *gin.Context) {
	var req services.ListProxyServicesRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		utils.BadRequest(c, "参数错误: "+err.Error())
		return
	}

	response, err := h.proxyServiceService.ListProxyServices(&req)
	if err != nil {
		utils.InternalServerError(c, err.Error())
		return
	}

	utils.Success(c, response)
}

// UpdateProxyService 更新代理服务
// @Summary 更新代理服务
// @Description 更新代理服务信息
// @Tags 代理服务
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "代理服务ID"
// @Param body body services.UpdateProxyServiceRequest true "代理服务信息"
// @Success 200 {object} utils.Response
// @Router /api/proxy-services/{id} [put]
func (h *ProxyServiceHandler) UpdateProxyService(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.BadRequest(c, "无效的ID")
		return
	}

	var req services.UpdateProxyServiceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, "参数错误: "+err.Error())
		return
	}

	proxyService, err := h.proxyServiceService.UpdateProxyService(uint(id), &req)
	if err != nil {
		utils.Error(c, 400, err.Error())
		return
	}

	utils.Success(c, proxyService)
}

// DeleteProxyService 删除代理服务
// @Summary 删除代理服务
// @Description 根据ID删除代理服务（软删除）
// @Tags 代理服务
// @Security BearerAuth
// @Produce json
// @Param id path int true "代理服务ID"
// @Success 200 {object} utils.Response
// @Router /api/proxy-services/{id} [delete]
func (h *ProxyServiceHandler) DeleteProxyService(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.BadRequest(c, "无效的ID")
		return
	}

	if err := h.proxyServiceService.DeleteProxyService(uint(id)); err != nil {
		utils.Error(c, 400, err.Error())
		return
	}

	utils.Success(c, nil)
}


