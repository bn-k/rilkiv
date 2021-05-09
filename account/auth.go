package account

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"math/rand"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/bn-k/rilkiv/exchange"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

const (
	initialBalance = 100
)

var validate *validator.Validate

type UserRegister struct {
	Email     string `validate:"required,email"`
	FirstName string `validate:"required,alphaunicode"`
	LastName  string `validate:"required,alphaunicode"`
	Password  string `validate:"required,gte=8,lte=50"`
}

func init() {
	validate = validator.New()
	rand.Seed(time.Now().UnixNano())
}

func (h *Handlers) fmtUser(ctx context.Context, req UserRegister, r *http.Request) (exchange.User, error) {
	res := exchange.User{}
	err := validate.Struct(req)
	if err != nil {
		return res, err
	}

	err = validateMX(req.Email)
	if err != nil {
		return res, err
	}

	confirmToken := base64.URLEncoding.EncodeToString([]byte(req.Email + h.Conf.Auth.Secret))
	password, err := hashPassword(req.Password)
	if err != nil {
		return res, err
	}

	return exchange.User{
		Email:     req.Email,
		Firstname: req.FirstName,
		Lastname:  req.LastName,
		Role:      exchange.Client,
		Auth: exchange.Auth{
			Password:     password,
			ConfirmToken: confirmToken,
			Confirmed:    false,
		},
		Wallets: []exchange.Wallet{
			{
				Address:  gofakeit.BitcoinAddress(),
				Currency: exchange.BTC,
			},
			{
				Address:  gofakeit.BitcoinAddress(),
				Currency: exchange.ETH,
			},
		},
		Registration: exchange.Registration{
			IPAddress: r.RemoteAddr,
			UserAgent: r.UserAgent(),
		},
	}, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func validateMX(email string) error {
	_, host := split(email)
	_, err := net.LookupMX(host)
	if err != nil {
		return fmt.Errorf("bad dns")
	}

	return nil
}

func split(email string) (account, host string) {
	i := strings.LastIndexByte(email, '@')
	account = email[:i]
	host = email[i+1:]
	return
}
