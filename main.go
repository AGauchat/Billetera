package main

import (
	"billetera/controllers/recibos"
	"billetera/controllers/users"
	"billetera/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDataBase()

	r.POST("/inicioSesion/", users.InicioSesion)

	r.GET("/GetIdUsrporDni/:dni", users.GetIdUsrporDni)

	//r.GET("/refreshToken", middleware.Refresh)
	//r.Use(middleware.VerifyToken)

	r.GET("/recibos", recibos.FindRecibos)
	r.GET("/recibos/:id", recibos.FindRecibo)
	r.POST("/recibos", recibos.CreateRecibo)
	r.PATCH("/recibos/:id", recibos.UpdateRecibo)
	r.DELETE("/recibos/:id", recibos.DeleteRecibo)

	r.POST("/usuario", users.CreateUsuario)

	r.GET("/tarjetaUsuario/:id", users.FindTarjeta)

	r.Run()
}
