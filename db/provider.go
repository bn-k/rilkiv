package db

import (
	"fmt"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/bn-k/rilkiv/config"
	"github.com/bn-k/rilkiv/exchange"
)

type handler struct {
	gorm *gorm.DB
}

type db struct {
	exchange.Users
}

func Provide(conf config.Config, log *zap.Logger) (exchange.DB, error) {
	conn, err := CreateConnection(conf, log)
	if err != nil {
		return nil, err
	}


	return db{
		users{conn},
	}, MigrateAll(conn)
}

func MigrateAll(conn *gorm.DB) error {
	return conn.AutoMigrate(
		&exchange.User{},
	)
}


func CreateConnection(conf config.Config, log *zap.Logger) (*gorm.DB, error) {
	return gorm.Open(postgres.New(postgres.Config{
		DSN:                  fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", conf.Db.Host, conf.Db.User, conf.Db.Password, conf.Db.Database, conf.Db.Port),
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
}
