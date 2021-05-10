// models/setup.go

package models

import (
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDataBase() {
	dsn := "sqlserver://sa:Luse_2010@localhost:1433?database=Billetera"
	database, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	_ = database.AutoMigrate(&Recibo{})
	_ = database.AutoMigrate(&ReciboBIND{})
	_ = database.AutoMigrate(&TipoMovimiento{})
	_ = database.AutoMigrate(&TarjetaUsuario{})
	_ = database.AutoMigrate(&MovimientosCtaCte{})
	_ = database.AutoMigrate(&Usuario{})

	DB = database
}

func ConnectDataBasePrueba() *gorm.DB {
	dsn := "sqlserver://sa:Luse_2010@localhost:1433?database=Billetera"
	database, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	return database
}

func Migrate() bool {
	var db = ConnectDataBasePrueba()

	_ = db.AutoMigrate(&Recibo{})
	_ = db.AutoMigrate(&ReciboBIND{})
	_ = db.AutoMigrate(&TipoMovimiento{})
	_ = db.AutoMigrate(&TarjetaUsuario{})
	_ = db.AutoMigrate(&MovimientosCtaCte{})
	_ = db.AutoMigrate(&Usuario{})

	return true
}
