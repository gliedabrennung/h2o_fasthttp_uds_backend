package server

import (
	"log/slog"
	"net"
	"time"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

func StartServer(router *router.Router, ln net.Listener) (*fasthttp.Server, error) {
	server := &fasthttp.Server{
		Handler: router.Handler,
		Name:    "caddy_uds_fasthttp",

		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,

		CloseOnShutdown: true,

		MaxRequestBodySize: 4 * 1024 * 1024,
	}

	slog.Info("Starting server on Unix Socket", "path", ln.Addr().String())

	go func() {
		if err := server.Serve(ln); err != nil {
			slog.Error("Server failed", "error", err)
		}
	}()
	return server, nil
}
