package grpcserver

import (
	"context"

	"github.com/pawpawchat/chat/api/pb"
	"github.com/pawpawchat/chat/internal/domain/model"
)

func (s *ChatGRPCServer) AddMember(ctx context.Context, req *pb.AddMemberRequest) (*pb.AddMemberResponse, error) {
	user, err := parseAddMemberRequest(req)
	if err != nil {
		return nil, err
	}

	if err := s.service.AddMember(ctx, user); err != nil {
		return nil, err
	}

	pbmember, err := parseMember(user)
	if err != nil {
		return nil, err
	}

	return &pb.AddMemberResponse{Member: pbmember}, nil
}

func parseMember(m *model.Member) (*pb.Member, error) {
	return &pb.Member{
		MemberId: m.MemberID,
		Username: m.Username,
		ChatId:   m.ChatID,
		Role:     m.Role,
	}, nil
}

func parseAddMemberRequest(pb *pb.AddMemberRequest) (*model.Member, error) {
	invited := &model.Member{
		MemberID: pb.MemberId,
		Username: pb.Username,
		ChatID:   pb.ChatId,
		Role:     pb.Role,
	}

	return invited, nil
}
