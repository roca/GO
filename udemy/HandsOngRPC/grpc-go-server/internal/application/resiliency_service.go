package application

import (
	"fmt"
	"grpc-go-server/internal/port"
	"math/rand"
	"time"
)

type ResiliencyService struct {
	ExitChan chan bool
}

func (r *ResiliencyService) GetResiliency(req *port.ResiliencyRequest) (*port.ResiliencyResponse, error) {

	delay := delay(req.MaxDelaySecond, req.MinDelaySecond)

	randomIndex := rand.Intn(len(req.StatusCodes))
	pick := req.StatusCodes[randomIndex]
	str := fmt.Sprintf("The time now is %v, execution delayed for %v seconds", time.Now().Format("15:04:05.000000"), delay)

	return &port.ResiliencyResponse{
		StatusCode: pick,
		Response:   str,
		Error:      nil,
	}, nil
}
func (r *ResiliencyService) GetResiliencyStream(req *port.ResiliencyRequest) (*port.ResiliencyResponse, error) {
	return r.GetResiliency(req)
}
func (r *ResiliencyService) SendResiliencyStream(reqs []*port.ResiliencyRequest) (*port.ResiliencyResponse, error) {
	resps := make([]*port.ResiliencyResponse, 0)
	for _, req := range reqs {
		resp, err := r.GetResiliency(req)
		if err != nil {
			return nil, err
		}
		resps = append(resps, resp)
	}
	return resps[len(resps)-1], nil
}
func (r *ResiliencyService) BidirectionalResiliencyStream(req *port.ResiliencyRequest) <-chan *port.ResiliencyResponse {
	ch := make(chan *port.ResiliencyResponse)
	go func() {
		resp, err := r.GetResiliency(req)
		if err != nil {
			ch <- &port.ResiliencyResponse{
				Response:   "",
				StatusCode: 0,
				Error:      err,
			}
			close(ch)
			return
		}
		ch <- resp
		close(ch)
	}()
	return ch
}

func delay(maxDelaySecond, minDelaySecond int32) int {
	n := rand.Intn(int(maxDelaySecond-minDelaySecond)+1) + int(minDelaySecond)
	duration := time.Duration(n) * time.Millisecond
	time.Sleep(duration)
	return n
}
