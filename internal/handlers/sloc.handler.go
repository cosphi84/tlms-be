package handlers

import (
	"errors"
	"net/http"
	"tlms/internal/dto"
	"tlms/internal/helpers"
	"tlms/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type StorageLocationHandler struct {
	slocService services.StogareLocationService
}

func NewStorageLocationHandler(slocService services.StogareLocationService) *StorageLocationHandler {
	return &StorageLocationHandler{
		slocService: slocService,
	}
}

func (h *StorageLocationHandler) CreateSloc(c *gin.Context) {
	var req dto.StorageLocationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	if err := h.slocService.Create(req, c.Request.Context()); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "storage location created"})
}

func (h *StorageLocationHandler) GetSLoc(c *gin.Context) {
	id, err := helpers.GetIDFromParam(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		c.Abort()
		return
	}

	sloc, err := h.slocService.FindById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"message": "storage location not found"})
			c.Abort()
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return

	}

	c.JSON(http.StatusOK, sloc)
}

func (h *StorageLocationHandler) GetAllSloc(c *gin.Context) {
	pagination := helpers.ParsePaginationQuery(c)

	result, err := h.slocService.GetAllSloc(pagination)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, result)
}

func (h *StorageLocationHandler) DeleteSLoc(c *gin.Context) {
	id, err := helpers.GetIDFromParam(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		c.Abort()
		return
	}
	if err := h.slocService.Delete(id, c.Request.Context()); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"message": "storage location not found"})
			c.Abort()
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "storage location deleted"})
}

func (h *StorageLocationHandler) UpdateSLoc(c *gin.Context) {
	id, err := helpers.GetIDFromParam(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		c.Abort()
		return
	}
	var req dto.StorageLocationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	if err := h.slocService.Update(id, req, c.Request.Context()); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "storage location updated"})
}
