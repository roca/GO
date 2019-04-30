package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"udemy.com/gRPC/code/calculator/calculatorpb"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	fmt.Printf("Recieved Sum RPC: %v\n", req)
	firstNum := req.GetNums().GetFirstNum()
	secondNum := req.GetNums().GetSecondNum()
	total := firstNum + secondNum
	res := &calculatorpb.SumResponse{
		Total: total,
	}
	return res, nil
}

// return list of primes less than N
func PrimeFactors(n int) (pfs []int) {
	// Get the number of 2s that divide n
	for n%2 == 0 {
		pfs = append(pfs, 2)
		n = n / 2
	}

	// n must be odd at this point. so we can skip one element
	// (note i = i + 2)
	for i := 3; i*i <= n; i = i + 2 {
		// while i divides n, append i and divide n
		for n%i == 0 {
			pfs = append(pfs, i)
			n = n / i
		}
	}

	// This condition is to handle the case when n is a prime number
	// greater than 2
	if n > 2 {
		pfs = append(pfs, n)
	}

	return
}

func (*server) PrimeNumbers(req *calculatorpb.PrimeNumbersRequest, stream calculatorpb.CalculatorService_PrimeNumbersServer) error {
	fmt.Printf("PrimeNumber function was invocked with %v\n", req)
	primeNumber := req.GetNumber()
	primeFactors := PrimeFactors(int(primeNumber))
	for _, p := range primeFactors {
		res := &calculatorpb.PrimeNumbersResponse{
			Result: int64(p),
		}
		stream.Send(res)
		time.Sleep(1 * time.Second)
	}
	return nil
}

func main() {
	fmt.Println("Calculator Server")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
