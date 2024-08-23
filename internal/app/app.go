package app

import (
	"context"
	"log"
	"log/slog"
	"net"
	"sync"

	"github.com/jmoiron/sqlx"
	"github.com/pawpawchat/chat/api/pb"
	"github.com/pawpawchat/chat/config"
	"github.com/pawpawchat/chat/internal/domain/service"
	"github.com/pawpawchat/chat/internal/infrastructure/repository"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/jackc/pgx/v5/stdlib"
	grpcserver "github.com/pawpawchat/chat/internal/app/grpc_server"
)

func Run(ctx context.Context, config *config.Config) error {
	l, err := net.Listen("tcp", config.Env().GRPC_SERVER_ADDR)
	if err != nil {
		return err
	}

	srv := newGRPCServer(config)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		slog.Debug("grpc server is up and running", "addr", config.Env().GRPC_SERVER_ADDR)
		err = srv.Serve(l)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()
		srv.GracefulStop()
		slog.Debug("grpc server was gracefuly stopped")
	}()

	wg.Wait()
	return err
}

func newGRPCServer(config *config.Config) *grpc.Server {
	db, err := sqlx.Connect("pgx", config.Env().DB_URL)
	if err != nil {
		log.Fatal(err)
	}

	ps := service.New(repository.NewMessageRepository(db), repository.NewChatRepository(db), repository.NewMemberRepository(db))

	chatServer := grpcserver.New(ps)

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	pb.RegisterChatServiceServer(grpcServer, chatServer)
	return grpcServer
}
