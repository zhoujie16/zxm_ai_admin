// Package handlers HTTP请求处理器
// 处理模型代理管理相关的 HTTP 请求
package handlers

import (
	"net/http"
	"strconv"
	"zxm_ai_admin/server/internal/services"
	"zxm_ai_admin/server/internal/utils"

	"github.com/gin-gonic/gin"
)

type AIModelHandler struct {
	aiModelService *services.AIModelService
}

// NewAIModelHandler 创建模型代理处理器实例
func NewAIModelHandler() *AIModelHandler {
	return &AIModelHandler{
		aiModelService: services.NewAIModelService(),
	}
}

// CreateAIModel 创建模型代理
func (h *AIModelHandler) CreateAIModel(c *gin.Context) {
	var req services.CreateAIModelRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, "参数错误: "+err.Error())
		return
	}

	aiModel, err := h.aiModelService.CreateAIModel(&req)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.Success(c, aiModel)
}

// GetAIModel 获取模型代理详情
func (h *AIModelHandler) GetAIModel(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.BadRequest(c, "无效的ID")
		return
	}

	aiModel, err := h.aiModelService.GetAIModel(uint(id))
	if err != nil {
		utils.NotFound(c, err.Error())
		return
	}

	utils.Success(c, aiModel)
}

// ListAIModels 获取模型代理列表
func (h *AIModelHandler) ListAIModels(c *gin.Context) {
	var req services.ListAIModelsRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		utils.BadRequest(c, "参数错误: "+err.Error())
		return
	}

	response, err := h.aiModelService.ListAIModels(&req)
	if err != nil {
		utils.InternalServerError(c, err.Error())
		return
	}

	utils.Success(c, response)
}

// UpdateAIModel 更新模型代理
func (h *AIModelHandler) UpdateAIModel(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.BadRequest(c, "无效的ID")
		return
	}

	var req services.UpdateAIModelRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, "参数错误: "+err.Error())
		return
	}

	aiModel, err := h.aiModelService.UpdateAIModel(uint(id), &req)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.Success(c, aiModel)
}

// DeleteAIModel 删除模型代理
func (h *AIModelHandler) DeleteAIModel(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.BadRequest(c, "无效的ID")
		return
	}

	if err := h.aiModelService.DeleteAIModel(uint(id)); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.Success(c, nil)
}
