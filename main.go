package main

import (
	"vincentcore_interview/config"
	"vincentcore_interview/handler"
	"vincentcore_interview/pkg/logs"
	"vincentcore_interview/repository"
	"vincentcore_interview/service"

	"github.com/gin-gonic/gin"
)

func main() {
	logging := logs.NewLogger()
	db, err := config.Mysql()
	if err != nil {
		logging.Error(err)
		return
	}
	handler := handler.NewHandler(service.NewService(repository.NewRepository(db), logging))

	r := gin.Default()
	auth := r.Group("/auth/v1")
	auth.POST("/register", handler.Register)
	auth.POST("/login", handler.Login)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
