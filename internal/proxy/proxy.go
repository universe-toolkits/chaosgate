package proxy

import (
	"net/http"
	"sync"

	"github.com/universe-toolkits/chaosgate/internal/config"
	"github.com/universe-toolkits/chaosgate/internal/rules"
)

type Proxy struct {
	engineMu      sync.RWMutex
	engine        *rules.Engine
	transport     http.RoundTripper
	currentConfig *config.Config
}

func New(cfg *config.Config) *Proxy {
	transport := defaultTransport()
	engine := rules.NewEngine(buildRules(cfg))

	return &Proxy{
		engine:        engine,
		transport:     transport,
		currentConfig: cfg,
	}
}

func (p *Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodConnect {
		p.handleHTTPS(w, r)
		return
	}

	ctx := rules.NewContext(w, r, p.transport)
	p.engine.Execute(ctx)
}

func (p *Proxy) Config() *config.Config {
	p.engineMu.RLock()
	defer p.engineMu.RUnlock()
	return p.currentConfig
}

func (p *Proxy) UpdateConfig(cfg *config.Config) error {
	engine := rules.NewEngine(buildRules(cfg)) // new engine

	p.engineMu.Lock()
	defer p.engineMu.Unlock()

	p.currentConfig = cfg
	p.engine = engine
	return nil
}
