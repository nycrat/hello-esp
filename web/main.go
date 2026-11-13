package main

import (
	"log/slog"
	"net/http"
)

func main() {
	slog.Info("started running server")
	http.ListenAndServe(":6767", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("hello"))
		if err != nil {
			slog.Warn("error writing response", "err", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}))
}
