package models

import (
	_ "github.com/jinzhu/gorm"
	"github.com/shopspring/decimal"
	"time"
)

type ReciboBIND struct {
	IdReciboBIND int             `json:"IdReciboBIND" gorm:"primary_key;AUTO_INCREMENT"`
	Fecha        time.Time       `json:"Fecha"`
	Monto        decimal.Decimal `json:"Monto"`
}
