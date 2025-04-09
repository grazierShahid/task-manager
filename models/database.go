package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type Task struct {
	gorm.Model
	Title       string `json:"title" gorm:"not null"`
	Description string `json:"description"`
	DueDate     string `json:"due_date"`
	Status      string `json:"status" gorm:"default:'pending'"`
	Priority    string `json:"priority" gorm:"default:'medium'"`
	UserID      uint   `json:"user_id" gorm:"not null"`
	User        User   `json:"user" gorm:"foreignKey:UserID"`
}
