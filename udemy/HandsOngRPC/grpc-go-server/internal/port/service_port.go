package port

type HelloServicePort interface {
	GenerateHello(name string) string
	GenerateManyHellos(name string, count int) []string
	GenerateHelloToEveryone(names []string) string
	GenerateContinuousHello(name string, count int) <-chan string
}

type BankServicePort interface {
	FindCurrentBalance(acct string) float64
}
