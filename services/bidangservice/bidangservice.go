package bidangservice

import (
	"net/http"

	"agungdh.com/bpkad_disposisi_be/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var bidangs []models.Bidang

	models.Db.Find(&bidangs)

	c.JSON(http.StatusOK, bidangs)
}

func Show(c *gin.Context) {
	var bidang models.Bidang
	id := c.Param("id")

	if err := models.Db.Where("id = ?", id).First(&bidang).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, bidang)
}

// func Create(c *gin.Context) {

// 	var product models.Product

// 	if err := c.ShouldBindJSON(&product); err != nil {
// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
// 		return
// 	}

// 	models.DB.Create(&product)
// 	c.JSON(http.StatusOK, gin.H{"product": product})
// }

// func Update(c *gin.Context) {
// 	var product models.Product
// 	id := c.Param("id")

// 	if err := c.ShouldBindJSON(&product); err != nil {
// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
// 		return
// 	}

// 	if models.DB.Model(&product).Where("id = ?", id).Updates(&product).RowsAffected == 0 {
// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "tidak dapat mengupdate product"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil diperbarui"})

// }

// func Delete(c *gin.Context) {

// 	var product models.Product

// 	var input struct {
// 		Id json.Number
// 	}

// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
// 		return
// 	}

// 	id, _ := input.Id.Int64()
// 	if models.DB.Delete(&product, id).RowsAffected == 0 {
// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat menghapus product"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
// }
