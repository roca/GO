package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math"
	"net"
	"time"

	"google.golang.org/grpc/codes"

	"udemy.com/gRPC/code/calculator/calculatorpb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
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

func (*server) ComputeAverage(stream calculatorpb.CalculatorService_ComputeAverageServer) error {
	fmt.Println("ComputeAverage function was invocked with a streaming request")
	sum := float64(0)
	count := float64(0)

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// we have finished reading the client stream
			return stream.SendAndClose(&calculatorpb.ComputeAverageResponse{
				Result: sum / count,
			})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}

		number := req.GetNumber()
		count++
		sum += float64(number)
	}
}

func (*server) FindMaximum(stream calculatorpb.CalculatorService_FindMaximumServer) error {
	fmt.Println("FindMaximum function was invocked with a streaming request")
	maximum := int64(0)

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
			return err
		}

		number := req.GetNumber()
		if number > maximum {
			maximum = number
			if err := stream.Send(&calculatorpb.FindMaximumResponse{Result: maximum}); err != nil {
				log.Fatalf("Error while sending data to client: %v", err)
				return err
			}
		}

	}
}

func (*server) SquareRoot(ctx context.Context, req *calculatorpb.SquareRootRequest) (*calculatorpb.SquareRootResponse, error) {
	fmt.Printf("Recieved SquareRoot RPC: %v\n", req)
	number := req.GetNumber()

	if number < 0 {
		return nil, status.Error(
			codes.InvalidArgument,
			fmt.Sprintf("Received a negative number: %v", number),
		)
	}

	res := &calculatorpb.SquareRootResponse{
		NumberRoot: math.Sqrt(float64(number)),
	}
	return res, nil
}

func main() {
	fmt.Println("Calculator Server")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	// Register reflection service on gRPC server.
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
