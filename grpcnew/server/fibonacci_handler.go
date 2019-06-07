package server

import (
	"context"
	"training/grpcnew/fibonaccipb/proto"
)

type Fibonacci_Handler struct{}

func (ch *Fibonacci_Handler) FibonacciSeries(ctx context.Context, request *proto.Request) (*proto.FibonacciResponse, error) {
	response := &proto.FibonacciResponse{}

	response.Result = request.GetNumber() * 56

	return response, nil
}
