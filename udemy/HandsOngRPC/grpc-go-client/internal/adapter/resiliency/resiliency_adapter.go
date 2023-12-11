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

type ResiliencyRequest struct {
	MaxDelaySecond int32
	MinDelaySecond int32
	StatusCodes    []uint32
}

type ResiliencyResponse struct {
	Response   string
	StatusCode uint32
}

func NewResiliencyAdapter(conn *grpc.ClientConn) (*ResiliencyAdapter, error) {
	client := pb.NewResiliencyServiceClient(conn)

	return &ResiliencyAdapter{
		resiliencyClient: client,
	}, nil
}

func (a *ResiliencyAdapter) GetResiliency(ctx context.Context, resiliencyRequest *ResiliencyRequest) (*ResiliencyResponse, error) {
	req := &pb.ResiliencyRequest{
		MaxDelaySecond: resiliencyRequest.MaxDelaySecond,
		MinDelaySecond: resiliencyRequest.MinDelaySecond,
		StatusCodes:    resiliencyRequest.StatusCodes,
	}
	resp, err := a.resiliencyClient.GetResiliency(ctx, req)
	if err != nil {
		return nil, err
	}
	return &ResiliencyResponse{
		Response:   resp.Response,
		StatusCode: resp.StatusCode,
	}, nil
}
