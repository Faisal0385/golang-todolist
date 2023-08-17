package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FullName string `json:"full_name"`
	Mobile   string `json:"mobile"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Photo    string `json:"photo"`
	Address  string `json:"address"` // text - null value
	Role     string `json:"role"`    // enum ['admin', 'user', 'vendor']
	Status   string `json:"status"`  // how to give default value
}

type Todo struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	User        string `json:"image"`
	Status      string `json:"status"`
}
