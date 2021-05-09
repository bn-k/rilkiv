package account

import (
	"context"
	"net/http"

	"github.com/go-chi/render"
	"go.uber.org/zap"
)

type ConfirmResponse struct {
}
func (b ConfirmResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
type ConfirmRequest struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
}
func (h *Handlers) Confirm(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := render.Render(w, r, ConfirmResponse{})
		if err != nil {
			h.Log.Error("cannot render ", zap.Error(err))
		}
	}
}

