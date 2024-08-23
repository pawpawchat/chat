package app

import (
	"context"
	"log"
	"sync"

	"github.com/jmoiron/sqlx"
	"github.com/pawpawchat/chat/config"
	"github.com/pawpawchat/chat/internal/domain/model"
	"github.com/pawpawchat/chat/internal/domain/service"
	"github.com/pawpawchat/chat/internal/infrastructure/repository"

	_ "github.com/jackc/pgx/v5/stdlib"
	gs "github.com/pawpawchat/chat/internal/app/grpc_server"
	ws "github.com/pawpawchat/chat/internal/app/websocket_server"
)

func Run(ctx context.Context, config *config.Config) {
	var wg sync.WaitGroup
	msgChan := make(chan *model.Message, 10)

	wg.Add(1)
	go func() {
		defer wg.Done()
		service := newService(config, msgChan)
		gs.Run(ctx, service, msgChan, config.Env().GRPCServerAddr)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		service := newService(config, msgChan)
		ws.Run(ctx, service, msgChan, config.Env().WebSocketAddr)
	}()

	wg.Wait()
}

func newService(config *config.Config, msgChan chan *model.Message) *service.Service {
	db := newDbConn(config.Env().DbUrl)
	return service.New(repository.NewMessageRepository(db), repository.NewChatRepository(db), repository.NewMemberRepository(db), msgChan)
}

func newDbConn(url string) *sqlx.DB {
	db, err := sqlx.Connect("pgx", url)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
