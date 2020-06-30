package main

import (
	"context"
	"log"

	calc "../pb"
	"google.golang.org/grpc"
)

var host string = "localhost:50051"

var calcClient calc.CalcServiceClient

func main() {
	println("Client starting...", host)

	//connect to gRPC server
	cnn, err := gRPCCnn()
	if err != nil {
		log.Println("Closing client...")
		return
	}
	defer cnn.Close()

	// get reference to Client interface
	// note this is loaded as global variable
	calcClient = createClient(cnn)

	// make some requests
	res := makeAddRequest(1, 3)
	log.Println("Server Add response...", res.Result)

	res = makeSubtractRequest(3, 3)
	log.Println("Server Subtract response...", res.Result)

	res = makeMultiplyRequest(4, 6)
	log.Println("Server Multiply response...", res.Result)

	res = makeDivideRequest(6, 3)
	log.Println("Server Divide response...", res.Result)

}

func gRPCCnn() (*grpc.ClientConn, error) {
	cnn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		log.Println("Client connection failed...", host)
	}
	return cnn, err
}

func createClient(cnn *grpc.ClientConn) calc.CalcServiceClient {
	c := calc.NewCalcServiceClient(cnn)
	return c
}

func makeAddRequest(num1 int32, num2 int32) *calc.CalcResponse {
	req := &calc.CalcRequest{
		Num_1:     num1,
		Num_2:     num2,
		Operation: "+",
	}
	res, err := calcClient.Add(context.Background(), req)
	if err != nil {
		log.Panicln(err)
	}
	return res
}

func makeSubtractRequest(num1 int32, num2 int32) *calc.CalcResponse {
	req := &calc.CalcRequest{
		Num_1:     num1,
		Num_2:     num2,
		Operation: "-",
	}
	res, err := calcClient.Subtract(context.Background(), req)
	if err != nil {
		log.Panicln(err)
	}
	return res
}

func makeMultiplyRequest(num1 int32, num2 int32) *calc.CalcResponse {
	req := &calc.CalcRequest{
		Num_1:     num1,
		Num_2:     num2,
		Operation: "*",
	}
	res, err := calcClient.Multiply(context.Background(), req)
	if err != nil {
		log.Panicln(err)
	}
	return res
}

func makeDivideRequest(num1 int32, num2 int32) *calc.CalcResponse {
	req := &calc.CalcRequest{
		Num_1:     num1,
		Num_2:     num2,
		Operation: "/",
	}
	res, err := calcClient.Divide(context.Background(), req)
	if err != nil {
		log.Panicln(err)
	}
	return res
}
