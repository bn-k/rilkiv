package db

import (
	"fmt"
	"time"

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
	exchange.Wallets
	exchange.Transactions
	exchange.Registrations
}

func Provide(conf config.Config, log *zap.Logger) (exchange.DB, error) {
	conn, err := CreateConnection(conf, log)
	if err != nil {
		return nil, err
	}

	return db{
		users{conn},
		wallets{conn},
		transactions{conn},
		registrations{conn},
	}, MigrateAll(conn)
}

func DropMigrateAll(conf config.Config) error {
	conn, err := CreateConnection(conf, nil)
	if err != nil {
		panic(err)
	}

	conn.Migrator().DropTable(
		&exchange.Transaction{},
		&exchange.Registration{},
		&exchange.User{},
		&exchange.Wallet{})

	return conn.AutoMigrate(
		&exchange.Transaction{},
		&exchange.Registration{},
		&exchange.Wallet{},
		&exchange.User{},
	)
}

func MigrateAll(conn *gorm.DB) error {
	return conn.AutoMigrate(
		&exchange.Transaction{},
		&exchange.Registration{},
		&exchange.User{},
		&exchange.Wallet{},
	)
}

func CreateConnection(conf config.Config, log *zap.Logger) (*gorm.DB, error) {
	var conn *gorm.DB
	err := fmt.Errorf("start")
	for err != nil {
		conn, err = gorm.Open(postgres.New(postgres.Config{
			DSN:                  fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", conf.Db.Host, conf.Db.User, conf.Db.Password, conf.Db.Database, conf.Db.Port),
			PreferSimpleProtocol: true,
		}), &gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
		})
		<-time.After(time.Second * 2)
	}
	return conn, nil
}
