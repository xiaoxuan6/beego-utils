package services

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/dchest/captcha"
)

type captchaService struct {
}

type Captcha struct {
	Id string `json:"id"` // 验证码ID
}

var CaptchaService = newCaptchaService()

func newCaptchaService() *captchaService {
	return new(captchaService)
}

func (c *captchaService) NewCaptcha() (cap *Captcha) {
	length := beego.AppConfig.DefaultInt("len", captcha.DefaultLen)

	captchaId := captcha.NewLen(length)

	cap = &Captcha{
		captchaId,
	}

	return cap
}

func (c *captchaService) Verify(captchaId, captchaVal string) error {
	if ok := captcha.VerifyString(captchaId, captchaVal); !ok {
		return fmt.Errorf("验证失败")
	}

	return nil
}

func (c *captchaService) GetCaptchaBase64ById(captchaId string) (str string, error error) {
	var content bytes.Buffer

	stdWidth := beego.AppConfig.DefaultInt("std_width", captcha.StdWidth)
	stdHeight := beego.AppConfig.DefaultInt("std_height", captcha.StdHeight)

	if err := captcha.WriteImage(&content, captchaId, stdWidth, stdHeight); err != nil {
		return "", err
	}

	str = base64.StdEncoding.EncodeToString(content.Bytes())
	return str, nil
}
