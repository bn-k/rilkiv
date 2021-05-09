package account

import (
	"context"
	"github.com/go-chi/jwtauth/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"net/http"
)

type Wallet struct {
	ID      uuid.UUID
	Address string
	Balance int64
}
type GetWalletsResponse struct {
	Wallets []Wallet
}

func (b GetWalletsResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (h *Handlers) GetWallets(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, claims, _ := jwtauth.FromContext(r.Context())
		confirmed, err := h.provideUserConfirmed(ctx, claims)
		if err != nil {
			h.Log.Debug("unauthorized", zap.Error(err))
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		ws, err := confirmed.UserWallets(ctx)
		if err != nil {
			h.Log.Debug("user wallets fail", zap.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		var wallets []Wallet
		for _, w := range ws {
			wallets = append(wallets, Wallet{
				ID:      w.ID,
				Address: w.Address,
				Balance: w.GetBalance(),
			})
		}
		err = render.Render(w, r, GetWalletsResponse{Wallets: wallets})
		if err != nil {
			h.Log.Error("cannot render ", zap.Error(err))
		}
	}
}
