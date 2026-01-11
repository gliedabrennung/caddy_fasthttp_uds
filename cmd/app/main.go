package main

import (
	"log/slog"
	"runtime"
	"wsp_go/cmd/uds"
	"wsp_go/internal"
	"wsp_go/internal/router"
	"wsp_go/internal/server"
)

func main() {
	runtime.GOMAXPROCS(2)
	r := router.InitRouter()
	listener, socketPath := uds.InitUDSListener()
	srv, err := server.StartServer(r, listener)
	if err != nil {
		slog.Error("Error starting server", "error", err.Error())
	}
	defer func() {
		internal.Shutdown(srv, socketPath)
	}()
}
