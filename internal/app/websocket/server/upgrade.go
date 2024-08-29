package server

import (
	"log/slog"
	"net/http"
	"strconv"
)

func UpgradeHandler(s *webSocketServer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(r.Header.Get("id"), 0, 10)
		if err != nil {
			http.Error(w, "missing/invalid id", http.StatusBadRequest)
			return
		}

		conn, err := s.upgrader.Upgrade(w, r, nil)
		if err != nil {
			slog.Debug("upgrade:", "err", err)
			return
		}

		slog.Debug("new connection", "id", id)
		go s.handleClient(conn, id)
	}
}
