package web

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/universe-toolkits/chaosgate/internal/config"
	"github.com/universe-toolkits/chaosgate/internal/proxy"
)

type API struct {
	proxy *proxy.Proxy
	mu    sync.RWMutex
}

func NewAPI(p *proxy.Proxy) *API {
	return &API{proxy: p}
}

func (a *API) Handler() http.Handler {
	r := chi.NewRouter()

	r.Get("/api/config", a.getConfig)
	r.Put("/api/config", a.updateConfig)
	//r.Post("/reload", a.reloadConfig)

	r.Get("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	return r
}

func (a *API) getConfig(w http.ResponseWriter, r *http.Request) {
	a.mu.RLock()
	defer a.mu.RUnlock()

	json.NewEncoder(w).Encode(a.proxy.Config())
}

func (a *API) updateConfig(w http.ResponseWriter, r *http.Request) {
	a.mu.Lock()
	defer a.mu.Unlock()

	var cfg config.Config
	if err := json.NewDecoder(r.Body).Decode(&cfg); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := a.proxy.UpdateConfig(&cfg); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
