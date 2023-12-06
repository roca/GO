package application

type ResiliencyService struct{}

func (r *ResiliencyService) GetResiliency() (interface{}, error) {
	return nil, nil
}
func (r *ResiliencyService) GetResiliencyStream() error {
	return nil
}
func (r *ResiliencyService) SendResiliencyStream() error {
	return nil
}
func (r *ResiliencyService) BidirectionalResiliencyStream() error {
	return nil
}
