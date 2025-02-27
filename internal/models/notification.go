package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type NotificationTemplate struct {
	ID           int
	TemplateName string `gorm:"column:template_name;type:varchar(255);unique"`
	Subject      string `gorm:"column:subject;type:varchar(255)"`
	Body         string `gorm:"column:body;type:text"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (*NotificationTemplate) TableName() string {
	return "notification_templates"
}

type NotificationHistory struct {
	ID           int
	Recipient    string `gorm:"column:recipient;type:varchar(255)"`
	TemplateID   int    `gorm:"column:template_id"`
	Status       string `gorm:"column:status;type:varchar(10)"`
	ErrorMessage string `gorm:"column:error_message;type:text"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (*NotificationHistory) TableName() string {
	return "notification_history"
}

type InternalNotificationRequest struct {
	TemplateName string `valid:"required"`
	Recipient    string `valid:"required;email"`
	Placeholders map[string]string
}

func (I *InternalNotificationRequest) Validate() error {
	v := validator.New()
	return v.Struct(I)
}
