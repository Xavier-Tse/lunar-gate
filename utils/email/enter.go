package email

import (
	"fmt"
	"github.com/Xavier-Tse/lunar-gate/global"
	"github.com/jordan-wright/email"
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
