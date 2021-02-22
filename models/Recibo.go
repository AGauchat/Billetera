// models/ReciboController.go

package models

import (
	_ "github.com/jinzhu/gorm"
)

type Recibo struct {
	IdRecibo       int    `json:"IdRecibo" gorm:"primary_key;AUTO_INCREMENT"`
	IdTablaApi     int    `json:"IdTablaApi"`
	NombreTablaApi string `json:"NombreTablaApi"`
}

type CreateReciboInput struct {
	IdTablaApi     int    `json:"IdTablaApi" binding:"required"`
	NombreTablaApi string `json:"NombreTablaApi"`
}

type UpdateReciboInput struct {
	IdTablaApi     int    `json:"IdTablaApi"`
	NombreTablaApi string `json:"NombreTablaApi"`
}
