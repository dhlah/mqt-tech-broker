package models

import (
	"time"
)

// ==================== VirtualPin ====================
type VirtualPin struct {
	ID          string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	VirtualName string // misalnya V0, V1, V2
	Type        string // ex: "switch", "range", "sensor"
	DeviceID    string
	Device      Device          `gorm:"foreignKey:DeviceID;references:ID"`
	Logs        []VirtualPinLog `gorm:"foreignKey:VirtualPinID;references:ID"`
	CreatedAt   time.Time       `gorm:"autoCreateTime"`
}

// ==================== VirtualPinLog ====================
type VirtualPinLog struct {
	ID           string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Value        string    // data/value dari pin
	Timestamp    time.Time `gorm:"autoCreateTime"`
	VirtualPinID string
	VirtualPin   VirtualPin `gorm:"foreignKey:VirtualPinID;references:ID"`
}

// ==================== User ====================
type User struct {
	ID        string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Username  string `gorm:"unique"`
	Password  string
	Token     string
	Role      string
	CreatedAt time.Time `gorm:"autoCreateTime"`

	RefreshTokens []RefreshToken `gorm:"foreignKey:UserID;references:ID"`
}

// ==================== RefreshToken ====================
type RefreshToken struct {
	ID        string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Token     string `gorm:"unique"`
	UserID    string
	User      User      `gorm:"foreignKey:UserID;references:ID"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

// ==================== Section ====================
type Section struct {
	ID      string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name    string
	Devices []Device `gorm:"foreignKey:RoomID;references:ID"`
}

// ==================== Devices ====================
type Device struct {
	ID          string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name        string
	Status      string
	Token       string
	RoomID      string       `gorm:"type:uuid"`
	Room        Section      `gorm:"foreignKey:RoomID;references:ID"`
	VirtualPins []VirtualPin `gorm:"foreignKey:DeviceID;references:ID"`
	CreatedAt   time.Time    `gorm:"autoCreateTime"`
}
