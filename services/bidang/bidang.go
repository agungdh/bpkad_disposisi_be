package bidang

import (
	"errors"
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
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, bidang)
}

func Create(c *gin.Context) {
	var bidang models.Bidang

	if err := c.ShouldBindJSON(&bidang); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	validationErrors := ValidateStoreUpdate(bidang)
	if len(validationErrors) > 0 {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, validationErrors)
		return
	}

	models.Db.Create(&bidang)

	c.JSON(http.StatusOK, bidang)
}

func Update(c *gin.Context) {
	models.Db.Transaction(func(tx *gorm.DB) error {
		var bidang models.Bidang
		var bidangModel models.Bidang

		if err := c.ShouldBindJSON(&bidang); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return err
		}

		validationErrors := ValidateStoreUpdate(bidang)
		if len(validationErrors) > 0 {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, validationErrors)
			return errors.New("validation error")
		}

		id := c.Param("id")

		if err := tx.Where("id = ?", id).First(&bidangModel).Error; err != nil {
			switch err {
			case gorm.ErrRecordNotFound:
				c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": err.Error()})
				return err
			default:
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
				return err
			}
		}

		if err := tx.Model(&bidangModel).Updates(&bidang).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return err
		}

		c.JSON(http.StatusOK, bidangModel)

		return nil
	})

}

func Delete(c *gin.Context) {
	models.Db.Transaction(func(tx *gorm.DB) error {
		var bidang models.Bidang
		id := c.Param("id")

		if err := tx.Where("id = ?", id).First(&bidang).Error; err != nil {
			switch err {
			case gorm.ErrRecordNotFound:
				c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": err.Error()})
				return err
			default:
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
				return err
			}
		}

		if err := tx.Delete(&bidang).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return err
		}

		c.JSON(http.StatusOK, bidang)

		return nil
	})
}
