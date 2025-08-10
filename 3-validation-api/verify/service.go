package verify

import (
	"fmt"
	"main/config"
	"net/smtp"

	"github.com/jordan-wright/email"
)

type VerifyService struct {
	config *config.AppConfig
}

type unencryptedAuth struct {
	smtp.Auth
}

func (a unencryptedAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	s := *server
	s.TLS = true
	return a.Auth.Start(&s)
}

func (verifyService *VerifyService) SendEmail() {
	cfg := verifyService.config
	e := email.NewEmail()
	e.From = "Test<test@gmail.com>"
	e.To = []string{"test@mail.ru"}
	e.Subject = "Blablabla"
	e.HTML = []byte("<h1>Verify email<br><a href='http://localhost:8080/verify/hash'>Click here</a></h1>")

	address := fmt.Sprintf("%s:%s", cfg.SmtpConfig.Address, cfg.SmtpConfig.Port)

	auth := unencryptedAuth{
		smtp.PlainAuth(
			"",
			cfg.SmtpConfig.Email,
			cfg.SmtpConfig.Password,
			cfg.SmtpConfig.Address,
		),
	}

	err := e.Send(address, auth)

	if err != nil {
		fmt.Printf(err.Error())
	}
}
