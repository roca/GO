package grpc

import (
	"context"
	"grpc-go-server/internal/port"
	"io"
	"log"
	pb "proto/protogen/go/resiliency"
	"time"
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
	var resiliencyRequests []*port.ResiliencyRequest = nil
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			resp, err := a.ResiliencyService.SendResiliencyStream(resiliencyRequests)
			if err != nil {
				return err
			}
			return stream.SendAndClose(&pb.ResiliencyResponse{Response: resp})
		}
		if err != nil {
			return err
		}
		resiliencyRequests = append(resiliencyRequests, &port.ResiliencyRequest{
			MaxDelaySecond: req.MaxDelaySecond,
			MinDelaySecond: req.MinDelaySecond,
		},
		)
	}
}
func (a *GrpcAdapter) BidirectionalResiliencyStream(stream pb.ResiliencyService_BidirectionalResiliencyStreamServer) error {
	context := stream.Context()
	for {
		select {
		case <-context.Done():
			log.Println("Client has cancelled stream")
			return nil
		default:
			req, err := stream.Recv()
			if err == io.EOF {
				return nil
			}
			if err != nil {
				return err
			}

			for resp := range a.ResiliencyService.BidirectionalResiliencyStream(req.MaxDelaySecond, req.MinDelaySecond) {
				if resp.Error != nil {
					return resp.Error
				}
				_ = stream.Send(&pb.ResiliencyResponse{Response: resp.Response})
				// if err != nil {
				// 	return err
				// }
				time.Sleep(500 * time.Millisecond)
			}
		}
	}
}
