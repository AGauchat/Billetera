package users

import (
	"billetera/controllers/middleware"
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

	dni := c.Request.Header.Get("dni")
	pass := c.Request.Header.Get("pass")

	db := models.ConnectDataBasePrueba()
	if err := db.Where("dni = ? AND contraseña = ?", dni, PassHash(pass)).First(&usuario).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No Encontrado!"})
		return
	}

	t := middleware.GenerateToken(usuario.IdUsuario)

	c.SetCookie("token", t, 360, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{"token": t})
}

func GetIdUsrporDni(c *gin.Context) {
	var usuario models.Usuario
	q := models.DB.Find(&usuario, "dni = ?", c.Param("dni"))
	if q.Error != nil || c.Errors != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error_massage": q.Error.Error() + " / " + c.Errors.String()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user_id": usuario.IdUsuario})
	return
}
