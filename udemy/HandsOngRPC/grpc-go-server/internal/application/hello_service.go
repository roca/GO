package application

type HelloService struct{}

func (a *HelloService) GenerateHello(name string) string {
	return "Hello " + name
}