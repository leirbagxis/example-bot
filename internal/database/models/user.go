package models

import "time"

type User struct {
	UserId    int64     `gorm:"primaryKey" json:"id"` // ID do Telegram
	FirstName string    `json:"firstName"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
