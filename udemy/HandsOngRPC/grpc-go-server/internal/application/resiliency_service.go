package application

import (
	"grpc-go-server/internal/port"
	"math/rand"
	"time"
)

type ResiliencyService struct{}

func (r *ResiliencyService) GetResiliency(req *port.ResiliencyRequest) (*port.ResiliencyResponse, error) {

	delay(req.MaxDelaySecond, req.MinDelaySecond)
	return nil, nil
}
func (r *ResiliencyService) GetResiliencyStream(req *port.ResiliencyRequest) (*port.ResiliencyResponse, error) {
	delay(req.MaxDelaySecond, req.MinDelaySecond)
	return nil, nil
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
	n := rand.Intn(int(maxDelaySecond-minDelaySecond)) + int(minDelaySecond)
	duration := time.Duration(n) * time.Millisecond
	time.Sleep(duration)
}
