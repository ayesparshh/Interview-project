package db

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Customers model
type Customers struct {
	ID      uint      `gorm:"primaryKey" json:"id"`
	Name    string    `gorm:"not null" json:"name"`
	Email   string    `gorm:"not null" json:"email"`
	Phone   string    `gorm:"not null" json:"phone"`
	Address string    `gorm:"not null" json:"address"`
	City    string    `gorm:"not null" json:"city"`
	State   string    `gorm:"not null" json:"state"`
	Zip     string    `gorm:"not null" json:"zip"`
	Date    time.Time `gorm:"not null" json:"date"`
}

// all customers
func (db DB) GetCustomers(c *gin.Context) {
	var customers []Customers
	if result := db.conn.Find(&customers); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, customers)
}

// create a new customer
func (db DB) CreateCustomer(c *gin.Context) {
	var customer Customers
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if result := db.conn.Create(&customer); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, customer)
}
