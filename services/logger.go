package services

import (
	"mqt-tech-broker/db"
	"mqt-tech-broker/models"
	"strings"
	"time"
)

func LogMessage(topic, payload string) {
	// topic format: deviceID/virtualPinID
	parts := strings.Split(topic, "/")
	if len(parts) < 2 {
		return // format salah, tidak disimpan
	}

	virtualPinID := parts[1] // ambil UUID virtual pin

	log := models.VirtualPinLog{
		VirtualPinID: virtualPinID,
		Value:        payload,
		Timestamp:    time.Now(),
	}

	db.Pg.Create(&log)
}
