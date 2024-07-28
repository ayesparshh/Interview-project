package db

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Billings model
type Billings struct {
	ID       uint      `gorm:"primaryKey" json:"id"`
	Customer uint      `gorm:"not null" json:"customer"`
	Invoice  uint      `gorm:"not null" json:"invoice"`
	Date     time.Time `gorm:"not null" json:"date"`
	Amount   float64   `gorm:"not null" json:"amount"`
}

func (db DB) GetBillings(c *gin.Context) {
	var billings []Billings
	if result := db.conn.Find(&billings); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, billings)
}

func (db DB) CreateBilling(c *gin.Context) {
	var billing Billings
	if err := c.ShouldBindJSON(&billing); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if result := db.conn.Create(&billing); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, billing)
}
