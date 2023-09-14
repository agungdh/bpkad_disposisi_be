package main

import (
	"io"
	"os"

	"agungdh.com/bpkad_disposisi_be/models"
	"agungdh.com/bpkad_disposisi_be/services/auth"
	"agungdh.com/bpkad_disposisi_be/services/bidang"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	models.Setup()

	r := gin.Default()

	// Logging to a file.
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	r.POST("/api/login", auth.Login)

	r.GET("/api/bidang", bidang.Index)
	r.GET("/api/bidang/:id", bidang.Show)
	r.POST("/api/bidang", bidang.Create)
	r.PUT("/api/bidang/:id", bidang.Update)
	r.DELETE("/api/bidang/:id", bidang.Delete)

	r.Run()
}
