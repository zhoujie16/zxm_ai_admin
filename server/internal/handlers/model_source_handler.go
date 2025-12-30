// Package handlers HTTP请求处理器
// 处理模型来源管理相关的 HTTP 请求
package handlers

import (
	"net/http"
	"strconv"
	"zxm_ai_admin/server/internal/services"
	"zxm_ai_admin/server/internal/utils"

	"github.com/gin-gonic/gin"
)

type ModelSourceHandler struct {
	modelSourceService *services.ModelSourceService
}

// NewModelSourceHandler 创建模型来源处理器实例
func NewModelSourceHandler() *ModelSourceHandler {
	return &ModelSourceHandler{
		modelSourceService: services.NewModelSourceService(),
	}
}

// CreateModelSource 创建模型来源
func (h *ModelSourceHandler) CreateModelSource(c *gin.Context) {
	var req services.CreateModelSourceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, "参数错误: "+err.Error())
		return
	}

	modelSource, err := h.modelSourceService.CreateModelSource(&req)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.Success(c, modelSource)
}

// GetModelSource 获取模型来源详情
func (h *ModelSourceHandler) GetModelSource(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.BadRequest(c, "无效的ID")
		return
	}

	modelSource, err := h.modelSourceService.GetModelSource(uint(id))
	if err != nil {
		utils.NotFound(c, err.Error())
		return
	}

	utils.Success(c, modelSource)
}

// ListModelSources 获取模型来源列表
func (h *ModelSourceHandler) ListModelSources(c *gin.Context) {
	var req services.ListModelSourcesRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		utils.BadRequest(c, "参数错误: "+err.Error())
		return
	}

	response, err := h.modelSourceService.ListModelSources(&req)
	if err != nil {
		utils.InternalServerError(c, err.Error())
		return
	}

	utils.Success(c, response)
}

// UpdateModelSource 更新模型来源
func (h *ModelSourceHandler) UpdateModelSource(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.BadRequest(c, "无效的ID")
		return
	}

	var req services.UpdateModelSourceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, "参数错误: "+err.Error())
		return
	}

	modelSource, err := h.modelSourceService.UpdateModelSource(uint(id), &req)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.Success(c, modelSource)
}

// DeleteModelSource 删除模型来源
func (h *ModelSourceHandler) DeleteModelSource(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.BadRequest(c, "无效的ID")
		return
	}

	if err := h.modelSourceService.DeleteModelSource(uint(id)); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.Success(c, nil)
}
