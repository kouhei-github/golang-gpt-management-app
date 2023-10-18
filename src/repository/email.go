package repository

import (
	"net-http/myapp/domain/model/email_domain"
)

type EmailModel struct {
}

func (repository *EmailModel) Create(emailEntity *email_domain.EmailEntity) error {
	result := db.Create(&emailEntity)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repository *EmailModel) GetAll() (*[]email_domain.EmailEntity, error) {
	var emailEntity []email_domain.EmailEntity
	// Limitは何個取得するか
	// Offsetは何個スキップするか
	result := db.Limit(3).Offset(3).Find(&emailEntity)
	if result.Error != nil {
		return &emailEntity, result.Error
	}
	return &emailEntity, nil
}
