package db

import (
	"errors"
	"time"
	
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

// Users model
type Users struct {
	ID        uint	 `gorm:"primaryKey" json:"id"`
	Name      string `gorm:"not null" json:"name"`
	Email     string `gorm:"not null" json:"email"`
	Password  string `gorm:"not null" json:"password"`
	Date      time.Time `gorm:"not null" json:"date"`
}