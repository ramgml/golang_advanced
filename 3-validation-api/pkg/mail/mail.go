package mail

import (
	"3-validation-api/pkg/vault"
	"fmt"
	"net/smtp"

	"github.com/jordan-wright/email"
)

func SendMail(address string, from string, password string, acc *vault.Account) error {
	link := fmt.Sprintf("http://localhost:8081/verify/%s", acc.Key)
	e := email.NewEmail()
	e.From = fmt.Sprintf("Jordan Wright <%s>", from)
	e.To = []string{acc.Email}
	e.Subject = "Verify your email"
	e.HTML = fmt.Appendf(nil, "<a href=\"%s\">Click me!</a>", link)
	err := e.Send(address, smtp.PlainAuth("", from, password, "smtp.gmail.com"))
	return err
}