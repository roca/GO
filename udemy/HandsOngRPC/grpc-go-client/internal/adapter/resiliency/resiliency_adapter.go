package resiliency

import (
	"context"
	"grpc-go-client/internal/port"
	"io"
	"log"
	pb "proto/protogen/go/resiliency"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
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

var dummyMetadata = map[string]string{
	"client_metadata": "some client metadata",
}

func NewResiliencyAdapter(conn *grpc.ClientConn) (*ResiliencyAdapter, error) {
	client := pb.NewResiliencyServiceClient(conn)

	return &ResiliencyAdapter{
		resiliencyClient: client,
	}, nil
}

func (a *ResiliencyAdapter) GetResiliency(ctx context.Context, resiliencyRequest *ResiliencyRequest) (*ResiliencyResponse, error) {

	requestMetadata := metadata.New(dummyMetadata)
	ctx = metadata.NewOutgoingContext(ctx, requestMetadata)

	var responseMetadata metadata.MD
	req := &pb.ResiliencyRequest{
		MaxDelaySecond: resiliencyRequest.MaxDelaySecond,
		MinDelaySecond: resiliencyRequest.MinDelaySecond,
		StatusCodes:    resiliencyRequest.StatusCodes,
	}
	resp, err := a.resiliencyClient.GetResiliency(ctx, req, grpc.Header(&responseMetadata))
	if err != nil {
		return nil, err
	}

	if responseMetadata.Len() == 0 {
		log.Println("Unary GetResiliency received no metadata")
	} else {
		for k, v := range responseMetadata {
			log.Println("Unary GetResiliency received metadata:", k, ":", v)
		}
	}
	return &ResiliencyResponse{
		Response:   resp.Response,
		StatusCode: resp.StatusCode,
	}, nil
}

func (a *ResiliencyAdapter) GetResiliencyStream(ctx context.Context, resiliencyRequest *ResiliencyRequest) error {
	req := &pb.ResiliencyRequest{
		MaxDelaySecond: resiliencyRequest.MaxDelaySecond,
		MinDelaySecond: resiliencyRequest.MinDelaySecond,
		StatusCodes:    resiliencyRequest.StatusCodes,
	}
	requestMetadata := metadata.New(dummyMetadata)
	ctx = metadata.NewOutgoingContext(ctx, requestMetadata)


	stream, err := a.resiliencyClient.GetResiliencyStream(ctx, req)
	if responseMetadata, err := stream.Header(); err == nil {
		if responseMetadata.Len() == 0 {
			log.Println("Unary GetResiliency received no metadata")
		} else {
			for k, v := range responseMetadata {
				log.Println("Unary GetResiliency received metadata:", k, ":", v)
			}
		}
	}
	if err != nil {
		log.Println("Error on GetResiliencyStream:", err)
		return err
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println("Error on FetchExchangeRates:", err)
			return err
		}
		log.Println(res.Response, res.StatusCode)
	}

	return nil
}

func (a *ResiliencyAdapter) SendResiliencyStream(ctx context.Context, reqs []*ResiliencyRequest) (*ResiliencyResponse, error) {
	requestMetadata := metadata.New(dummyMetadata)
	ctx = metadata.NewOutgoingContext(ctx, requestMetadata)
	stream, err := a.resiliencyClient.SendResiliencyStream(ctx)
	if err != nil {
		log.Println("Error getting stream on SummarizeTransactions:", err)
		return nil, err
	}

	for _, req := range reqs {
		err := stream.Send(&pb.ResiliencyRequest{
			MaxDelaySecond: req.MaxDelaySecond,
			MinDelaySecond: req.MinDelaySecond,
			StatusCodes:    req.StatusCodes,
		})
		if err != nil {
			log.Println("Error sending name on stream SummarizeTransactions:", err)
			return nil, err
		}
		// time.Sleep(500 * time.Millisecond)
	}

	resp, err := stream.CloseAndRecv()
	if responseMetadata, err := stream.Header(); err == nil {
		if responseMetadata.Len() == 0 {
			log.Println("Unary GetResiliency received no metadata")
		} else {
			for k, v := range responseMetadata {
				log.Println("Unary GetResiliency received metadata:", k, ":", v)
			}
		}
	}
	if err != nil {
		log.Println("Error closing and receiving response on stream SayHelloToEveryone:", err)
		return nil, err
	}

	return &ResiliencyResponse{
		Response:   resp.Response,
		StatusCode: resp.StatusCode,
	}, nil
}

func (a *ResiliencyAdapter) BidirectionalResiliencyStream(ctx context.Context, reqs []*ResiliencyRequest) error {
	stream, err := a.resiliencyClient.BidirectionalResiliencyStream(ctx)
	if err != nil {
		log.Println("Error getting stream on BidirectionalResiliencyStream:", err)
		return err
	}

	resiliencyChan := make(chan struct{})

	go func() {

		for _, req := range reqs {
			err := stream.Send(&pb.ResiliencyRequest{
				MaxDelaySecond: req.MaxDelaySecond,
				MinDelaySecond: req.MinDelaySecond,
				StatusCodes:    req.StatusCodes,
			})
			if err != nil {
				log.Println("Error sending name on stream BidirectionalResiliencyStream:", err)
			}
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				s := status.Convert(err)
				log.Fatalln("Can not invoke SummarizeTransactions on the BankAdapter:", "\nCode:", s.Code(), "\nMessage:", s.Message(), "\nDetails:", s.Details())
			}
			log.Println(res)
		}
		close(resiliencyChan)
	}()

	<-resiliencyChan

	return nil
}
