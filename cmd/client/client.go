package main

// for testing the grpc server
// go run cmd/client/client.go
import (
	"context"
	"log"
	"time"

	"atlassian.carcgl.com/bitbucket/ls/lms/pkg/api/proto/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

// local server address
const (
	lms_server_address = "localhost:9094"
)

// main entry point
func main() {
	// create a connection to the lms gRPC server
	conn, err := grpc.Dial(lms_server_address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect %v", err)
	} else {
		log.Printf("connnect to gRPC server at %s", lms_server_address)
	}

	// close the connection when done
	defer conn.Close()

	c := proto.NewLMSRecordServiceClient(conn)

	// https://go.dev/blog/context
	// In Go servers, each incoming request is handled in its own goroutine.
	// Request handlers often start additional goroutines to access backends such as databases
	// and RPC services. The set of goroutines working on a request typically needs access to
	// request-specific values such as the identity of the end user, authorization tokens,
	// and the request’s deadline. When a request is canceled or times out, all the goroutines
	// working on that request should exit quickly so the system can reclaim any resources they are using.
	// At Google, we developed a context package that makes it easy to pass request-scoped values,
	// cancelation signals, and deadlines across API boundaries to all the goroutines involved
	// in handling a request. The package is publicly available as context.

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	var new_operators = make(map[string]bool)

	new_operators["Qi"] = true
	new_operators["Stig"] = true
	new_operators["Espen"] = true

	for name, valid := range new_operators {
		r, err := c.CreateOperator(ctx, &proto.CreateOperatorRequest{Name: name, Valid: valid})
		if err != nil {
			log.Fatalf("could not create operator %v", err)
		}
		log.Printf(`Operator ID: %s`, r.GetOperatorId())
	}

	r, err := c.ListOperators(ctx, &emptypb.Empty{})
	if err != nil {
		log.Fatalf("could not get operator list %v", err)
	}
	log.Print("\n Operators List: \n")
	log.Printf(" %v\n", r.GetOperators())
}
