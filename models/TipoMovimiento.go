// models/ReciboController.go

package models

import (
	_ "github.com/jinzhu/gorm"
)

type TipoMovimiento struct {
	IdTipoMovimiento int    `json:"IdTipoMovimiento" gorm:"primary_key"`
	Nombre           string `json:"Nombre"`
}
