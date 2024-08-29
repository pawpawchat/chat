package server

import (
	"log/slog"
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
