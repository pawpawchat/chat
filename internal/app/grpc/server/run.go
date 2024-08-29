package server

import (
	"context"
	"log"
	"log/slog"
	"net"
	"sync"

	"github.com/pawpawchat/chat/api/pb"
	"github.com/pawpawchat/chat/internal/domain/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func Run(ctx context.Context, service profileService, msgChan chan *model.Message, addr string) {
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
