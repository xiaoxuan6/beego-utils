package email

import (
	"fmt"
	"github.com/astaxie/beego"
	email2 "github.com/jordan-wright/email"
	"net/smtp"
	"runtime"
	"strconv"
)

var Send chan *email2.Email

func init() {

	host := beego.AppConfig.String("e_host")
	port, _ := beego.AppConfig.Int("e_port")
	username := beego.AppConfig.String("e_username")
	password := beego.AppConfig.String("e_password")
	chanlimit := beego.AppConfig.DefaultInt("chan_limit", runtime.NumCPU())

	if len(host) < 1 || port < 1 || len(username) < 1 || len(password) < 1 {
		panic("email 配置信息错误")
	}

	add := fmt.Sprintf("%s:%s", host, strconv.Itoa(port))

	Send = make(chan *email2.Email, chanlimit)

	go func() {
		for {
			select {
			case m, ok := <-Send:
				if !ok {
					panic("发送失败")
					return
				}

				if er := m.Send(add, smtp.PlainAuth("", username, password, host)); er != nil {
					panic(fmt.Sprintf("email 发送失败: %s", er.Error()))
					return
				}
			}
		}
	}()
}
