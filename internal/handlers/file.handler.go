package handlers

import (
	"errors"
	"net/http"
	"tlms/internal/services"
	"tlms/internal/storage"
	"tlms/internal/validators"

	"github.com/gin-gonic/gin"
)

type FileHandler struct {
	fileService services.FileService
}

func NewFileHandler(fileService services.FileService) *FileHandler {
	return &FileHandler{fileService}
}

func (h *FileHandler) Upload(c *gin.Context) {
	folder := c.PostForm("folder")
	if folder == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "folder is required"})
		c.Abort()
		return
	}

	fh, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "file is required"})
		c.Abort()
		return
	}

	result, err := h.fileService.Upload(c.Request.Context(), folder, fh)
	if err != nil {
		h.handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func (h *FileHandler) GetMetadata(c *gin.Context) {
	fileUUID := c.Param("uuid")

	result, err := h.fileService.GetMetadata(fileUUID)
	if err != nil {
		h.handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *FileHandler) Download(c *gin.Context) {
	fileUUID := c.Param("uuid")

	reader, meta, err := h.fileService.OpenForDownload(c.Request.Context(), fileUUID)
	if err != nil {
		h.handleServiceError(c, err)
		return
	}
	if closer, ok := reader.(interface{ Close() error }); ok {
		defer closer.Close()
	}

	c.Header("Content-Disposition", "attachment; filename=\""+meta.OriginalName+"\"")
	c.Header("Content-Type", meta.MimeType)
	c.DataFromReader(http.StatusOK, meta.Size, meta.MimeType, reader, nil)
}

func (h *FileHandler) Delete(c *gin.Context) {
	fileUUID := c.Param("uuid")

	if err := h.fileService.Delete(c.Request.Context(), fileUUID); err != nil {
		h.handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "file deleted"})
}

// handleServiceError memetakan sentinel error dari Service/Validator/Storage
// layer ke HTTP status yang sesuai — mengikuti pola errors.Is() di sloc.handler.go.
func (h *FileHandler) handleServiceError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, services.ErrFileMetaNotFound), errors.Is(err, storage.ErrFileNotFound):
		c.JSON(http.StatusNotFound, gin.H{"message": "file not found"})
	case errors.Is(err, validators.ErrUnsupportedMimeType),
		errors.Is(err, validators.ErrFileTooLarge),
		errors.Is(err, validators.ErrEmptyFile),
		errors.Is(err, validators.ErrMimeMismatch):
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.Abort()
}
