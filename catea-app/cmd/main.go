package main

import (
	"context"
	"flag"
	"log"
	"time"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"github.com/MikaelHans/catea/session-management/pkg/session"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	name = flag.String("name", "dsf", "Name to greet")
)

func main() {
	flag.Parse()
	// Set up a connection to the server
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := session.NewSessionManagementClient(conn)
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	// r, err := c.GetSessionInfo(ctx, &session.SessionID{Sessionid: "asd"})
	data, err := c.SetSession(ctx, &session.SessionData{
		SessionID: &session.SessionID{Sessionid:  "asd"}, 
		Email: "mikael.hans30@gmail.com",
		Firstname: "Jonathan",
		Lastname: "Cahyo",
		Membersince: "2022-05-07 17:39:05",})
	
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", data)
}