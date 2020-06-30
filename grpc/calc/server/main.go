package main

import (
	"context"
	"log"
	"net"

	calc "../pb"
	"google.golang.org/grpc"
)

var host string = "0.0.0.0:50051"

// gRPC calc Service
type calcSvc struct{}

// gRPC Service implements Calc interface as defined in protobuff requrements
// see pb/calc.pb.go
func (*calcSvc) Add(ctx context.Context, req *calc.CalcRequest) (*calc.CalcResponse, error) {

	// log.Println("Request received on Calc channel")

	n1 := req.Num_1
	n2 := req.Num_2

	log.Println("Add channel...num1...", n1)
	log.Println("Add channel...num2...", n2)

	res := &calc.CalcResponse{
		Result: int64(n1 + n2),
	}

	return res, nil
}

func (*calcSvc) Subtract(ctx context.Context, req *calc.CalcRequest) (*calc.CalcResponse, error) {

	// log.Println("Request received on Send channel")

	n1 := req.Num_1
	n2 := req.Num_2

	log.Println("Subtract channel...num1...", n1)
	log.Println("Subtract channel...num2...", n2)

	res := &calc.CalcResponse{
		Result: int64(n1 - n2),
	}

	return res, nil
}

func (*calcSvc) Multiply(ctx context.Context, req *calc.CalcRequest) (*calc.CalcResponse, error) {

	// log.Println("Request received on Send channel")

	n1 := req.Num_1
	n2 := req.Num_2

	log.Println("Multiply channel...num1...", n1)
	log.Println("Multiply channel...num2...", n2)

	res := &calc.CalcResponse{
		Result: int64(n1 * n2),
	}

	return res, nil
}

func (*calcSvc) Divide(ctx context.Context, req *calc.CalcRequest) (*calc.CalcResponse, error) {

	// log.Println("Request received on Send channel")

	n1 := req.Num_1
	n2 := req.Num_2

	log.Println("Divide channel...num1...", n1)
	log.Println("Divide channel...num2...", n2)

	res := &calc.CalcResponse{
		Result: int64(n1 / n2),
	}

	return res, nil
}

func logError(e error) {
	log.Println("Server error: ", e)
}

func main() {
	println("Server starting...", host)

	//start tcp session
	tcp, err := net.Listen("tcp", host)
	if err != nil {
		logError(err)
	}
	//create new gRPC server
	s := grpc.NewServer()
	//pass gRPC server to protobuffer
	calc.RegisterCalcServiceServer(s, &calcSvc{})
	//start server on tcp net
	err = s.Serve(tcp)
	if err != nil {
		log.Fatalln("Failed to start gRPC Calc server")
	}
}
