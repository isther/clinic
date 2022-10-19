package model

import "gorm.io/gorm"

type User struct {
	Name        string `json:"name" gorm:"name"`
	StudentID   string `json:"student_id" gorm:"student_id"`
	PhoneNumber string `json:"phone_number" gorm:"phone_number"`
}

type Form struct {
	User         `json:"user"`
	Content      string `json:"content" gorm:"content"`
	ExpectedTime string `json:"expected_time" gorm:"expected_time"`
}

type FormSql struct {
	gorm.Model
	Form `json:"form"`
}
