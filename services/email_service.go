package services

import (
	"fmt"
	"github.com/astaxie/beego"
	email2 "github.com/jordan-wright/email"
	email1 "github.com/xiaoxuan6/beego-utils/email"
	"time"
)

type emailService struct {
	to      []string
	cc      []string
	subject string
}

var EmailService = newEmailService()

func newEmailService() *emailService {
	return new(emailService)
}

func (e *emailService) To(to []string) *emailService {
	e.to = to
	return e
}

func (e *emailService) CC(cc []string) *emailService {
	e.cc = cc
	return e
}

func (e *emailService) Subject(subject string) *emailService {
	e.subject = subject
	return e
}

func (e *emailService) Send(text string) error {

	to := e.to
	if len(to) < 1 {
		return fmt.Errorf("接受者不能为空")
	}

	email := &email2.Email{
		From: beego.AppConfig.String("e_from"),
		Text: []byte(text),
		To:   to,
	}

	if len(e.cc) > 0 {
		email.Cc = e.cc
	}

	if e.subject != "" {
		email.Subject = e.subject
	}

	select {
	case email1.Send <- email:
		return nil
	case <-time.After(time.Second * 3): // (管道阻塞时)等待 3s
		return fmt.Errorf("发送超时")
	}
}
