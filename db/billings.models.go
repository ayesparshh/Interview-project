package db

import (
	"errors"
	"time"
	
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

// Billings model
type Billings struct {
	ID        uint	    `gorm:"primaryKey" json:"id"`
	Customer  uint	    `gorm:"not null" json:"customer"`
	Invoice   uint	    `gorm:"not null" json:"invoice"`
	Date      time.Time `gorm:"not null" json:"date"`
	Amount    float64   `gorm:"not null" json:"amount"`
}