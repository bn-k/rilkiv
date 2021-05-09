package db

import (
	"context"

	"github.com/bn-k/rilkiv/exchange"
	"github.com/google/uuid"
)

type registrations handler

func (w registrations) GetRegistrationByID(ctx context.Context, id uuid.UUID) (exchange.Registration, error) {
	res := exchange.Registration{Orm: exchange.Orm{ID: id}}
	err := w.gorm.First(&res).Error
	if err != nil {
		return exchange.Registration{}, err
	}

	return res, err
}

func (w registrations) CreateRegistration(ctx context.Context, reg exchange.Registration) (exchange.Registration, error) {
	err := w.gorm.Create(&reg).Error
	if err != nil {
		return exchange.Registration{}, err
	}
	return exchange.Registration{}, nil
}
