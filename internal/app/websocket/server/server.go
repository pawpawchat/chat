package server

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/pawpawchat/chat/internal/domain/model"
	"golang.org/x/net/context"
)

type chatMembersProvider interface {
	GetMember(context.Context, int64, int64) (*model.Member, error)
}

type webSocketServer struct {
	clients    map[int64]*websocket.Conn
	msgChan    chan *model.Message
	upgrader   *websocket.Upgrader
	provider   chatMembersProvider
	httpServer *http.Server
}

func newWebSocketServer(provider chatMembersProvider, msgChan chan *model.Message, addr string) *webSocketServer {
	srv := &webSocketServer{
		clients: make(map[int64]*websocket.Conn),
		msgChan: msgChan,
		upgrader: &websocket.Upgrader{
			ReadBufferSize:  1024,         // 1  kb
			WriteBufferSize: 1024 * 8 * 8, // 64 kb
		},
		provider: provider,
		httpServer: &http.Server{
			Addr: addr,
		},
	}

	srv.httpServer.Handler = UpgradeHandler(srv)
	return srv
}
