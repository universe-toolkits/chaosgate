package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func (a *App) shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := a.apiServer.Shutdown(ctx); err != nil {
		return err
	}
	return a.proxyServer.Shutdown(ctx)
}

func (a *App) waitForShutdown() error {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop
	return a.shutdown()
}
