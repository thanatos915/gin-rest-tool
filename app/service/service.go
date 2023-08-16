package service

import (
	"gin-rest-tool/app/dto"
	"gin-rest-tool/global"
	"github.com/gin-gonic/gin"
)

type Service struct {
	dto *dto.Dto
	c   *gin.Context
}

func New(c *gin.Context) Service {
	svc := Service{c: c}
	svc.dto = dto.New(global.DB)
	return svc
}
