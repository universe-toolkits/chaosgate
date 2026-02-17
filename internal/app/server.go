package app

import (
	"net/http"

	"github.com/universe-toolkits/chaosgate/internal/config"
	"github.com/universe-toolkits/chaosgate/internal/proxy"
	"github.com/universe-toolkits/chaosgate/internal/web"
)

type App struct {
	proxyServer *http.Server
	apiServer   *http.Server
}

func New() (*App, error) {
	cfg, err := config.Load("configs/sample.yaml")
	if err != nil {
		return nil, err
	}

	proxyEngine := proxy.New(cfg)

	apiHandler := web.NewAPI(proxyEngine)

	return &App{
		proxyServer: &http.Server{
			Addr:    ":8080",
			Handler: proxyEngine,
		},
		apiServer: &http.Server{
			Addr:    ":9090",
			Handler: apiHandler.Handler(),
		},
	}, nil
}

func (a *App) Run() error {
	go a.apiServer.ListenAndServe()
	return a.proxyServer.ListenAndServe()
}
