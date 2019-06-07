package main

import (
	"grpcMultiplication/grpcnew/server"
)

func main() {

	server.Start("50051")

	//time.Sleep(2 * time.Second)

	//client.Start("50051")
}
