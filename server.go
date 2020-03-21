// +build !lambda

package main

import (
	"context"
	"net/http"
	"time"

	"github.com/go-training/line-login/config"

	"golang.org/x/sync/errgroup"
)

// RunHTTPServer provide run http or https protocol.
func RunHTTPServer(ctx context.Context) (err error) {
	conf := config.MustLoad()

	server := &http.Server{
		Addr:    ":" + conf.HTTP.Port,
		Handler: setupRouter(),
	}

	return startServer(ctx, server)
}

func listenAndServe(ctx context.Context, s *http.Server) error {
	var g errgroup.Group
	g.Go(func() error {
		select {
		case <-ctx.Done():
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			return s.Shutdown(ctx)
		}
	})
	g.Go(func() error {
		return s.ListenAndServe()
	})
	return g.Wait()
}

func startServer(ctx context.Context, s *http.Server) error {
	return listenAndServe(ctx, s)
}
