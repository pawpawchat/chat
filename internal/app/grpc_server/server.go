package grpcserver

import (
	"context"
	"log"
	"log/slog"
	"net"
	"sync"

	"github.com/pawpawchat/chat/api/pb"
	"github.com/pawpawchat/chat/internal/domain/model"
	"github.com/pawpawchat/chat/internal/domain/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type _service interface {
	CreateChat(context.Context, *model.Member, *model.Chat) error
	GetChat(context.Context, uint64) (*model.Chat, error)

	AddMember(context.Context, *model.Member) error
	GetMembers(context.Context, uint64) (*[]model.Member, error)

	SendMessage(context.Context, *model.Message) error
	GetMessages(context.Context, uint64) (*[]model.Message, error)
}

var _ _service = (*service.Service)(nil)

type ChatGRPCServer struct {
	pb.UnimplementedChatServiceServer
	service _service
}

func NewGRPCServer(service _service) *ChatGRPCServer {
	return &ChatGRPCServer{service: service}
}

func Run(ctx context.Context, service _service, msgChan chan *model.Message, addr string) {
	grpcSrv := grpc.NewServer()
	chatSrv := NewGRPCServer(service)

	pb.RegisterChatServiceServer(grpcSrv, chatSrv)
	reflection.Register(grpcSrv)

	lsn, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		slog.Debug("running a grpc server", "addr", addr)
		if err := grpcSrv.Serve(lsn); err != nil {
			slog.Error("grpc server", "error", err)
		}
	}()

	wg.Add(1)
	go func() {
		wg.Done()
		<-ctx.Done()
		grpcSrv.GracefulStop()
		slog.Debug("the grpc server has been gracefully shut down")
	}()

	wg.Wait()
}
