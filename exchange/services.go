package exchange

import (
	"context"
	"github.com/google/uuid"
)

type Exchange interface {
}

type EmailClient interface {
	SendConfirmation(email, token string) error
}

type DB interface {
	Users
}

type Users interface {
	CreateUser(ctx context.Context, user User) (User, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (User, error)
}
