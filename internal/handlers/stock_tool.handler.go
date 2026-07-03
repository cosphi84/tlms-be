package handlers

import (
	"net/http"
	"tlms/internal/dto"
	"tlms/internal/models"
	"tlms/internal/services"

	"github.com/gin-gonic/gin"
)

type StockToolHandler struct {
	stockToolService services.StockToolService
}

func NewStockToolHandler(stockToolService services.StockToolService) *StockToolHandler {
	return &StockToolHandler{
		stockToolService: stockToolService,
	}
}

func (h *StockToolHandler) AddStockTool(c *gin.Context) {
	req := dto.ModifyStokToolRequest{Mode: string(models.ModifyStockToolsTypeIncoming)}
	h.execHandler(c, req)
}

func (h *StockToolHandler) RemoveStockTool(c *gin.Context) {
	req := dto.ModifyStokToolRequest{Mode: string(models.ModifyStockToolsTypeOutgoing)}
	h.execHandler(c, req)
}

func (h *StockToolHandler) execHandler(c *gin.Context, req dto.ModifyStokToolRequest) {
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	if err := h.stockToolService.ModifyStock(&req, c.Request.Context()); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "stock tool saved"})
}
