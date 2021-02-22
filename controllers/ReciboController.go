// controllers/ReciboController.go

package controllers

import (
	"billetera/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FindRecibos(c *gin.Context) {
	var Recibo []models.Recibo
	models.DB.Find(&Recibo)

	c.JSON(http.StatusOK, Recibo)
}

func CreateRecibo(c *gin.Context) {
	// Validate input
	var input models.CreateReciboInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errorACA": err.Error()})
		return
	}

	// Create recibo
	recibo := models.Recibo{
		IdTablaApi:     input.IdTablaApi,
		NombreTablaApi: input.NombreTablaApi,
	}
	models.DB.Create(&recibo)

	c.JSON(http.StatusOK, recibo)
}
