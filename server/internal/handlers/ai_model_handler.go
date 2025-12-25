// Package handlers HTTP请求处理器
// 处理 AI 模型管理相关的 HTTP 请求
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

// NewAIModelHandler 创建 AI 模型处理器实例
func NewAIModelHandler() *AIModelHandler {
	return &AIModelHandler{
		aiModelService: services.NewAIModelService(),
	}
}

// CreateAIModel 创建 AI 模型
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

// GetAIModel 获取 AI 模型详情
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

// ListAIModels 获取 AI 模型列表
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

// UpdateAIModel 更新 AI 模型
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

// DeleteAIModel 删除 AI 模型
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
