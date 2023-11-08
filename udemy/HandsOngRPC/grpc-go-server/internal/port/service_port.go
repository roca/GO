package port

type HelloServicePort interface {
	GenerateHello(name string) string
}