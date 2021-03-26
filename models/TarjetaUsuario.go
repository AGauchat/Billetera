// models/ReciboController.go

package models

import (
	_ "github.com/jinzhu/gorm"
)

type TarjetaUsuario struct {
	IdTarjeta        int64  `json:"IdCuentaCorriente" gorm:"primary_key;AUTO_INCREMENT;not null"`
	IdUsuario        int64  `json:"IdUsuario"`
	NroTarjeta       string `json:"NroTarjeta" gorm:"not null"`
	CodVerificacion  int    `json:"CodVerificacion" gorm:"not null"`
	FechaVencimiento string `json:"FechaVencimiento" gorm:"not null"`
	ApellidosTarjeta string `json:"ApellidosTarjeta" gorm:"not null"`
	NombresTarjeta   string `json:"NombresTarjeta" gorm:"not null"`
}
