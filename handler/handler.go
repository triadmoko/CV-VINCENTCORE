package handler

import (
	"vincentcore_interview/service"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	Register(c *gin.Context)
}
type handler struct {
	service service.Service
}

func NewHandler(service service.Service) *handler {
	return &handler{service: service}
}
