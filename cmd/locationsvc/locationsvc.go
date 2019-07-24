package main

import (
	"context"
	"log"
	"net"

	locationpb "github.com/jaddqiu/grpc-location-service/pb"
	"google.golang.org/grpc"
)

type server struct {
	Locations map[string]string
}

func (s *server) Query(ctx context.Context, in *locationpb.QueryRequest) (*locationpb.QueryResponse, error) {
	log.Printf("Query request received, name: %v\n", in.Name)
	location := s.Locations[in.Name]
	if location == "" {
		location = "Not Found"
	}
	return &locationpb.QueryResponse{Location: location}, nil
}

func (s *server) Post(ctx context.Context, in *locationpb.PostRequest) (*locationpb.PostResponse, error) {
	log.Printf("Post request received, name: %v, location: %v", in.Name, in.Location)
	_, ok := s.Locations[in.Name]
	var message string
	if ok {
		message = "Update location successfully"
	} else {
		message = "Add location successfully"
	}
	s.Locations[in.Name] = in.Location

	return &locationpb.PostResponse{Result: 0, Message: message}, nil

}

func main() {
	var Location = make(map[string]string)
	Location["jadd"] = "pingzhou"

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	locationpb.RegisterLocationServer(s, &server{Locations: Location})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
