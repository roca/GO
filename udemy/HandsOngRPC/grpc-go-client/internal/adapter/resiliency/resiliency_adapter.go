package resiliency

import (
	"grpc-go-client/internal/port"
	pb "proto/protogen/go/resiliency"

	"google.golang.org/grpc"
)

type ResiliencyAdapter struct {
	resiliencyClient port.ResiliencyClientPort
}

func NewResiliencyAdapter(conn *grpc.ClientConn) (*ResiliencyAdapter, error) {
	client := pb.NewResiliencyServiceClient(conn)

	return &ResiliencyAdapter{
		resiliencyClient: client,
	}, nil
}

// func (a *ResiliencyAdapter) GetResiliency(ctx context.Context, resiliencyResuest interface{}) (interface{}, error) {
// 	return a.resiliencyClient.GetResiliency(nil, nil)
// }

