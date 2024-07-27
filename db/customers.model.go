package db

import (
	"errors"
	"time"
	
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

// Customers model
type Customers struct {
	ID        uint	    `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"not null" json:"name"`
	Email     string    `gorm:"not null" json:"email"`
	Phone     string    `gorm:"not null" json:"phone"`
	Address   string    `gorm:"not null" json:"address"`
	City      string    `gorm:"not null" json:"city"`
	State     string    `gorm:"not null" json:"state"`
	Zip       string    `gorm:"not null" json:"zip"`
	Date      time.Time `gorm:"not null" json:"date"`
}
