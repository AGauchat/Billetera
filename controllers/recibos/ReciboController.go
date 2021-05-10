// controllers/ReciboController.go

package recibos

import (
	"billetera/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GET /recibos
func FindRecibos(c *gin.Context) {
	var Recibo []models.Recibo
	var db = models.DB
	db.Find(&Recibo)

	c.JSON(http.StatusOK, Recibo)
}

// POST /recibos
func CreateRecibo(c *gin.Context) {
	// Validate input
	var input models.CreateReciboInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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

// GET /recibos/:id
// Buscar un recibo
func FindRecibo(c *gin.Context) { // Get model if exist
	var recibo models.Recibo

	if err := models.DB.Where("id_recibo = ?", c.Param("id")).First(&recibo).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No encontrado!"})
		return
	}

	c.JSON(http.StatusOK, recibo)
}

// PATCH /recibos/:id
// Editar Recibo
func UpdateRecibo(c *gin.Context) {
	// Get model if exist
	var recibo models.Recibo
	if err := models.DB.Where("id_recibo = ?", c.Param("id")).First(&recibo).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No encontrado!"})
		return
	}

	// Validate input
	var input models.UpdateReciboInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&recibo).Updates(input)

	c.JSON(http.StatusOK, recibo)
}

// DELETE /recibos/:id
// Eliminar Recibo
func DeleteRecibo(c *gin.Context) {
	// Get model if exist
	var recibo models.Recibo
	if err := models.DB.Where("id_recibo = ?", c.Param("id")).First(&recibo).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No encontrado!"})
		return
	}

	models.DB.Delete(&recibo)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
