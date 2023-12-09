package application

import (
	"grpc-go-server/internal/port"
	"math/rand"
	"time"
)

type ResiliencyService struct {
	ExitChan chan bool
}

func (r *ResiliencyService) GetResiliency(req *port.ResiliencyRequest) (*port.ResiliencyResponse, error) {

	delay(req.MaxDelaySecond, req.MinDelaySecond)

	randomIndex := rand.Intn(len(req.StatusCodes))
	pick := req.StatusCodes[randomIndex]

	return &port.ResiliencyResponse{
		StatusCode: pick,
		Response:   "OK",
		Error:      nil,
	}, nil
}
func (r *ResiliencyService) GetResiliencyStream(req *port.ResiliencyRequest) (*port.ResiliencyResponse, error) {
	return r.GetResiliency(req)
}
func (r *ResiliencyService) SendResiliencyStream(reqs []*port.ResiliencyRequest) (*port.ResiliencyResponse, error) {
	for _, req := range reqs {
		delay(req.MaxDelaySecond, req.MinDelaySecond)
	}
	return nil, nil
}
func (r *ResiliencyService) BidirectionalResiliencyStream(req *port.ResiliencyRequest) <-chan *port.ResiliencyResponse {
	delay(req.MaxDelaySecond, req.MinDelaySecond)
	return nil
}

func delay(maxDelaySecond, minDelaySecond int32) {
	n := rand.Intn(int(maxDelaySecond-minDelaySecond)+1) + int(minDelaySecond)
	duration := time.Duration(n) * time.Millisecond
	time.Sleep(duration)
}
