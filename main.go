package main

import (
	"io"
	"os"

	"agungdh.com/bpkad_disposisi_be/models"
	"agungdh.com/bpkad_disposisi_be/services/bidang"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	r := gin.Default()

	// Logging to a file.
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	models.Setup()

	r.GET("/api/bidang", bidang.Index)
	r.GET("/api/bidang/:id", bidang.Show)
	r.POST("/api/bidang", bidang.Create)
	r.PUT("/api/bidang/:id", bidang.Update)
	r.DELETE("/api/bidang/:id", bidang.Delete)

	r.Run()
}
