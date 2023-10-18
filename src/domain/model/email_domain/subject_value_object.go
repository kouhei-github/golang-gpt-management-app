package email_domain

import "fmt"

// MailBody パスワードオブジェクト
type Subject struct {
	value string
}

func NewSubject(body string) (*Subject, error) {
	// emailアドレスのフォーマットが正しいか確認
	if len(body) < 5 {
		return nil, fmt.Errorf("文章をちゃんと書いてください")
	}
	return &Subject{value: body}, nil
}

// 文字列
// @return
// パスワード
func (body *Subject) String() string {
	return string(body.value)
}
