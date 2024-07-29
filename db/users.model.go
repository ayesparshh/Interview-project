package db

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Users model
type Users struct {
	ID       uint      `gorm:"primaryKey" json:"id"`
	Name     string    `gorm:"not null" json:"name"`
	Email    string    `gorm:"not null" json:"email"`
	Password string    `gorm:"not null" json:"password"`
	Date     time.Time `gorm:"not null" json:"date"`
}

// all users
func (db DB) GetUsers(c *gin.Context) {
	var users []Users
	if result := db.conn.Find(&users); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

// create a new user
func (db DB) CreateUser(c *gin.Context) {
	var user Users
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if result := db.conn.Create(&user); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}
