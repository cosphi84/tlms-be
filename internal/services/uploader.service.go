package services

import (
	"context"

	"github.com/gin-gonic/gin"
)

type UploaderService interface {
	UploadHandler(req, ctx *gin.Context)
}

type uploaderService struct{}

func NewUploaderService() UploaderService {
	return &uploaderService{}
}

func (srv *UploaderService) UploadHandler(req, ctx *context.Context)
