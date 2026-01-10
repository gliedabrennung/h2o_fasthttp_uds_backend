package uds

import (
	"log/slog"
	"net"
	"os"
)

func InitUDSListener() (net.Listener, string) {
	socketPath := "/tmp/fasthttp.sock"

	if err := os.RemoveAll(socketPath); err != nil {
		slog.Error("Error deleting socket", "error", err)
	}

	ln, err := net.Listen("unix", socketPath)
	if err != nil {
		slog.Error("Error listening socket", "error", err)
		os.Exit(1)
	}

	if err := os.Chmod(socketPath, 0777); err != nil {
		slog.Error("Error chmod", "error", err)
	}

	slog.Info("UDS Listener initialized", "path", socketPath)
	return ln, socketPath
}
