package db

import (
	"context"
	"github.com/bn-k/rilkiv/exchange"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type transactions handler

func (w transactions) GetTransactionByID(ctx context.Context, id uuid.UUID) (exchange.Transaction, error) {
	res := exchange.Transaction{Orm: exchange.Orm{ID: id}}
	err := w.gorm.First(&res).Error
	if err != nil {
		return exchange.Transaction{}, err
	}

	return res, err
}

func (w transactions) GetUserTransactions(ctx context.Context, userID uuid.UUID, address string, date *exchange.Date, page uint) ([]exchange.Transaction, error) {
	if page == 0 {
		page++
	}
	const pageSize = 100
	var res []exchange.Transaction
	sql := w.gorm.
		Scopes(func(d *gorm.DB) *gorm.DB {
			offset := (page - 1) * pageSize
			return d.Offset(int(offset)).Limit(pageSize)
		}).
		Order("created_at").
		Joins("JOIN wallets ON wallets.address = transactions.from OR wallets.address = transactions.to AND wallets.user_id = ?", userID)

	if address != "" {
		sql = sql.Where("transactions.from = ? OR transactions.to = ? ", address, address)
	}

	if date != nil {
		sql = sql.Where("transactions.created_at BETWEEN ? AND ? ", date.From, date.To)
	}

	err := sql.Find(&res).Error

	return res, err
}

func (w transactions) CreateTransaction(ctx context.Context, trx exchange.Transaction) (exchange.Transaction, error) {
	err := w.gorm.Create(&trx).Error
	if err != nil {
		return exchange.Transaction{}, err
	}
	return trx, nil
}
