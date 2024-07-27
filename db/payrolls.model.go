package db

import (
	"errors"
	"time"
	
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

// Payrolls model
type Payrolls struct {
	ID        	uint	  `gorm:"primaryKey" json:"id"`
	EmployeeID 	uint	  `gorm:"not null" json:"employee_id"`
	Date      	time.Time `gorm:"not null" json:"date"`
	Amount    	uint	  `gorm:"not null" json:"amount"`
}