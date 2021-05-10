package users

import (
	"billetera/models"
	"billetera/models/views"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func FindTarjeta(c *gin.Context) {
	pi, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	var Tarjetas []views.Tarjetas
	var db = models.DB

	if err := db.Table("tarjeta_usuarios").Select("tarjeta_usuarios.nro_tarjeta, tarjeta_usuarios.fecha_vencimiento, tarjeta_usuarios.nombres_tarjeta, tarjeta_usuarios.apellidos_tarjeta, usuarios.nombres, usuarios.apellidos").Joins("inner join usuarios on usuarios.id_usuario = tarjeta_usuarios.id_usuario").Where("tarjeta_usuarios.id_usuario = ?", pi).Scan(&Tarjetas).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No encontrado!"})
		return
	}

	c.JSON(http.StatusOK, Tarjetas)
}
