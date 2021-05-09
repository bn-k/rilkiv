package app

import (
	"github.com/bn-k/rilkiv/config"
	"github.com/bn-k/rilkiv/db"
	"github.com/bn-k/rilkiv/email"
	"github.com/bn-k/rilkiv/exchange"
	"go.uber.org/zap"
)

type App struct {
	Conf config.Config
	Log  *zap.Logger
	DB   exchange.DB
	Mail exchange.EmailClient
}

func Provide() (App, error) {
	conf, err := config.Provide()
	if err != nil {
		panic(err)
	}

	log, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	log.Info("start provider", zap.Any("conf", conf))
	mail, err := email.ProvideEmail(conf)
	if err != nil {
		panic(err)
	}

	dbal, err := db.Provide(conf, log)
	return App{
		conf,
		log,
		dbal,
		mail,
	}, nil
}
