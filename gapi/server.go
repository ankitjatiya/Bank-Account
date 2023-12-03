package gapi

import (
	db "Bank-Account/db/sqlc"
	"Bank-Account/pb"
	"Bank-Account/token"
	"Bank-Account/util"
	"fmt"
)

// Server serves gRPC request for our banking service
type Server struct {
	pb.UnimplementedBankAccountServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	return server, nil
}
