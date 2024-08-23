package grpcserver

import (
	"context"

	"github.com/pawpawchat/chat/api/pb"
)

func (s *ChatGRPCServer) GetMembers(ctx context.Context, req *pb.GetMembersRequest) (*pb.GetMembersResponse, error) {
	chatID := req.ChatId

	members, err := s.usecase.GetMembers(ctx, chatID)
	if err != nil {
		return nil, err
	}

	pbmembers := make([]*pb.Member, 0)

	for _, m := range *members {
		pbmembers = append(pbmembers, &pb.Member{
			MemberId: m.MemberID,
			Username: m.Username,
			Role:     m.Role,
		})
	}

	return &pb.GetMembersResponse{Members: pbmembers, ChatId: chatID}, nil
}
