package exchange

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	Orm
	Email        string `gorm:"uniqueIndex"`
	Firstname    string
	Lastname     string
	Registration Registration
	Wallets      []Wallet
	Role         Role
	Auth
}

type Wallet struct {
	Orm
	Address  string
	Currency Currency `gorm:"not null"`
	UserID   uuid.UUID
	Received []Transaction `gorm:"foreignKey:To;references:Address"`
	Emitted  []Transaction `gorm:"foreignKey:From;references:Address"`
}

type Transaction struct {
	Orm
	From       string `gorm:"type:text"`
	To         string `gorm:"type:text"`
	Currency   Currency
	Amount     int64
	Commission int
	Status     TransactionStatus
}

type Registration struct {
	Orm
	IPAddress string
	UserAgent string
	UserID    uuid.UUID
}

type Orm struct {
	ID        uuid.UUID      `json:"id" gorm:"primarykey;type:uuid;default:gen_random_uuid()"`
	CreatedAt time.Time      `json:"created_at" gorm:"type:timestamp"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"type:timestamp"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index,type:timestamp"`
}

type Auth struct {
	Password     string
	ConfirmToken string
	Confirmed    bool
}

type (
	Currency          string
	Role              string
	TransactionStatus string
)

const (
	BTC = Currency("BTC")
	ETH = Currency("ETH")

	Client = Role("client")
	Admin  = Role("admin")

	Pending = TransactionStatus("pending")
	Done    = TransactionStatus("done")
)

func GetWalletCurrency(c Currency, wallets []Wallet) Wallet {
	for _, w := range wallets {
		if w.Currency == c {
			return w
		}
	}

	return Wallet{}
}

func (u User) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (c Currency) String() string {
	return string(c)
}
func (w Wallet) GetBalance() int64 {
	var emitted int64
	var received int64
	for _, transaction := range w.Emitted {
		emitted += transaction.Amount
	}
	for _, transaction := range w.Received {
		received += transaction.Amount
	}

	return received - emitted
}
