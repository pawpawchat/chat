package adapter

import (
	"context"

	"github.com/pawpawchat/chat/api/pb"
	"github.com/pawpawchat/chat/internal/domain/model"
)

type memberProvider interface {
	AddMember(context.Context, *model.Member) error
	GetMembers(context.Context, int64) (*[]model.Member, error)
}

func AddMemberAdapter(ctx context.Context, provider memberProvider, req *pb.AddMemberRequest) (*pb.AddMemberResponse, error) {
	member := &model.Member{
		MemberID: req.MemberId,
		Username: req.Username,
		ChatID:   req.ChatId,
		Role:     req.Role,
	}

	if err := provider.AddMember(ctx, member); err != nil {
		return nil, err
	}

	return &pb.AddMemberResponse{Member: member.ToPb()}, nil
}

func GetMembersAdapater(ctx context.Context, provider memberProvider, req *pb.GetMembersRequest) (*pb.GetMembersResponse, error) {
	chatID := req.ChatId

	members, err := provider.GetMembers(ctx, chatID)
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
