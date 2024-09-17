package server

import (
	"log/slog"
	"net/http"
	"sync"
	"time"

	"github.com/pawpawchat/chat/internal/domain/model"
	"golang.org/x/net/context"
)

func Run(ctx context.Context, provider chatMembersProvider, msgChan chan *model.Message, Addr string) {
	srv := newWebSocketServer(provider, msgChan, Addr)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		slog.Debug("chat: websocket server startup", "addr", Addr)
		if err := srv.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("chat: websocket server", "error", err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()

		if err := srv.httpServer.Shutdown(ctx); err != nil {
			slog.Error("chat: weboscket server shutdown", "error", err)
		}
		slog.Debug("chat: websocket server has been gracefully shutdown")
	}()

	wg.Wait()
}
