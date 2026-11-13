package main

import (
	"log/slog"
	"net/http"
	"time"
)

func main() {
	srv := http.Server{
		Addr:         ":6767",
		Handler:      http.FileServer(http.Dir("static")),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  20 * time.Second,
	}
	slog.Info("started running server")
	err := srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		slog.Error("server unexpectedly ended", "err", err)
	}
}
