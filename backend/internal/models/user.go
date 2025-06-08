package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID         uuid.UUID      `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Email      string         `json:"email" gorm:"uniqueIndex;not null"`
	Password   string         `json:"-" gorm:"not null"`
	Name       string         `json:"name" gorm:"not null"`
	Avatar     string         `json:"avatar"`
	Provider   string         `json:"provider" gorm:"default:'local'"` // 'local' or 'google'
	ProviderID string         `json:"provider_id"`
	IsVerified bool           `json:"is_verified" gorm:"default:false"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"-" gorm:"index"`

	Notes []Note `json:"notes,omitempty" gorm:"foreignKey:UserID"`
}
