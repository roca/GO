package grpc

import (
	"context"
	"grpc-go-server/internal/port"
	"io"
	"log"
	pb "proto/protogen/go/resiliency"
)

func (a *GrpcAdapter) GetResiliency(ctx context.Context, req *pb.ResiliencyRequest) (*pb.ResiliencyResponse, error) {
	log.Println("GetResiliency")

	resiliencyRequest := &port.ResiliencyRequest{
		MaxDelaySecond: req.MaxDelaySecond,
		MinDelaySecond: req.MinDelaySecond,
		StatusCodes:    req.StatusCodes,
	}

	resp, err := a.ResiliencyService.GetResiliency(resiliencyRequest)

	return &pb.ResiliencyResponse{
		Response:   resp.Response,
		StatusCode: resp.StatusCode,
	}, err
}
func (a *GrpcAdapter) GetResiliencyStream(req *pb.ResiliencyRequest, stream pb.ResiliencyService_GetResiliencyStreamServer) error {
	context := stream.Context()
	resiliencyRequest := &port.ResiliencyRequest{
		MaxDelaySecond: req.MaxDelaySecond,
		MinDelaySecond: req.MinDelaySecond,
		StatusCodes:    req.StatusCodes,
	}
	for {
		select {
		case <-context.Done():
			log.Println("Client has cancelled stream")
			return nil
		default:
			_, err := a.ResiliencyService.GetResiliencyStream(resiliencyRequest)
			if err != nil {
				return err
			}
			stream.Send(&pb.ResiliencyResponse{})
		}
	}
}
func (a *GrpcAdapter) SendResiliencyStream(stream pb.ResiliencyService_SendResiliencyStreamServer) error {
	var resiliencyRequests []*port.ResiliencyRequest = nil
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			_, err := a.ResiliencyService.SendResiliencyStream(resiliencyRequests)
			if err != nil {
				return err
			}
			return stream.SendAndClose(&pb.ResiliencyResponse{})
		}
		if err != nil {
			return err
		}
		resiliencyRequests = append(resiliencyRequests, &port.ResiliencyRequest{
			MaxDelaySecond: req.MaxDelaySecond,
			MinDelaySecond: req.MinDelaySecond,
			StatusCodes:    req.StatusCodes,
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
			resiliencyRequest := &port.ResiliencyRequest{
				MaxDelaySecond: req.MaxDelaySecond,
				MinDelaySecond: req.MinDelaySecond,
				StatusCodes:    req.StatusCodes,
			}

			for resp := range a.ResiliencyService.BidirectionalResiliencyStream(resiliencyRequest) {
				if resp.Error != nil {
					return resp.Error
				}
				_ = stream.Send(&pb.ResiliencyResponse{Response: resp.Response})
				// if err != nil {
				// 	return err
				// }

			}
		}
	}
}
