package email

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	email2 "github.com/jordan-wright/email"
	"goelastic/core/config"
	"net/smtp"
	"strconv"
)

var Send chan *email2.Email

func init() {
	Econf := config.GetInstace().Email
	add := fmt.Sprintf("%s:%s", Econf.Host, strconv.Itoa(Econf.Port))

	Send = make(chan *email2.Email, Econf.ChanLimit)

	go func() {
		for {
			select {
			case m, ok := <-Send:
				if !ok {
					return
				}

				if er := m.Send(add, smtp.PlainAuth("", Econf.Username, Econf.Password, Econf.Host)); er != nil {
					logs.GetLogger().Println(er.Error())
				}
			}
		}
	}()
}
