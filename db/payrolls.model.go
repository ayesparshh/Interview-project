package db

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Payrolls model
type Payrolls struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	EmployeeID uint      `gorm:"not null" json:"employee_id"`
	Date       time.Time `gorm:"not null" json:"date"`
	Amount     uint      `gorm:"not null" json:"amount"`
}

// all payrolls
func (db DB) GetPayrolls(c *gin.Context) {
	var payrolls []Payrolls
	if result := db.conn.Find(&payrolls); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, payrolls)
}

// create a new payroll
func (db DB) CreatePayroll(c *gin.Context) {
	var payroll Payrolls
	if err := c.ShouldBindJSON(&payroll); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if result := db.conn.Create(&payroll); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, payroll)
}
