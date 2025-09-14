package services

import (
	"mqt-tech-broker/db"
)

func ValidateDevice(deviceID, token string) bool {
	var dbToken string

	// Ambil token langsung dari database (plain)
	err := db.Pg.Table("devices").
		Select("token").
		Where("id = ?", deviceID).
		Scan(&dbToken).Error

	if err != nil || dbToken == "" {
		return false
	}

	// Bandingkan plain string
	return dbToken == token
}
