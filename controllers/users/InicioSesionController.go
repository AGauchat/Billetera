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

	println(c.Request.Header.Get("dni"))
	println(c.Request.Header.Get("pass"))

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
