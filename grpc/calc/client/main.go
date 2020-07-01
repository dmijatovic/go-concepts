package main

import (
	"context"
	"log"
	"time"

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

	time.Sleep(1 * time.Second)
	res = makeSubtractRequest(3, 3)
	log.Println("Server Subtract response...", res.Result)

	time.Sleep(1 * time.Second)
	res = makeMultiplyRequest(4, 6)
	log.Println("Server Multiply response...", res.Result)

	time.Sleep(1 * time.Second)
	res = makeDivideRequest(6, 3)
	log.Println("Server Divide response...", res.Result)

	time.Sleep(1 * time.Second)
	nums := []int32{1, 2, 4, 5, 6}
	r := makeAverageRequest(nums)
	log.Println("Server Mean response...", r)

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

// This will stram number to server one by one
func makeAverageRequest(nums []int32) float64 {
	stream, err := calcClient.Average(context.Background())
	if err != nil {
		log.Panicln(err)
	}

	for _, num := range nums {
		msg := calc.AverageRequest{
			Num: num,
		}
		//send a value to a server?!?
		stream.Send(&msg)
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Panicln(err)
	}
	return resp.GetMean()
}
