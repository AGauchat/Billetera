// models/ReciboController.go

package models

import (
	_ "github.com/jinzhu/gorm"
)

//goland:noinspection ALL
type Usuario struct {
	IdUsuario           int    `json:"IdUsuario" gorm:"primary_key;AUTO_INCREMENT;not null"`
	Activo              bool   `json:"Activo"`
	IdentidadVerificada bool   `json:"IdentidadVerificada"`
	Nombres             string `json:"Nombres" gorm:"not null"`
	Apellidos           string `json:"Apellidos" gorm:"not null"`
	DNI                 string `json:"DNI"`
	Cuil                string `json:"Cuil" gorm:"not null"`
	Contraseña          string `json:"Contraseña"`
	Direccion           string `json:"Direccion" gorm:"not null"`
	NroDireccion        string `json:"NroDireccion"`
	NroTransaccion      string `json:"NroTransaccion"`
	IniciosFallidos     int    `json:"IniciosFallidos"`
	NroCelular          string `json:"NroCelular"`
}

type CreateUsuarioInput struct {
	Nombres        string `json:"Nombres" gorm:"not null"`
	Apellidos      string `json:"Apellidos" gorm:"not null"`
	DNI            string `json:"DNI"`
	Cuil           string `json:"Cuil" gorm:"not null"`
	Contraseña     string `json:"Contraseña"`
	Direccion      string `json:"Direccion" gorm:"not null"`
	NroDireccion   string `json:"NroDireccion"`
	NroTransaccion string `json:"NroTransaccion"`
	NroCelular     string `json:"NroCelular"`
}
