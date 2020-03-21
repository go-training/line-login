// +build lambda

package main

import (
	"context"

	"github.com/go-training/line-login/config"

	"github.com/apex/gateway"
)

// RunHTTPServer provide run http or https protocol.
func RunHTTPServer(ctx context.Context) error {
	conf := config.MustLoad()

	return gateway.ListenAndServe(":"+conf.HTTP.Port, setupRouter())
}
