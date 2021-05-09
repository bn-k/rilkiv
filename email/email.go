package email

import (
	"bytes"
	"fmt"
	"github.com/bn-k/rilkiv/config"
	"github.com/bn-k/rilkiv/exchange"
	"gopkg.in/gomail.v2"
	"html/template"
)

type email struct {
	main string
	conf config.Config
}

func ProvideEmail(conf config.Config) (exchange.EmailClient, error) {
	return email{
		conf.Email.Address,
		conf,
	}, nil
}

func (e email) SendConfirmation(to, token string) error {
	t, err := template.New("hello.gohtml").ParseFiles("email/hello.gohtml")
	if err != nil {
		return err
	}

	buf := bytes.Buffer{}
	err = t.Execute(&buf, struct {
		Email   string
		Token   string
		Address string
	}{
		to,
		token,
		e.conf.WebClientAddress,
	})
	if err != nil {
		return err
	}

	m := gomail.NewMessage()
	m.SetHeader("From", e.main)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", buf.String())

	fmt.Println("conf: ", e.conf.Email.Address)
	d := gomail.NewDialer(e.conf.Email.SMTP, e.conf.Email.Port, e.conf.Email.Username, e.conf.Email.Password)


	err = d.DialAndSend(m)
	if err != nil {
		return err
	}

	return nil
}
