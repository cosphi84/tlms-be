package handlers

import (
	"net/http"
	"tlms/internal/dto"
	"tlms/internal/helpers"
	"tlms/internal/services"

	"github.com/gin-gonic/gin"
)

type ToolsHandler struct {
	toolService services.ToolsService
}

func NewToolHandler(toolService services.ToolsService) *ToolsHandler {
	return &ToolsHandler{toolService}
}

func (h *ToolsHandler) GetAllTools(c *gin.Context) {
	pag := helpers.ParsePaginationQuery(c)

	res, err := h.toolService.FindAll(pag)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *ToolsHandler) GetTool(c *gin.Context) {
	id, err := helpers.ParseIDParam(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	tool, err := h.toolService.FindById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, tool)
}

func (h *ToolsHandler) RegisterTool(c *gin.Context) {
	var req dto.RegisterToolRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	if err := h.toolService.Create(&req, c.Request.Context()); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "tool created"})
}

func (h *ToolsHandler) UpdateTool(c *gin.Context) {
	id, err := helpers.ParseIDParam(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	var req dto.RegisterToolRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	if err = h.toolService.Update(id, &req, c.Request.Context()); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "tool updated"})
}

func (h *ToolsHandler) DeleteTool(c *gin.Context) {
	id, err := helpers.ParseIDParam(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	if err := h.toolService.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "tool deleted"})
}
