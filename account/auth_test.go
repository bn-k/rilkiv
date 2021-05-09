package account

import (
	"context"
	"github.com/bn-k/rilkiv/config"
	"github.com/bn-k/rilkiv/exchange"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestHandlers_fmtRegister(t *testing.T) {
	validate = validator.New()
	ctx := context.Background()
	h := Handlers{
		Conf: config.Config{
			Auth:             config.Auth{
				Secret:     "secret",
				ExpireTime: time.Hour,
				Encoder:    "encoder",
			},
		},
	}
	type args struct {
		req UserRegister
	}
	tests := []struct {
		name    string
		h       Handlers
		args    args
		want    exchange.User
		wantErr bool
	}{
		{
			name:    "Regular",
			h:       h,
			args:    args{
				req: UserRegister{
					Email:     "email@gmail.com",
					FirstName: "ngẫunhiên",
					LastName:  "ランダム",
					Password:  "asdfasdf",
				},
			},
			want:    exchange.User{
				Orm:       exchange.Orm{},
				Email:     "email@gmail.com",
				Firstname: "ngẫunhiên",
				Lastname:  "ランダム",
				Auth:      exchange.Auth{
					ConfirmToken: "ZW1haWxAZ21haWwuY29tc2VjcmV0",
					Confirmed:    false,
				},
			},
			wantErr: false,
		},
		{
			name:    "ShortPassword",
			h:       h,
			args:    args{
				req: UserRegister{
					Email:     "email@gmail.com",
					FirstName: "firstname",
					LastName:  "lastname",
					Password:  "asdfasd",
				},
			},
			want:    exchange.User{},
			wantErr: true,
		},
		{
			name:    "LongPassword",
			h:       h,
			args:    args{
				req: UserRegister{
					Email:     "email@gmail.com",
					FirstName: "firstname",
					LastName:  "lastname",
					Password:  "mqyvhrgbznpixtxexdsmorbxilkknrcyhiflgbypadsliylokrz",
				},
			},
			want:    exchange.User{},
			wantErr: true,
		},
		{
			name:    "WrongEmail",
			h:       h,
			args:    args{
				req: UserRegister{
					Email:     "emailgmail.com",
					FirstName: "firstname",
					LastName:  "lastname",
					Password:  "asdfasdf",
				},
			},
			want:    exchange.User{},
			wantErr: true,
		},
		{
			name:    "WrongEmailDomain",
			h:       h,
			args:    args{
				req: UserRegister{
					Email:     "email@asdfbhbkf.com",
					FirstName: "firstname",
					LastName:  "lastname",
					Password:  "asdfasdf",
				},
			},
			want:    exchange.User{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.h.fmtRegister(ctx, tt.args.req)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want.Email, got.Email)
				assert.Equal(t, tt.want.Lastname, got.Lastname)
				assert.Equal(t, tt.want.Firstname, got.Firstname)
				assert.Equal(t, tt.want.ConfirmToken, got.ConfirmToken)
				assert.Equal(t, tt.want.Confirmed, got.Confirmed)
			}
		})
	}
}
