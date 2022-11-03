package model

import (
	"database/sql/driver"
	"encoding/json"

	"gorm.io/gorm"
)

type FormSql struct {
	gorm.Model
	Form `json:"form"`
}

type Form struct {
	FormID       string        `json:"form_id" gorm:"form_id"` // ID
	User         `json:"user"` // 用户
	Content      string        `json:"content" gorm:"content"`             // 问题
	ContentType  string        `json:"content_type" gorm:"content_type"`   // 问题类型
	ExpectedTime string        `json:"expected_time" gorm:"expected_time"` // 期望时间
	Status       string        `json:"status" gorm:"status"`               // 当前状态
	Imgs         Strs          `json:"imgs" gorm:"imgs"`
}

type User struct {
	Name        string `json:"name" gorm:"name"`
	StudentID   string `json:"student_id" gorm:"student_id"`
	ContactInfo string `json:"contact_info" gorm:"contact_info"`
	ContactWay  string `json:"contact_way" gorm:"contact_way"` // 联系方式
}

type Strs []string

func (s *Strs) Scan(value interface{}) error {
	var bytesValue = value.([]byte)
	return json.Unmarshal(bytesValue, s)
}

func (s Strs) Value() (driver.Value, error) {
	return json.Marshal(s)
}
