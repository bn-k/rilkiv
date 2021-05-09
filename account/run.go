package account

import (
	"context"
	"fmt"
	"github.com/bn-k/rilkiv/app"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/docgen"
	"github.com/go-chi/jwtauth/v5"
	"github.com/go-chi/render"
	"net/http"
)

type Handlers app.App

var tokenAuth *jwtauth.JWTAuth

const (
	orgAddressBTC = "1FfmbHfnpaZjKFvyi1okTjJJusN455paPH"
	orgAddressETH = "0x29D7d1dd5B6f9C864d9db560D72a247c178aE86B"
)

func Server(ap app.App, doc bool) error {
	ctx := context.Background()
	tokenAuth = jwtauth.New("HS256", []byte(ap.Conf.Auth.Secret), nil)
	ap.Log.Info("token auth init")
	h := Handlers(ap)
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Post("/login", h.Login(ctx))
	r.Post("/register", h.Register(ctx))
	r.Get("/confirm/{email}/{token}", h.Confirm(ctx))

	// Protected routes
	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator)

		r.Get("/transactions", h.GetTransactions(ctx))
		r.Post("/transaction", h.MakeTransaction(ctx))

		r.Get("/wallets", h.GetWallets(ctx))
	})

	if doc {
		fmt.Println(docgen.MarkdownRoutesDoc(r, docgen.MarkdownOpts{
			ProjectPath: "doc",
			Intro:       "Short documentation ",
		}))
		return nil
	}

	ap.Log.Info("listen :8080")
	return http.ListenAndServe(":8080", r)
}
