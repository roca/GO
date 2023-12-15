package grpc

import (
	"context"
	"fmt"
	"grpc-go-server/internal/application"
	"grpc-go-server/internal/port"
	"io"
	"log"
	pb "proto/protogen/go/resiliency"
	"strings"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func ExtractMetadata(ctx context.Context) metadata.MD {
	var extractedMetadata metadata.MD = nil
	if contextMetadata, ok := metadata.FromIncomingContext(ctx); ok {
		md := map[string]string{}
		for k, v := range contextMetadata {
			md[k] = strings.Join(v, ", ")
			log.Println("Unary GetResiliency received metadata:", k, ":", md[k])
		}
		extractedMetadata = metadata.New(md)
	}
	log.Println("Unary GetResiliency received no metadata")

	// Add some additional default metadata
	extractedMetadata["grpc-server-time"] = []string{fmt.Sprint(time.Now().Format("15:04:05.000000"))}
	extractedMetadata["grpc-server-url"] = []string{"localhost:9090"}
	extractedMetadata["grpc-response-uuid"] = []string{uuid.New().String()}

	return extractedMetadata
}

func (a *GrpcAdapter) GetResiliency(ctx context.Context, req *pb.ResiliencyRequest) (*pb.ResiliencyResponse, error) {
	status_codes_strings := []string{}
	responseMetadata := ExtractMetadata(ctx)

	for _, v := range req.StatusCodes {
		status_codes_strings = append(status_codes_strings, application.StatusCodeMap[v].String())
	}
	status_codes_string := strings.Join(status_codes_strings, ", ")
	log.Println("Unary GetResiliency with status codes:[", status_codes_string, "] and delay:", req.MinDelaySecond, "-", req.MaxDelaySecond, "seconds")

	resiliencyRequest := &port.ResiliencyRequest{
		MaxDelaySecond: req.MaxDelaySecond,
		MinDelaySecond: req.MinDelaySecond,
		StatusCodes:    req.StatusCodes,
	}

	resp, err := a.ResiliencyService.GetResiliency(resiliencyRequest)
	if err != nil {
		return nil, err
	}

	if responseMetadata != nil {
		if err := grpc.SendHeader(ctx, responseMetadata); err != nil {
			log.Println("Error while sending response metadata : ", err)
			return nil, err
		}
	}

	return &pb.ResiliencyResponse{
		Response:   resp.Response,
		StatusCode: resp.StatusCode,
	}, nil
}
func (a *GrpcAdapter) GetResiliencyStream(req *pb.ResiliencyRequest, stream pb.ResiliencyService_GetResiliencyStreamServer) error {
	context := stream.Context()
	responseMetadata := ExtractMetadata(context)
	resiliencyRequest := &port.ResiliencyRequest{
		MaxDelaySecond: req.MaxDelaySecond,
		MinDelaySecond: req.MinDelaySecond,
		StatusCodes:    req.StatusCodes,
	}
	status_codes_strings := []string{}
	for _, v := range req.StatusCodes {
		status_codes_strings = append(status_codes_strings, application.StatusCodeMap[v].String())
	}
	status_codes_string := strings.Join(status_codes_strings, ", ")
	log.Println("Server stream GetResiliencyStream with status codes:[", status_codes_string, "] and delay:", req.MinDelaySecond, "-", req.MaxDelaySecond, "seconds")
	if responseMetadata != nil {
		if err := grpc.SendHeader(context, responseMetadata); err != nil {
			log.Println("Error while sending response metadata : ", err)
			return err
		}
	}
	for {
		select {
		case <-context.Done():
			log.Println("Client has cancelled stream")
			return nil
		default:
			resp, err := a.ResiliencyService.GetResiliencyStream(resiliencyRequest)
			if err != nil {
				return err
			}
			stream.Send(&pb.ResiliencyResponse{
				Response:   resp.Response,
				StatusCode: resp.StatusCode,
			})
			log.Printf("Response sent to client %d : %s\n", resp.StatusCode, resp.Response)
		}
	}
}
func (a *GrpcAdapter) SendResiliencyStream(stream pb.ResiliencyService_SendResiliencyStreamServer) error {
	var resiliencyRequests []*port.ResiliencyRequest = nil
	var responseMetadata metadata.MD = nil
	for {
		context := stream.Context()
		req, err := stream.Recv()
		if err == io.EOF {
			if responseMetadata != nil {
				if responseMetadata != nil {
					if err := grpc.SendHeader(context, responseMetadata); err != nil {
						log.Println("Error while sending response metadata : ", err)
						return err
					}
				}
			}
			resp, err := a.ResiliencyService.SendResiliencyStream(resiliencyRequests)
			if err != nil {
				return err
			}
			return stream.SendAndClose(&pb.ResiliencyResponse{
				Response:   fmt.Sprintf("Received %d requests from client", len(resiliencyRequests)),
				StatusCode: resp.StatusCode,
			})
		}
		if err != nil {
			return err
		}
		responseMetadata = ExtractMetadata(context)
		resiliencyRequests = append(resiliencyRequests, &port.ResiliencyRequest{
			MaxDelaySecond: req.MaxDelaySecond,
			MinDelaySecond: req.MinDelaySecond,
			StatusCodes:    req.StatusCodes,
		},
		)
	}
}
func (a *GrpcAdapter) BidirectionalResiliencyStream(stream pb.ResiliencyService_BidirectionalResiliencyStreamServer) error {
	var responseMetadata metadata.MD = nil
	context := stream.Context()
	responseMetadata = ExtractMetadata(context)
	if responseMetadata != nil {
		if responseMetadata != nil {
			if err := grpc.SendHeader(context, responseMetadata); err != nil {
				log.Println("Error while sending response metadata : ", err)
				return err
			}
		}
	}
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
