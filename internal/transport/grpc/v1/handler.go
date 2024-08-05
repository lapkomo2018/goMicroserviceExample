package v1

import (
	"context"
	grpcService "github.com/lapkomo2018/goServiceExample/pkg/grpc/example"
)

type (
	Handler struct {
		grpcService.UnimplementedExampleServer
	}
)

func New() *Handler {
	return &Handler{}
}

func (h *Handler) Authenticate(ctx context.Context, r *grpcService.DoSomeRequest) (*grpcService.DoSomeResponse, error) {
	return &grpcService.DoSomeResponse{Message: r.GetMessage()}, nil
}
