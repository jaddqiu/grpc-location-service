package main

import (
	"context"
	"log"
	"os"
	"time"

	locationpb "github.com/jaddqiu/grpc-location-service/pb"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:8080"
	defaultName = "jadd"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("cannot connect: %v", err)
	}

	defer conn.Close()

	c := locationpb.NewLocationClient(conn)

	name := defaultName

	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	switch len(os.Args) {
	case 1, 2:
		r, err := c.Query(ctx, &locationpb.QueryRequest{Name: name})

		if err != nil {
			log.Fatalf("cannot query: %v", err)
		}
		log.Printf("location for %s: %s", name, r.Location)
	default:
		location := os.Args[2]
		r, err := c.Post(ctx, &locationpb.PostRequest{Name: name, Location: location})

		if err != nil {
			log.Fatalf("cannot post: %v", err)
		}
		log.Printf("location post for %s message: %v", name, r.Message)

	}

}
