package email

import (
	"encoding/json"
	"net-http/myapp/controller"
	"net-http/myapp/usecase/email_usecase"
	"net-http/myapp/utils"
	"net/http"
)

type responseBody struct {
	Message string `json:"message"`
}

type EmailCreate struct {
	Service *email_usecase.EmailUseCaseStrcut
}

func NewEmailCreate(s *email_usecase.EmailUseCaseStrcut) *EmailCreate {
	return &EmailCreate{Service: s}
}

func (receiver EmailCreate) CreateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "POST" {
		response := controller.Response{Status: 405, Text: "Method Not Allowed"}
		w.WriteHeader(405)
		json.NewEncoder(w).Encode(response)
		return
	}

	// jwtTokenの認証
	jwtToken := r.Header.Get("Authorization")
	// パラメータやメソッドなどが諸々正しいことが確認できたら、

	// ユーザのリクエストパラメータを構造体にマッピング
	var input struct {
		Subject string `json:"subject" binding:"required"`
		Body    string `json:"body" binding:"required"`
		To      string `json:"send_to" binding:"required"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response := controller.Response{Status: 401, Text: "入力内容をお確かめください"}
		w.WriteHeader(response.Status)
		json.NewEncoder(w).Encode(response)
		utils.WriteLogFile(err.Error())
		return
	}

	// サインイン処理はユースケースに依頼
	err := receiver.Service.CreateEmailUseCase(jwtToken, input.Subject, input.Body, input.To)
	if err != nil {
		response := controller.Response{Status: 500, Text: err.Error()}
		w.WriteHeader(response.Status)
		json.NewEncoder(w).Encode(response)
		utils.WriteLogFile(err.Error())
		return
	}
	response := responseBody{Message: "成功しました"}
	json.NewEncoder(w).Encode(response)
	utils.WriteLogFile("完了しました")
}
