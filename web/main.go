package main

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"time"
)

func main() {
	mux := http.NewServeMux()
	apiMux := http.NewServeMux()
	apiMux.HandleFunc("/endpoint/{$}", func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			slog.Warn("error reading body of request", "err", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.Write(fmt.Appendf(nil, "hello %s", body))
	})
	mux.Handle("/", http.FileServer(http.Dir("static")))
	mux.Handle("/api/", http.StripPrefix("/api", apiMux))
	srv := http.Server{
		Addr:         ":" + os.Getenv("PORT"),
		Handler:      mux,
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
