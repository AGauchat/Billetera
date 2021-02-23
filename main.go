package main

import (
	"billetera/controllers"
	"billetera/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	models.ConnectDataBase()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})
	r.GET("/recibos", controllers.FindRecibos)
	r.GET("/recibos/:id", controllers.FindRecibo)
	r.POST("/recibos", controllers.CreateRecibo)
	r.PATCH("/recibos/:id", controllers.UpdateRecibo)
	r.DELETE("/recibos/:id", controllers.DeleteRecibo)

	r.GET("/inicioSesion/:dni/:pass", controllers.InicioSesion)

	r.POST("/usuario", controllers.CreateUsuario)

	r.Run()
}
