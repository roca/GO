package application

type ResiliencyService struct{}

func (r *ResiliencyService) GetResiliency(maxDelay, minDelay int32) (string, error) {
	return "", nil
}
func (r *ResiliencyService) GetResiliencyStream(maxDelay, minDelay int32) (string, error) {
	return "", nil
}
func (r *ResiliencyService) SendResiliencyStream(maxDelay, minDelay int32) (string, error) {
	return "",nil
}
func (r *ResiliencyService) BidirectionalResiliencyStream() error {
	return nil
}
