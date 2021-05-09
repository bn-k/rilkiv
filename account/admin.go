package account

import (
	"context"
	"encoding/json"
	"github.com/go-chi/jwtauth/v5"
	"net/http"

	"github.com/go-chi/render"
	"go.uber.org/zap"
)

type AdminResponse struct {
}

func (b AdminResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type AdminRequest struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

func (h *Handlers) Admin(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, claims, _ := jwtauth.FromContext(r.Context())
		if claims["role"] != "admin" {
			h.Log.Debug("unauthorized")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		req := new(AdminRequest)
		err := json.NewDecoder(r.Body).Decode(req)
		if err != nil {
			h.Log.Debug("bad request", zap.Any("req", r.Body))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		u, err := h.DB.GetUserByEmail(ctx, req.Email)
		if err != nil {
			h.Log.Debug("not found", zap.Any("req", r.Body))
			w.WriteHeader(http.StatusNotFound)
			return
		}

		err = render.Render(w, r, u)
		if err != nil {
			h.Log.Error("cannot render ", zap.Error(err))
		}
	}
}
