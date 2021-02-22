// models/setup.go

package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

var DB *gorm.DB

func ConnectDataBase() {
	database, err := gorm.Open("mssql", "sqlserver://sa:Luse_2010@localhost:1433?database=Billetera")

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&Recibo{})
	database.AutoMigrate(&ReciboBIND{})
	database.AutoMigrate(&CuentaCorriente{})
	database.AutoMigrate(&TipoMovimiento{})
	database.AutoMigrate(&TarjetaUsuario{})
	database.AutoMigrate(&MovimientosCtaCte{})
	database.AutoMigrate(&Usuario{})

	DB = database
}
