package router

import (
	"github.com/go-chi/chi/v5"
)

func (p *routes) apiHandlers() chi.Router {
	r := chi.NewRouter()
	r.Post("/login", p.login) // POST /users - create a new user and persist it
	// r.Get("/stocks", p.stocks)
	// r.Get("/trades", p.trades)
	// r.Post("/profile", p.profile)
	// r.Post("/stop", p.stop)
	// r.Post("/margins", p.margin)
	// r.Post("/full", p.full)
	// r.Post("/transactions", p.transactions)

	// r.Get("/market-status", p.marketStauts)
	// r.Get("/holiday", p.holiday)
	// r.Get("/loosers", p.loosers)
	// r.Get("/gainers", p.gainers)
	// r.Get("/option-chain", p.optionChain)

	return r
}
