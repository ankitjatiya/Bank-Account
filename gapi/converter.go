package gapi

import (
	db "Bank-Account/db/sqlc"
	"Bank-Account/pb"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertUser(user db.User) *pb.User {
	return &pb.User{
		Username:          user.Username,
		FullName:          user.FullName,
		Email:             user.Email,
		PasswordChangedAt: timestamppb.New(user.PasswordChanged),
		CreatedAt:         timestamppb.New(user.CreatedAt),
	}
}
