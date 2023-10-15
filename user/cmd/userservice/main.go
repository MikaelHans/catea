package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/MikaelHans/catea/user/api"
	"github.com/MikaelHans/catea/user/internal/user"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 6602, "The server port")
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	s := user.Server{}
	api.RegisterUserServiceServer(server, &s)
	log.Printf("server listening at %v", lis.Addr())
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
