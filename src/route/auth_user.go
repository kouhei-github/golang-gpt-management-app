package route

import (
	"net-http/myapp/controller"
	"net-http/myapp/controller/email"
	userHandling "net-http/myapp/controller/user"
	"net-http/myapp/repository"
	"net-http/myapp/repository/auth_infra"
	"net-http/myapp/usecase/email_usecase"
	"net-http/myapp/usecase/user"
)

func (router *Router) GetAuthRouter() {
	router.Mutex.HandleFunc("/auth", controller.HandlerTwo)

	// 認証許可
	getUser := userHandling.NewGetUserHandler(&user.GetUser{
		AdminUserRepo: &repository.Administer{},
		JwtRepo:       &auth_infra.JwtToken{},
	})
	router.Mutex.HandleFunc("/api/v1/users", getUser.GetUserHandler)

	// 認証許可
	userExport := userHandling.NewUserExportCsvFile(&user.UserExportCsvFile{
		AdminUserRepo: &repository.Administer{},
		JwtRepo:       &auth_infra.JwtToken{},
	})
	router.Mutex.HandleFunc("/api/v1/users-export", userExport.ExportCsvUserHandler)

	// プロンプト一覧

	// プロンプト作成

	// プロンプト表示

	// email_usecase
	emailHandle := email.NewEmailCreate(&email_usecase.EmailUseCaseStrcut{
		AdminUserRepo: &repository.Administer{},
		JwtRepo:       &auth_infra.JwtToken{},
		EmailRepo:     &repository.EmailModel{},
	})
	router.Mutex.HandleFunc("/api/v1/email-create", emailHandle.CreateHandler)

	// emailの文章の取得
	emailGetHandle := email.EmailList{}
	router.Mutex.HandleFunc("/api/v1/email/list", emailGetHandle.ListEmailHandler)
}
