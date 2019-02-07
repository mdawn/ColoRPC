package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/mdawn/ColoRPC/colorspb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {

	tls := false
	opts := grpc.WithInsecure()
	if tls {
		certFile := "ssl/ca.crt" // Certificate Authority Trust certificate
		creds, sslErr := credentials.NewClientTLSFromFile(certFile, "")
		if sslErr != nil {
			log.Fatalf("Error while loading CA trust certificate: %v", sslErr)
			return
		}
		opts = grpc.WithTransportCredentials(creds)
	}

	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := colorspb.NewColorServiceClient(cc)

	doStreaming(c)
	// doUnary(c)
}

func doUnary(c colorspb.ColorServiceClient) {
	req := &colorspb.ColorRequest{
		Colors: &colorspb.Coloring{
			Adjective: "Piss ",
			BaseColor: "Yellow",
		},
	}

	res, err := c.Color(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling color rpc: %v", err)
	}
	log.Printf("Response from color: %v", res.Result)

}

func doStreaming(c colorspb.ColorServiceClient) {
	fmt.Println("Starting streaming RPC...")

	// invoke the client
	stream, err := c.ColorEverywhere(context.Background())
	if err != nil {
		log.Fatalf("Error while creating stream: %v", err)
		return
	}

	requests := []*colorspb.ColorEverywhereRequest{
		&colorspb.ColorEverywhereRequest{
			Coloring: &colorspb.Coloring{
				Adjective: "Grass",
			},
		},
		&colorspb.ColorEverywhereRequest{
			Coloring: &colorspb.Coloring{
				Adjective: "Pearly",
			},
		},
		&colorspb.ColorEverywhereRequest{
			Coloring: &colorspb.Coloring{
				Adjective: "Kelly",
			},
		},
		&colorspb.ColorEverywhereRequest{
			Coloring: &colorspb.Coloring{
				Adjective: "Olive",
			},
		},
		&colorspb.ColorEverywhereRequest{
			Coloring: &colorspb.Coloring{
				Adjective: "Puke",
			},
		},
		&colorspb.ColorEverywhereRequest{
			Coloring: &colorspb.Coloring{
				Adjective: "Forest",
			},
		},
		&colorspb.ColorEverywhereRequest{
			Coloring: &colorspb.Coloring{
				Adjective: "Moss",
			},
		},
		&colorspb.ColorEverywhereRequest{
			Coloring: &colorspb.Coloring{
				Adjective: "Flourescent",
			},
		},
		&colorspb.ColorEverywhereRequest{
			Coloring: &colorspb.Coloring{
				Adjective: "Sea",
			},
		},
		&colorspb.ColorEverywhereRequest{
			Coloring: &colorspb.Coloring{
				Adjective: "Neon",
			},
		},
		&colorspb.ColorEverywhereRequest{
			Coloring: &colorspb.Coloring{
				Adjective: "Pastel",
			},
		},
		&colorspb.ColorEverywhereRequest{
			Coloring: &colorspb.Coloring{
				Adjective: "Mint",
			},
		},
		&colorspb.ColorEverywhereRequest{
			Coloring: &colorspb.Coloring{
				Adjective: "Two-tone",
			},
		},
		&colorspb.ColorEverywhereRequest{
			Coloring: &colorspb.Coloring{
				Adjective: "Bottle",
			},
		},
		&colorspb.ColorEverywhereRequest{
			Coloring: &colorspb.Coloring{
				Adjective: "Translucent",
			},
		},
		&colorspb.ColorEverywhereRequest{
			Coloring: &colorspb.Coloring{
				Adjective: "Mantis",
			},
		},
		&colorspb.ColorEverywhereRequest{
			Coloring: &colorspb.Coloring{
				Adjective: "Psychedelic",
			},
		},
		&colorspb.ColorEverywhereRequest{
			Coloring: &colorspb.Coloring{
				Adjective: "Tea",
			},
		},
		&colorspb.ColorEverywhereRequest{
			Coloring: &colorspb.Coloring{
				Adjective: "Shamrock",
			},
		},
		&colorspb.ColorEverywhereRequest{
			Coloring: &colorspb.Coloring{
				Adjective: "Money",
			},
		},
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range requests {
			stream.Send(req)
			time.Sleep(1000 * time.Millisecond)
		}
		stream.CloseSend()
	}()
	// we receive a bunch of messages from the client (go routine)
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while receiving: %v", err)
				break
			}
			fmt.Printf("%v\n", res.GetResult())
		}
		close(waitc)
	}()

	// block until everything is done
	<-waitc
}
