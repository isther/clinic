package model

import "gorm.io/gorm"

type User struct {
	Name        string `json:"name" gorm:"name"`
	StudentID   string `json:"student_id" gorm:"student_id"`
	PhoneNumber string `json:"phone_number" gorm:"phone_number"`
	QQ          string `json:"qq" gorm:"qq"`
	Wechat      string `json:"wechat" gorm:"wechat"`
}

type Form struct {
	FormID       string `json:"form_id" gorm:"form_id"`
	User         `json:"user"`
	Content      string `json:"content" gorm:"content"`
	ExpectedTime string `json:"expected_time" gorm:"expected_time"`
	Status       bool   `json:"status" gorm:"status"`
}

type FormSql struct {
	gorm.Model
	Form `json:"form"`
}
