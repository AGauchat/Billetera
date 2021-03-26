package controllers

import (
	"billetera/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUsuario(c *gin.Context) {
	// Validate input
	var input models.CreateUsuarioInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Verifico que el CUIL no se repita
	var usuarioF models.Usuario
	if err := models.DB.Where("cuil = ?", input.Cuil).Find(&usuarioF).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Usuario ya registrado!"})
		return
	}

	// Create recibo
	usuario := models.Usuario{
		Nombres:        input.Nombres,
		Apellidos:      input.Apellidos,
		DNI:            input.DNI,
		Cuil:           input.Cuil,
		Contraseña:     PassHash(input.Contraseña),
		Direccion:      input.Direccion,
		NroDireccion:   input.NroDireccion,
		NroTransaccion: input.NroTransaccion,
		NroCelular:     input.NroCelular,
	}
	models.DB.Create(&usuario)

	c.JSON(http.StatusOK, usuario)
}
