package websocketserver

import (
	"log/slog"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/pawpawchat/chat/internal/domain/model"
	"golang.org/x/net/context"
)

type service interface {
	GetMember(context.Context, uint64, uint64) (*model.Member, error)
}

type webSocketServer struct {
	clients    map[uint64]*websocket.Conn
	msgChan    chan *model.Message
	upgrader   *websocket.Upgrader
	service    service
	httpServer *http.Server
}

func newWebSocketServer(service service, msgChan chan *model.Message, addr string) *webSocketServer {
	// var srv *webSocketServer
	srv := &webSocketServer{
		clients: make(map[uint64]*websocket.Conn),
		msgChan: msgChan,
		upgrader: &websocket.Upgrader{
			ReadBufferSize:  1024,         // 1  kb
			WriteBufferSize: 1024 * 8 * 8, // 64 kb
		},
		service: service,
		httpServer: &http.Server{
			Addr: addr,
		},
	}

	srv.httpServer.Handler = UpgradeHandler(srv)
	return srv
}

func Run(ctx context.Context, service service, msgChan chan *model.Message, Addr string) {
	srv := newWebSocketServer(service, msgChan, Addr)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		slog.Debug("running a websocket server", "addr", Addr)
		if err := srv.httpServer.ListenAndServe(); err != nil {
			slog.Error("http server", "error", err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()

		if err := srv.httpServer.Shutdown(ctx); err != nil {
			slog.Error("shutdown", "error", err)
		}
		slog.Debug("the websocket server has been gracefully shut down")
	}()

	wg.Wait()
}
