package mail

import (
	"testing"
)

func TestSend(t *testing.T) {
	options := &Options{
		MailHost: "smtp.163.com",
		MailPort: 465,
		MailUser: "77086539@163.com",
		MailPass: "UBHULVZFJCYZINQY", //密码或授权码
		MailTo:   "zhou.jie@xhsoftware.cn",
		Subject:  "subject",
		Body:     "body",
	}
	err := Send(options)
	if err != nil {
		t.Error("Mail Send error", err)
		return
	}
	t.Log("success")
}
