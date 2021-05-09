package account

import (
	"context"
	"github.com/bn-k/rilkiv/exchange"
	"github.com/go-chi/chi/v5"
	"net/http"

	"github.com/go-chi/render"
	"go.uber.org/zap"
)

type ConfirmResponse struct {
	Message string
}

func (b ConfirmResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type ConfirmRequest struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

func (h *Handlers) Confirm(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := new(ConfirmRequest)

		req.Email = chi.URLParam(r, "email")
		req.Token = chi.URLParam(r, "token")

		u, err := h.DB.GetUserByEmailToken(ctx, req.Email, req.Token)
		if err != nil {
			h.Log.Debug("not found", zap.Any("req", r.Body))
			w.WriteHeader(http.StatusNotFound)
			return
		}

		err = h.DB.SetUserConfirmed(ctx, u.ID)
		if err != nil {
			h.Log.Debug("not found", zap.Any("req", r.Body))
			w.WriteHeader(http.StatusNotFound)
			return
		}

		_, err = h.DB.CreateTransaction(ctx, exchange.Transaction{
			Orm:      exchange.Orm{},
			From:     orgAddressETH,
			To:       exchange.GetWalletCurrency(exchange.ETH, u.Wallets).Address,
			Currency: exchange.ETH,
			Amount:   initialBalance,
		})
		if err != nil {
			h.Log.Debug("cannot send initial ETH amount")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		_, err = h.DB.CreateTransaction(ctx, exchange.Transaction{
			Orm:      exchange.Orm{},
			From:     orgAddressBTC,
			To:       exchange.GetWalletCurrency(exchange.BTC, u.Wallets).Address,
			Currency: exchange.BTC,
			Amount:   initialBalance,
		})
		if err != nil {
			h.Log.Debug("cannot send initial ETH amount")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = render.Render(w, r, ConfirmResponse{
			Message: "success",
		})
		if err != nil {
			h.Log.Error("cannot render ", zap.Error(err))
		}
	}
}
