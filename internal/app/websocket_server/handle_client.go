package websocketserver

import (
	"log/slog"

	"github.com/gorilla/websocket"
	"golang.org/x/net/context"
)

func (s *webSocketServer) handleClient(conn *websocket.Conn, clientID uint64) {
	defer conn.Close()
	closeCh := make(chan struct{})

	go func() {
		defer close(closeCh)
		for {
			if _, _, err := conn.ReadMessage(); err != nil {
				slog.Debug("client has disconnected", "id", clientID, "err", err)
				return
			}
		}
	}()

	for {
		select {
		case msg := <-s.msgChan:
			if _, err := s.service.GetMember(context.TODO(), msg.ChatID, clientID); err != nil {
				// TODO: redirect to notify service
				continue
			}

			go func() {
				if err := conn.WriteJSON(msg); err != nil {
					slog.Debug("failed to delivery the message", "err", err)
				}
			}()

		case <-closeCh:
			return
		}
	}
}
