package exchange

import (
	"context"
	"github.com/google/uuid"
	"time"
)

type UnconfirmedUser interface {
}

type UserExchange interface {
	MakeTransaction(ctx context.Context, wid uuid.UUID, amount int64, dest string) (Transaction, error)
	UserTransactions(ctx context.Context, address string, date *Date, page uint) ([]Transaction, error)
	UserWallets(ctx context.Context) ([]Wallet, error)
}

type EmailClient interface {
	SendConfirmation(email, token string) error
}

type DB interface {
	Users
	Wallets
	Registrations
	Transactions
}

type Users interface {
	CreateUser(ctx context.Context, user User) (User, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	GetUserByEmailToken(ctx context.Context, email string, token string) (User, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (User, error)
	SetUserConfirmed(ctx context.Context, userID uuid.UUID) error
}

type Wallets interface {
	CreateWallet(ctx context.Context, wallet Wallet) (Wallet, error)
	GetWalletByID(ctx context.Context, id uuid.UUID) (Wallet, error)
	GetUserWallets(ctx context.Context, id uuid.UUID) ([]Wallet, error)
	GetUserWallet(ctx context.Context, uid, wid uuid.UUID) (Wallet, error)
}

type Registrations interface {
	CreateRegistration(ctx context.Context, reg Registration) (Registration, error)
	GetRegistrationByID(ctx context.Context, id uuid.UUID) (Registration, error)
}

type Transactions interface {
	CreateTransaction(ctx context.Context, trx Transaction) (Transaction, error)
	GetTransactionByID(ctx context.Context, id uuid.UUID) (Transaction, error)
	GetUserTransactions(ctx context.Context, userID uuid.UUID, address string, date *Date, page uint) ([]Transaction, error)
}

type Date struct {
	From time.Time
	To   time.Time
}
