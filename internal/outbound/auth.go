package outbound

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"

	api "art-item/internal/api/proto"
)

type AuthServiceClient struct {
	client api.AuthServiceClient
}

func NewAuthServiceClient(serverAddr string) AuthServiceClient {
	var conn *grpc.ClientConn

	if serverAddr == "" {
		log.Fatalf("Invalid server address")
	}

	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	client := api.NewAuthServiceClient(conn)

	return AuthServiceClient{client: client}
}

func (ac *AuthServiceClient) VerifyToken(ctx context.Context, token string, tokenType string) (*api.VerifyTokenResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	r, err := ac.client.VerifyToken(ctx, &api.VerifyTokenRequest{Token: token, TokenType: tokenType})
	if err != nil {
		return nil, err
	}

	return r, nil
}
