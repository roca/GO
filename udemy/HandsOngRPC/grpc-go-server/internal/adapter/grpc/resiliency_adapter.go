package grpc

import (
	"context"
	"log"
	pb "proto/protogen/go/resiliency"
)

func (a *GrpcAdapter) GetResiliency(ctx context.Context, req *pb.ResiliencyRequest) (*pb.ResiliencyResponse, error) {
	_, err := a.ResiliencyService.GetResiliency()

	resp := &pb.ResiliencyResponse{}
	return resp, err
}
func (a *GrpcAdapter) GetResiliencyStream(req *pb.ResiliencyRequest, stream pb.ResiliencyService_GetResiliencyStreamServer) error {
	context := stream.Context()
	for {
		select {
		case <-context.Done():
			log.Println("Client has cancelled stream")
			return nil
		default:
			_, err := a.ResiliencyService.GetResiliencyStream()
			if err != nil {
				//log.Println("Error getting current balance: %s\n", s.Err())
				return err
			}
			stream.Send(&pb.ResiliencyResponse{})
		}
	}
}
func (a *GrpcAdapter) SendResiliencyStream(stream pb.ResiliencyService_SendResiliencyStreamServer) error {
	return a.ResiliencyService.SendResiliencyStream()
}
func (a *GrpcAdapter) BidirectionalResiliencyStream(stream pb.ResiliencyService_BidirectionalResiliencyStreamServer) error {
	return a.ResiliencyService.BidirectionalResiliencyStream()
}
