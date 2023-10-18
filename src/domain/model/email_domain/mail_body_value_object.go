package email_domain

import (
	"fmt"
)

// MailBody パスワードオブジェクト
type MailBody struct {
	value string
}

func NewMailBody(body string) (*MailBody, error) {
	// emailアドレスのフォーマットが正しいか確認
	if len(body) < 5 {
		return nil, fmt.Errorf("文章をちゃんと書いてください")
	}
	return &MailBody{value: body}, nil
}

// 文字列
// @return
// パスワード
func (body *MailBody) String() string {
	return string(body.value)
}
