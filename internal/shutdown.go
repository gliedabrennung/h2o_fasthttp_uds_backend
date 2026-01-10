package internal

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/valyala/fasthttp"
)

func Shutdown(server *fasthttp.Server, socketPath string) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	sig := <-stop
	slog.Info("Received signal, starting shutdown...", "signal", sig.String())

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.ShutdownWithContext(ctx); err != nil {
		slog.Error("Fasthttp shutdown error", "error", err)
	}
	time.Sleep(100 * time.Millisecond)
	if err := os.Remove(socketPath); err != nil {
		if !os.IsNotExist(err) {
			slog.Error("Failed to remove socket file", "error", err)
		}
	} else {
		slog.Info("Socket file removed successfully")
	}

	slog.Info("Socket file removed successfully", "path", socketPath)
}
