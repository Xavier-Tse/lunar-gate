package email

import (
	"fmt"
	"github.com/jordan-wright/email"
	"github.com/lunarise-dev/lunar-gate/global"
	"net/smtp"
	"strings"
)

func SendEmail(to, subject, text string) error {
	em := global.Config.Email
	e := email.NewEmail()
	e.From = fmt.Sprintf("%s <%s>", em.SendNickname, em.SendEmail)
	e.To = []string{to}
	e.Subject = subject
	e.Text = []byte(text)
	err := e.Send(fmt.Sprintf("%s:%d", em.Domain, em.Port), smtp.PlainAuth("", em.SendEmail, em.AuthCode, em.Domain))
	if err != nil && !strings.Contains(err.Error(), "short response:") {
		return err
	}
	return nil
}
