package db

import (
	"context"
	"fmt"
	"gorm.io/gorm/clause"

	"github.com/bn-k/rilkiv/exchange"
	"github.com/google/uuid"
)

type wallets handler

func (w wallets) GetWalletByID(ctx context.Context, id uuid.UUID) (exchange.Wallet, error) {
	res := exchange.Wallet{Orm: exchange.Orm{ID: id}}
	err := w.gorm.First(&res).Error
	if err != nil {
		return exchange.Wallet{}, err
	}

	return res, err
}

func (w wallets) GetUserWallet(ctx context.Context, uid, wid uuid.UUID) (exchange.Wallet, error) {
	wallets, err := w.GetUserWallets(ctx, uid)
	if err != nil {
		return exchange.Wallet{}, err
	}

	for _, wallet := range wallets {
		if wallet.ID == wid {
			return wallet, nil
		}
	}

	return exchange.Wallet{}, fmt.Errorf("not found")
}

func (w wallets) GetUserWallets(ctx context.Context, id uuid.UUID) ([]exchange.Wallet, error) {
	var res []exchange.Wallet
	err := w.gorm.
		Where("user_id = ?", id).
		Preload(clause.Associations).
		Find(&res).Error
	if err != nil {
		return nil, err
	}

	return res, err
}

func (w wallets) CreateWallet(ctx context.Context, user exchange.Wallet) (exchange.Wallet, error) {
	err := w.gorm.Create(&user).Error
	if err != nil {
		return exchange.Wallet{}, err
	}
	return exchange.Wallet{}, nil
}
