package email_usecase

import (
	"net-http/myapp/domain"
	"net-http/myapp/domain/model/email_domain"
	"net-http/myapp/utils"
)

type EmailUseCaseStrcut struct {
	AdminUserRepo domain.AdminUserRepository
	JwtRepo       domain.AuthJwtToken
	EmailRepo     domain.EmailRepository
}

func (receiver EmailUseCaseStrcut) CreateEmailUseCase(jwtToken string, subject string, body string, to string) error {
	userId, err := receiver.JwtRepo.AuthorizationProcess(jwtToken)
	if err != nil {
		return err
	}

	userData, err := receiver.AdminUserRepo.FindAdminUserById(userId)
	if err != nil {
		return err
	}

	if userData.IsLogin == false {
		return utils.MyError{Message: "ログインしてください"}
	}

	mailVo, err := email_domain.NewEmail(to)
	if err != nil {
		return err
	}

	mailBodyVo, err := email_domain.NewMailBody(body)
	if err != nil {
		return err
	}

	subjectVo, err := email_domain.NewSubject(subject)
	if err != nil {
		return err
	}

	emailEntity, err := email_domain.NewEmailEntity(subjectVo.String(), mailVo.String(), mailBodyVo.String())
	if err = receiver.EmailRepo.Create(emailEntity); err != nil {
		return err
	}
	return nil
}
