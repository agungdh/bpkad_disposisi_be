package auth

import (
	"net/http"

	"agungdh.com/bpkad_disposisi_be/models"
	"github.com/gin-gonic/gin"
)

type login struct {
	Username string
	Password string
}

func Login(c *gin.Context) {
	var login login

	if err := c.ShouldBindJSON(&login); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var user models.User

	result := models.Db.Where("username = ?", login.Username).First(&user)

	if result.RowsAffected > 0 && models.CheckPasswordHash(login.Password, user.Password) {
		c.JSON(http.StatusOK, user)
	} else {
		c.JSON(http.StatusNotFound, nil)
	}
}
