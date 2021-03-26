// models/ReciboController.go

package models

import (
	_ "github.com/jinzhu/gorm"
)

type Recibo struct {
	IdRecibo       int64  `json:"IdRecibo" gorm:"primary_key;AUTO_INCREMENT"`
	IdTablaApi     int64  `json:"IdTablaApi"`
	NombreTablaApi string `json:"NombreTablaApi"`
}

type CreateReciboInput struct {
	IdTablaApi     int64  `json:"IdTablaApi" binding:"required"`
	NombreTablaApi string `json:"NombreTablaApi"`
}

type UpdateReciboInput struct {
	IdTablaApi     int64  `json:"IdTablaApi"`
	NombreTablaApi string `json:"NombreTablaApi"`
}
