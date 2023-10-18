package email_domain

import (
	"gorm.io/gorm"
)

type EmailEntity struct {
	gorm.Model
	SendTo   string `json:"send_to" binding:"required"`
	MailBody string `json:"mail_body"  binding:"required"`
	Subject  string `json:"subject"`
}

func NewEmailEntity(subject string, to string, mailBody string) (*EmailEntity, error) {
	return &EmailEntity{SendTo: to, MailBody: mailBody, Subject: subject}, nil
}

func (email EmailEntity) ArrayString() ([]string, error) {
	return []string{email.SendTo, email.Subject, email.MailBody}, nil
}
