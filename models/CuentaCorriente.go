// models/ReciboController.go

package models

import (
	_ "github.com/jinzhu/gorm"
)

type CuentaCorriente struct {
	IdCuentaCorriente int    `json:"IdCuentaCorriente" gorm:"primary_key;AUTO_INCREMENT"`
	IdUsuario         int    `json:"IdUsuario" gorm:"not null"`
	CVU               string `json:"CVU" gorm:"not null"`
	Alias             string `json:"Alias" gorm:"not null"`
}
