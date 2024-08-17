package persistence

import (
	"github.com/google/uuid"
	"time"
)

type UserEntity struct {
	ID        uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Username  string     `gorm:"type:varchar(50);unique;not null"`
	Password  string     `gorm:"type:varchar(255);not null"`
	Email     string     `gorm:"type:varchar(50);unique;not null"`
	CreatedAt time.Time  `gorm:"type:timestamp with time zone;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time  `gorm:"type:timestamp with time zone;default:CURRENT_TIMESTAMP"`
	DeletedAt *time.Time `gorm:"type:timestamp with time zone"`
}

func (*UserEntity) TableName() string {
	return "users"
}
