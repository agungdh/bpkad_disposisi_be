package main

import (
	"agungdh.com/bpkad_disposisi_be/models"
	"agungdh.com/bpkad_disposisi_be/services/bidangservice"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	r := gin.Default()

	models.Setup()

	r.GET("/api/bidang", bidangservice.Index)
	r.GET("/api/bidang/:id", bidangservice.Show)
	// r.POST("/api/bidang", bidangservice.Create)
	// r.PUT("/api/bidang/:id", bidangservice.Update)
	// r.DELETE("/api/bidang", bidangservice.Delete)

	r.Run()
}
