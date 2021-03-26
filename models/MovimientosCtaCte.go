// models/ReciboController.go

package models

import (
	_ "github.com/jinzhu/gorm"
	"github.com/shopspring/decimal"
	"time"
)

type MovimientosCtaCte struct {
	IdMovimientosCtaCte int64           `json:"IdMovimientosCtaCte" gorm:"primary_key;AUTO_INCREMENT;not null"`
	IdTipoMovimiento    int64           `json:"IdTipoMovimiento" gorm:"not null"`
	IdUsuario           int64           `json:"IdUsuario" gorm:"not null"`
	IdRecibo            int64           `json:"IdRecibo" gorm:"not null"`
	CVUHasta            string          `json:"CVUHasta" gorm:"not null"`
	Fecha               time.Time       `json:"Fecha" gorm:"not null"`
	Estado              string          `json:"Estado"`
	Monto               decimal.Decimal `json:"Monto" gorm:"not null"`
	Saldo               decimal.Decimal `json:"Saldo" gorm:"not null"`
}
