package account

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/render"
	"go.uber.org/zap"
)

type LoginResponse struct {
	Bearer string `json:"bearer"`
}

func (b LoginResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handlers) Login(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := new(LoginRequest)
		err := json.NewDecoder(r.Body).Decode(req)
		if err != nil {
			h.Log.Debug("bad request", zap.Any("req", r.Body))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		u, err := h.DB.GetUserByEmail(ctx, req.Email)
		if err != nil {
			h.Log.Debug("unknown user", zap.Any("req", r.Body))
			w.WriteHeader(http.StatusForbidden)
			return
		}
		if !checkPasswordHash(req.Password, u.Password) {
			h.Log.Debug("wrong password", zap.Any("req", r.Body))
			w.WriteHeader(http.StatusForbidden)
			return
		}

		_, tks, err := tokenAuth.Encode(map[string]interface{}{"user_id": u.ID, "conf": u.Confirmed, "role": u.Role})
		if err != nil {
			h.Log.Error("cannot encode ", zap.Error(err))
		}

		err = render.Render(w, r, LoginResponse{Bearer: tks})
		if err != nil {
			h.Log.Error("cannot render ", zap.Error(err))
		}
	}
}
