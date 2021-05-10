package views

import (
	_ "github.com/jinzhu/gorm"
)

type Tarjetas struct {
	Nro_tarjeta       string
	Fecha_vencimiento string
	Nombres_tarjeta   string
	Apellidos_tarjeta string
	Nombres           string
	Apellidos         string
}
