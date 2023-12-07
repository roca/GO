package application

import "grpc-go-server/internal/port"

type ResiliencyService struct{}

func (r *ResiliencyService) GetResiliency(maxDelay, minDelay int32) (string, error) {
	return "", nil
}
func (r *ResiliencyService) GetResiliencyStream(maxDelay, minDelay int32) (string, error) {
	return "", nil
}
func (r *ResiliencyService) SendResiliencyStream([]*port.ResiliencyRequest) (string, error) {
	return "", nil
}
func (r *ResiliencyService) BidirectionalResiliencyStream(maxDelay, minDelay int32) <-chan *port.ResiliencyResponse {
	return nil
}
