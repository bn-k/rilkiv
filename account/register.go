package account

import (
	"context"
	"encoding/json"
	"github.com/go-chi/render"
	"go.uber.org/zap"
	"net/http"
)

type RegisterResponse struct {
}

func (b RegisterResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type RegisterRequest struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	Lastname  string `json:"lastname"`
	Firstname string `json:"firstname"`
}

func (h *Handlers) Register(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := new(RegisterRequest)
		err := json.NewDecoder(r.Body).Decode(req)
		if err != nil {
			h.Log.Debug("bad request", zap.Any("req", r.Body))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		user, err := h.fmtRegister(ctx, UserRegister{
			Email:     req.Email,
			FirstName: req.Firstname,
			LastName:  req.Lastname,
			Password:  req.Password,
		})
		if err != nil {
			h.Log.Debug("cannot register user", zap.Error(err))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		_, err = h.DB.CreateUser(ctx, user)
		if err != nil {
			h.Log.Debug("cannot register user", zap.Error(err))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = h.Mail.SendConfirmation(user.Email, user.ConfirmToken)
		if err != nil {
			h.Log.Debug("cannot sent confirm email", zap.Error(err))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = render.Render(w, r, RegisterResponse{})
		if err != nil {
			h.Log.Error("cannot render ", zap.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
