package resiliency

import (
	"context"
	"grpc-go-client/internal/port"
	pb "proto/protogen/go/resiliency"

	"google.golang.org/grpc"
)

type ResiliencyAdapter struct {
	resiliencyClient port.ResiliencyClientPort
}

type ResiliencyRequest struct{}
type ResiliencyResponse struct{}

func NewResiliencyAdapter(conn *grpc.ClientConn) (*ResiliencyAdapter, error) {
	client := pb.NewResiliencyServiceClient(conn)

	return &ResiliencyAdapter{                       
		resiliencyClient: client,
	}, nil
}

func (a *ResiliencyAdapter) GetResiliency(ctx context.Context, resiliencyRequest *ResiliencyRequest) (*pb.ResiliencyResponse, error) {
	resp, err := a.resiliencyClient.GetResiliency(ctx, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil   
}
