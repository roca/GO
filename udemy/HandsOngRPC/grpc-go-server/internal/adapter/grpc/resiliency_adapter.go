package grpc

import (
	"context"
	pb "proto/protogen/go/resiliency"
)

func (a *GrpcAdapter) GetResiliency(context.Context, *pb.ResiliencyRequest) (*pb.ResiliencyResponse, error) {
	_, err := a.ResiliencyService.GetResiliency()

	resp := &pb.ResiliencyResponse{}
	return resp, err
}
func (a *GrpcAdapter) GetResiliencyStream(*pb.ResiliencyRequest, pb.ResiliencyService_GetResiliencyStreamServer) error {
	return a.ResiliencyService.GetResiliencyStream()
}
func (a *GrpcAdapter) SendResiliencyStream(pb.ResiliencyService_SendResiliencyStreamServer) error {
	return a.ResiliencyService.SendResiliencyStream()
}
func (a *GrpcAdapter) BidirectionalResiliencyStream(pb.ResiliencyService_BidirectionalResiliencyStreamServer) error {
	return a.ResiliencyService.BidirectionalResiliencyStream()
}
