package helpers

import (
	"strconv"
	"tlms/internal/dto"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetIDFromParam(c *gin.Context) (int64, error) {
	return strconv.ParseInt(c.Param("id"), 10, 64)
}

func GetUuidFromParam(c *gin.Context) (uuid.UUID, error) {
	return uuid.Parse(c.Param("id"))
}

func ParsePaginationQuery(c *gin.Context) *dto.PaginationRequest {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	return &dto.PaginationRequest{
		Page:     page,
		Limit:    limit,
		Search:   c.Query("search"),
		SortedBy: c.DefaultQuery("sorted_by", "created_at"),
		SortDir:  c.DefaultQuery("sort_dir", "desc"),
	}
}

func ParseIDParam(c *gin.Context) (int64, error) {
	return strconv.ParseInt(c.Param("id"), 10, 64)
}
