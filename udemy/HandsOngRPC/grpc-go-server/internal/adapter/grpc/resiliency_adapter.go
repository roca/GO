package grpc

import (
	"context"
	"io"
	"log"
	pb "proto/protogen/go/resiliency"
)

func (a *GrpcAdapter) GetResiliency(ctx context.Context, req *pb.ResiliencyRequest) (*pb.ResiliencyResponse, error) {
	resp, err := a.ResiliencyService.GetResiliency(req.MaxDelaySecond, req.MinDelaySecond)

	return &pb.ResiliencyResponse{Response: resp}, err
}
func (a *GrpcAdapter) GetResiliencyStream(req *pb.ResiliencyRequest, stream pb.ResiliencyService_GetResiliencyStreamServer) error {
	context := stream.Context()
	for {
		select {
		case <-context.Done():
			log.Println("Client has cancelled stream")
			return nil
		default:
			resp, err := a.ResiliencyService.GetResiliencyStream(req.MaxDelaySecond, req.MinDelaySecond)
			if err != nil {
				return err
			}
			stream.Send(&pb.ResiliencyResponse{Response: resp})
		}
	}
}
func (a *GrpcAdapter) SendResiliencyStream(stream pb.ResiliencyService_SendResiliencyStreamServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			resp, err := a.ResiliencyService.SendResiliencyStream(req.MaxDelaySecond, req.MinDelaySecond)
			if err != nil {
				return err
			}
			return stream.SendAndClose(&pb.ResiliencyResponse{Response: resp})
		}
		if err != nil {
			return err
		}
	}
}
func (a *GrpcAdapter) BidirectionalResiliencyStream(stream pb.ResiliencyService_BidirectionalResiliencyStreamServer) error {
	return a.ResiliencyService.BidirectionalResiliencyStream()
}
