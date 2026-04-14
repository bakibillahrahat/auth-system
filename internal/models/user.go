package models

import "time"

type User struct {
	ID	       uint 		`gorm:"primaryKey" json:"id"`
	FirstName  string		`validate:"required" json:"first_name"`
	LastName   string  		`validate:"required" json:"last_name"`
	Email 	   string  		`gorm:"unique,not null" validate:"required,email" json:"email"`
	Password   string  		`gorm:"not null" validate:"required" json:"password"`
	Address    Address 		`gorm:"embedded" json:"address"`
	AvatarURL  string  		`validate:"omitempty, url" json:"avatar_url"`
	CreatedAt  time.Time 	`json:"created_at"`
	UpdatedAt  time.Time 	`json:"updated_at"`
}

type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode string `json:"zip_code"`
}