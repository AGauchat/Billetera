// models/ReciboController.go

package models

import (
	_ "github.com/jinzhu/gorm"
	"github.com/shopspring/decimal"
	"time"
)

type MovimientosCtaCte struct {
	IdMovimientosCtaCte int             `json:"IdMovimientosCtaCte" gorm:"primary_key;AUTO_INCREMENT;not null"`
	IdTipoMovimiento    int             `json:"IdTipoMovimiento" gorm:"not null"`
	IdCuentaCorriente   int             `json:"IdCuentaCorriente" gorm:"not null"`
	IdRecibo            int             `json:"IdRecibo" gorm:"not null"`
	CVUDesde            string          `json:"CVUDesde" gorm:"not null"`
	CVUHasta            string          `json:"CVUHasta" gorm:"not null"`
	Fecha               time.Time       `json:"Fecha" gorm:"not null"`
	Estado              string          `json:"Estado"`
	Monto               decimal.Decimal `json:"Monto" gorm:"not null"`
	Saldo               decimal.Decimal `json:"Saldo" gorm:"not null"`
}
