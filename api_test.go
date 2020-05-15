package grpc_test

import (
	"context"
	"fmt"
	api "github.com/phiskills/grpc-api.go"
	"github.com/phiskills/grpc-api.go/example"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"io"
	"strings"
	"testing"
)

const grpcPort = 50051

type exampleServer struct {
	example.UnimplementedServiceServer
}

func TestApi(t *testing.T) {
	fmt.Println("# Initialize test API")
	a := api.New("test-api")
	a.UsePort(grpcPort)
	if a.Status() != api.Unknown {
		t.Fatalf("New API failed: wrong status, %s should be %s", a.Status(), api.Unknown)
	}
	fmt.Println("# Register services")
	example.RegisterServiceServer(a.Server(), &exampleServer{})
	fmt.Println("# Start API")
	go a.Start()
	status := a.Change()
	if a.Status() != api.Serving || a.Status() != status {
		t.Fatalf("Start API failed: wrong status, a.status %s & a.change %s should be %s", a.Status(), status, api.Serving)
	}
	fmt.Println("# Dial API")
	addr := fmt.Sprintf("localhost:%d", grpcPort)
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Client connection failed: %v", err)
	}
	defer conn.Close()

	fmt.Println("# Initialize health client")
	healthClient := grpc_health_v1.NewHealthClient(conn)
	callHealthCheck(t, healthClient)

	fmt.Println("# Initialize main client")
	client := example.NewServiceClient(conn)

	callSimpleRequest(t, client)
	callServerStream(t, client)
	callClientStream(t, client)
	callBidirectionalStream(t, client)

	fmt.Println("# Stop API")
	a.Stop()
	status = a.Change()
	if a.Status() != api.NotServing || a.Status() != status {
		t.Fatalf("Stop API failed: wrong status, a.status %s & a.change %s should be %s", a.Status(), status, api.NotServing)
	}
}

func callHealthCheck(t *testing.T, client grpc_health_v1.HealthClient) {
	fmt.Println("## Test Health Check")
	request := grpc_health_v1.HealthCheckRequest{Service: "test-api"}
	response, err := client.Check(context.Background(), &request)
	if err != nil {
		t.Fatalf("Health Check failed: %v", err)
	}
	fmt.Printf("- Response: %v\n", response)
	expected := grpc_health_v1.HealthCheckResponse_SERVING
	if response.Status != expected {
		t.Errorf("Health Check failed: invalid response %v, expected %v", response.Status, expected)
	}
}

func callSimpleRequest(t *testing.T, client example.ServiceClient) {
	fmt.Println("## Test Simple Request")
	response, err := client.Simple(context.Background(), &example.Request{Value: "Test"})
	if err != nil {
		t.Fatalf("Simple Request failed: %v", err)
	}
	fmt.Printf("- Response: %v\n", response)
	expected := "Test"
	if response.Value != expected {
		t.Errorf("Simple Request failed: invalid response %v, expected %v", response.Value, expected)
	}
}

func (s *exampleServer) Simple(_ context.Context, request *example.Request) (*example.Response, error) {
	return &example.Response{Value: request.Value}, nil
}

func callServerStream(t *testing.T, client example.ServiceClient) {
	fmt.Println("## Test Server Stream")
	stream, err := client.ServerStream(context.Background(), &example.Request{Value: "Test"})
	if err != nil {
		t.Fatalf("Server Stream failed: %v", err)
	}
	for i := 1;; i++ {
		response, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatalf("Server Stream failed: %v", err)
		}
		fmt.Printf("- Response: %v\n", response)
		expected := fmt.Sprintf("Test %d", i)
		if response.Value != expected {
			t.Errorf("Server Stream failed: invalid response %v, expected %s", response.Value, expected)
		}
	}
}

func (s *exampleServer) ServerStream(request *example.Request, stream example.Service_ServerStreamServer) error {
	for i := 1; i < 4; i++ {
		message := fmt.Sprintf("%s %d", request.Value, i)
		if err := stream.Send(&example.Response{Value: message}); err != nil {
			return err
		}
	}
	return nil
}

func callClientStream(t *testing.T, client example.ServiceClient) {
	fmt.Println("## Test Client Stream")
	stream, err := client.ClientStream(context.Background())
	if err != nil {
		t.Fatalf("Client Stream failed: %v", err)
	}
	var messages []string
	for i := 1; i < 4; i++ {
		message := fmt.Sprintf("Test %d", i)
		err := stream.Send(&example.Request{Value: message})
		if err == io.EOF {
			break
		}
		if  err != nil {
			t.Fatalf("Client Stream failed: %v", err)
		}
		messages = append(messages, message)
	}
	response, err := stream.CloseAndRecv()
	if err != nil {
		t.Fatalf("Client Stream failed: %v", err)
	}
	fmt.Printf("- Response: %v\n", response)
	expected :=  strings.Join(messages, ", ")
	if response.Value != expected {
		t.Errorf("Client Stream failed: invalid response %v, expected %s", response.Value, expected)
	}
}

func (s *exampleServer) ClientStream(stream example.Service_ClientStreamServer) error {
	var requests []string
	for {
		request, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		requests = append(requests, request.Value)
	}
	return stream.SendAndClose(&example.Response{
		Value: strings.Join(requests, ", "),
	})
}

func callBidirectionalStream(t *testing.T, client example.ServiceClient) {
	fmt.Println("## Test Bidirectional Stream")
	stream, err := client.BidirectionalStream(context.Background())
	if err != nil {
		t.Fatalf("Bidirectional Stream failed: %v", err)
	}
	for i := 1; i < 4; i++ {
		message := fmt.Sprintf("Test %d", i)
		if err := stream.Send(&example.Request{Value: message}); err != nil {
			t.Fatalf("Bidirectional Stream failed: %v", err)
		}
		for j := 1; j < 4; j++ {
			response, err := stream.Recv()
			if err != nil {
				t.Fatalf("Bidirectional Stream failed: %v", err)
			}
			fmt.Printf("- Response: %v\n", response)
			expected :=  fmt.Sprintf("Test %d.%d", i, j)
			if response.Value != expected {
				t.Errorf("Bidirectional Stream failed: invalid response %v, expected %s", response.Value, expected)
			}
		}
	}
	stream.CloseSend()
}

func (s *exampleServer) BidirectionalStream(stream example.Service_BidirectionalStreamServer) error {
	for {
		request, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		for i := 1; i < 4; i++ {
			message := fmt.Sprintf("%s.%d", request.Value, i)
			if err = stream.Send(&example.Response{Value: message}); err != nil {
				return err
			}
		}
	}
}
