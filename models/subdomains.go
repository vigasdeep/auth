package models

import (
	"github.com/jinzhu/gorm"
)

//Subdomains struct declaration
type Subdomains struct {
	gorm.Model

	Name        string `binding:"required"`
	Pointer     string `binding:"required"`
	Description string `binding:"required"`
	UserID      int
	User        User
}
