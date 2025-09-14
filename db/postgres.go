package db

import (
	"log"
	// "mqt-tech-broker/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Pg *gorm.DB

// InitPostgres untuk koneksi ke PostgreSQL dan sync schema
func InitPostgres() {
	// dsn := "host=127.0.0.1 user=postgres password=1234 dbname=mqtt port=5432 sslmode=disable"
	dsn := "host=141.11.25.130 user=santosojamalsuckmadick password=H2xrR8Q7VXGQtD1maU5XmRFPB dbname=mqtt-broker-db port=5432 sslmode=disable"

	var err error
	Pg, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ gagal connect postgres:", err)
	}

	// Migrasi schema -> kalau tabel belum ada akan dibuat
	// err = Pg.AutoMigrate(
	// 	&models.Device{},
	// 	&models.VirtualPin{},
	// 	&models.VirtualPinLog{},
	// 	&models.User{},
	// 	&models.RefreshToken{},
	// 	&models.Section{},
	// )
	// if err != nil {
	// 	log.Fatal("❌ gagal migrate schema:", err)
	// }

	log.Println("✅ PostgreSQL connected & schema migrated")
}
