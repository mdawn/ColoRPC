package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/streamingColors/colorspb"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) Color(ctx context.Context, req *colorspb.ColorRequest) (*colorspb.ColorResponse, error) {
	fmt.Printf("Color function was invoked with %v\n", req)
	adjective := req.GetColors().GetAdjective()
	result := adjective + "Blue"
	res := &colorspb.ColorResponse{
		Result: result,
	}
	return res, nil
}

func (*server) ColorEverywhere(stream colorspb.ColorService_ColorEverywhereServer) error {
	fmt.Printf("You have invoked a stream of GREEN")

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while reading color recipe: %v", err)
			return err
		}
		shade := req.GetColoring().GetAdjective()
		result := shade + " green!"

		sendErr := stream.Send(&colorspb.ColorEverywhereResponse{
			Result: result,
		})
		if sendErr != nil {
			log.Fatalf("Error while sending greens: %v", err)
			return err
		}
	}
}

func main() {
	fmt.Println("Sit tight! Your greens are coming.")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	colorspb.RegisterColorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
