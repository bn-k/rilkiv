package account

import (
	"context"
	"fmt"

	"github.com/bn-k/rilkiv/exchange"
	"github.com/google/uuid"
)

type userExchange struct {
	user exchange.User
	h    Handlers
}

func (c userExchange) UserWallets(ctx context.Context) ([]exchange.Wallet, error) {
	return c.h.DB.GetUserWallets(ctx, c.user.ID)
}

type ErrInsufficientFund string

func (e ErrInsufficientFund) Error() string {
	return "not enough fund: " + string(e)
}

func (c userExchange) MakeTransaction(ctx context.Context, wid uuid.UUID, amount int64, dest string) (exchange.Transaction, error) {
	var res exchange.Transaction
	wallet, err := c.h.DB.GetUserWallet(ctx, c.user.ID, wid)
	if err != nil {
		return res, err
	}

	balance := wallet.GetBalance()
	if amount > balance {
		return res, ErrInsufficientFund(fmt.Sprintf("current balance: %d", balance))
	}

	res, err = c.h.DB.CreateTransaction(ctx, exchange.Transaction{
		From:       wallet.Address,
		To:         dest,
		Currency:   wallet.Currency,
		Amount:     amount,
		Commission: 0,
		Status:     exchange.Pending,
	})

	return res, err
}

func (c userExchange) UserTransactions(ctx context.Context, address string, date *exchange.Date, page uint) ([]exchange.Transaction, error) {
	return c.h.DB.GetUserTransactions(ctx, c.user.ID, address, date, page)
}

func (h Handlers) provideUserConfirmed(ctx context.Context, claims map[string]interface{}) (exchange.UserExchange, error) {
	res := userExchange{
		h: h,
	}
	userId, ok := claims["user_id"].(string)
	if !ok {
		return nil, fmt.Errorf("user_id not found")
	}

	uid, err := uuid.Parse(userId)
	if err != nil {
		return nil, err
	}

	res.user, err = h.DB.GetUserByID(ctx, uid)
	if err != nil {
		return nil, err
	}

	if !res.user.Confirmed {
		return nil, fmt.Errorf("user_id not found")
	}

	return res, nil
}
