package account

import (
	"context"
	"encoding/json"
	"github.com/bn-k/rilkiv/exchange"
	"github.com/go-chi/jwtauth/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"time"
)

type MakeTransactionRequest struct {
	FromWalletID string `json:"from_wallet_id"`
	ToAddress    string `json:"to_address"`
	Amount       int64  `json:"amount"`
}
type MakeTransactionResponse struct {
	Transaction exchange.Transaction
}

func (b MakeTransactionResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (h *Handlers) MakeTransaction(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, claims, _ := jwtauth.FromContext(r.Context())
		confirmed, err := h.provideUserConfirmed(ctx, claims)
		if err != nil {
			h.Log.Debug("unauthorized", zap.Error(err))
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		req := new(MakeTransactionRequest)
		err = json.NewDecoder(r.Body).Decode(req)
		if err != nil {
			h.Log.Debug("bad request", zap.Any("req", r.Body))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		wid, err := uuid.Parse(req.FromWalletID)
		if err != nil {
			h.Log.Debug("wrong wallet id", zap.Error(err))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		trx, err := confirmed.MakeTransaction(ctx, wid, req.Amount, req.ToAddress)
		if err != nil {
			h.Log.Debug("cannot make transaction", zap.Error(err))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = render.Render(w, r, MakeTransactionResponse{Transaction: trx})
		if err != nil {
			h.Log.Error("cannot render ", zap.Error(err))
		}
	}
}

type GetTransactionsResponse struct {
	Transactions []exchange.Transaction
}

func (b GetTransactionsResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (h *Handlers) GetTransactions(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, claims, _ := jwtauth.FromContext(r.Context())
		confirmed, err := h.provideUserConfirmed(ctx, claims)
		if err != nil {
			h.Log.Debug("unauthorized", zap.Error(err))
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		address := r.URL.Query().Get("address")
		fromRAW := r.URL.Query().Get("from")
		toRAW := r.URL.Query().Get("to")
		pageRAW := r.URL.Query().Get("page")

		var date *exchange.Date
		if fromRAW != "" && toRAW != "" {
			from, err := time.Parse("2006-01-02", fromRAW)
			if err != nil {
				h.Log.Debug("cannot parse `from` date", zap.Error(err), zap.String("query", fromRAW))
			}
			to, err := time.Parse("2006-01-02", toRAW)
			if err != nil {
				h.Log.Debug("cannot parse `to` date", zap.Error(err), zap.String("query", toRAW))
			}
			if !from.IsZero() && !to.IsZero() {
				date = &exchange.Date{
					From: from,
					To:   to,
				}
				h.Log.Debug("get transaction for date", zap.Any("date", date))
			}
		}
		page, _ := strconv.Atoi(pageRAW)
		trxs, err := confirmed.UserTransactions(ctx, address, date, uint(page))
		if err != nil {
			h.Log.Debug("user transaction fail", zap.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = render.Render(w, r, GetTransactionsResponse{Transactions: trxs})
		if err != nil {
			h.Log.Error("cannot render ", zap.Error(err))
		}
	}
}
