package services

import (
	"mqt-tech-broker/db"
	"time"
)

func SetOnline(deviceID string) {
	db.Rdb.Set(db.Ctx, "device:"+deviceID+":status", "online", time.Hour)
}

func SetOffline(deviceID string) {
	db.Rdb.Set(db.Ctx, "device:"+deviceID+":status", "offline", time.Hour)
}
