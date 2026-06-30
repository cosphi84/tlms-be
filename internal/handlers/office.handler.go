package handlers

import (
	"net/http"
	"strconv"
	"tlms/internal/dto"
	"tlms/internal/services"

	"github.com/gin-gonic/gin"
)

type OfficeHandler struct {
	officeService services.OfficeService
}

func NewOfficeHandler(officeService services.OfficeService) *OfficeHandler {
	return &OfficeHandler{officeService: officeService}
}

// Create godoc
// @Summary      Create a new office
// @Tags         offices
// @Accept       json
// @Produce      json
// @Param        body body dto.CreateOfficeRequest true "Office payload"
// @Success      201  {object}  map[string]string
// @Failure      400  {object}  map[string]string
// @Router       /offices [post]
func (h *OfficeHandler) Create(c *gin.Context) {
	var req dto.CreateOfficeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.officeService.CreateOffice(req, c.Request.Context()); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "office created successfully"})
}

// FindAll godoc
// @Summary      List all offices with pagination
// @Tags         offices
// @Produce      json
// @Param        page      query  int    false  "Page number (default: 1)"
// @Param        limit     query  int    false  "Items per page (default: 10)"
// @Param        sorted_by query  string false  "Column to sort by (default: created_at)"
// @Param        sort_dir  query  string false  "asc or desc (default: desc)"
// @Success      200  {object}  dto.PaginationResponse
// @Router       /offices [get]
func (h *OfficeHandler) FindAll(c *gin.Context) {
	pagination := parsePaginationQuery(c)

	result, err := h.officeService.GetOffices(pagination)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

// FindOptions godoc
// @Summary      Get offices as label/value pairs for Select components
// @Tags         offices
// @Produce      json
// @Success      200  {array}  dto.OfficeOptionResponse
// @Router       /offices/options [get]
func (h *OfficeHandler) FindOptions(c *gin.Context) {
	opts, err := h.officeService.GetOfficeOptions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, opts)
}

// Update godoc
// @Summary      Update an existing office (partial update supported)
// @Tags         offices
// @Accept       json
// @Produce      json
// @Param        id   path  int  true  "Office ID"
// @Param        body body  dto.UpdateOfficeRequest true "Fields to update"
// @Success      200  {object}  map[string]string
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /offices/{id} [put]
func (h *OfficeHandler) Update(c *gin.Context) {
	id, err := parseIDParam(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid office id"})
		return
	}

	var req dto.UpdateOfficeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.officeService.UpdateOffice(id, req, c.Request.Context()); err != nil {
		status := http.StatusBadRequest
		if err.Error() == "office not found" {
			status = http.StatusNotFound
		}
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "office updated successfully"})
}

// Delete godoc
// @Summary      Soft-delete an office
// @Tags         offices
// @Produce      json
// @Param        id   path  int  true  "Office ID"
// @Success      200  {object}  map[string]string
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /offices/{id} [delete]
func (h *OfficeHandler) Delete(c *gin.Context) {
	id, err := parseIDParam(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid office id"})
		return
	}

	if err := h.officeService.DeleteOffice(id, c.Request.Context()); err != nil {
		status := http.StatusBadRequest
		if err.Error() == "office not found" {
			status = http.StatusNotFound
		}
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "office deleted successfully"})
}

// --- helpers ---

func parsePaginationQuery(c *gin.Context) *dto.PaginationRequest {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	return &dto.PaginationRequest{
		Page:     page,
		Limit:    limit,
		SortedBy: c.DefaultQuery("sorted_by", "created_at"),
		SortDir:  c.DefaultQuery("sort_dir", "desc"),
	}
}

func parseIDParam(c *gin.Context) (int64, error) {
	return strconv.ParseInt(c.Param("id"), 10, 64)
}
