package db

import (
	"context"
	"github.com/bn-k/rilkiv/exchange"
	"github.com/google/uuid"
)

type users handler

func (u users) GetUserByEmail(ctx context.Context, email string) (exchange.User, error) {
	user := exchange.User{}
	result := u.gorm.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}

func (u users) GetUserByID(ctx context.Context, id uuid.UUID) (exchange.User, error) {
	res := exchange.User{Orm: exchange.Orm{ID: id}}
	err := u.gorm.First(&res).Error
	if err != nil {
		return exchange.User{}, err
	}

	return res, err
}

func (u users) CreateUser(ctx context.Context, user exchange.User) (exchange.User, error) {
	err := u.gorm.Create(&user).Error
	if err != nil {
		return exchange.User{}, err
	}
	return exchange.User{}, nil
}

//func (u users) SetUserInvitation(uid, token string) error {
//	return u.gorm.
//		Model(&exchange.User{Orm: exchange.Orm{ID: uid}}).
//		Update("invitation", token).Error
//}
