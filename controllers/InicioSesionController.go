package controllers

import (
	"billetera/models"
	"crypto/sha256"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PassHash(pas string) string {
	h := sha256.New()
	h.Write([]byte(pas))
	return string(h.Sum(nil))
}

func InicioSesion(c *gin.Context) {
	var usuario models.Usuario

	if err := models.DB.Where("dni = ? AND contraseña = ?", c.Param("dni"), PassHash(c.Param("pass"))).First(&usuario).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No Encontrado!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": true})
}
