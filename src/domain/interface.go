package domain

import (
	"net-http/myapp/domain/model/email_domain"
	"net-http/myapp/domain/model/user"
)

// 管理者ユーザーのリポジトリ
type AdminUserRepository interface {
	SaveAdminUser(user *user.AdminUser) error
	FindAdminUserByEmail(email string) (*user.AdminUser, error)
	UpdateAdminUser(user *user.AdminUser) error
	FindAdminUserById(id float64) (*user.AdminUser, error)
}

// JWTTokenの処理
type AuthJwtToken interface {
	CreateJwtToken(serId uint) (string, error)
	AuthorizationProcess(tokenString string) (float64, error)
}

// Companyに紐つく処理
type CompanyRepository interface {
	GetUserData() ([]user.AdminUser, error)
}

// Emailに紐つく処理
type EmailRepository interface {
	Create(emailEntity *email_domain.EmailEntity) error
}
