package models

import (
	"time"

	"gorm.io/gorm"
)

type Hotel struct {
	ID          int            `gorm:"primaryKey;column:id" json:"id"`
	Name        string         `json:"name" gorm:"type:varchar(255);not null"`
	Description string         `json:"description" gorm:"type:text"`
	Address     string         `json:"address" gorm:"type:varchar(255)"`
	City        string         `json:"city" gorm:"type:varchar(100)"`
	State       string         `json:"state" gorm:"type:varchar(100)"`
	Country     string         `json:"country" gorm:"type:varchar(100)"`
	PostalCode  string         `json:"postal_code" gorm:"type:varchar(20)"`
	PhoneNumber string         `json:"phone_number" gorm:"type:varchar(20)"`
	Email       string         `json:"email" gorm:"type:varchar(100)"`
	WebsiteURL  string         `json:"website_url" gorm:"type:varchar(255)"`
	Latitude    float64        `json:"latitude" gorm:"type:decimal(10,8)"`
	Longitude   float64        `json:"longitude" gorm:"type:decimal(11,8)"`
	Rating      float32        `json:"rating" gorm:"type:decimal(2,1)"`
	CreatedAt   time.Time      `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at" json:"-"`
}
