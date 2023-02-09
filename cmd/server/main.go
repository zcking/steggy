package main

import (
	"fmt"
	v1 "github.com/zcking/steggy/gen/proto/go/api/v1"
	"github.com/zcking/steggy/internal"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
)

var (
	Port       = os.Getenv("PORT")
	Reflection = os.Getenv("REFLECTION")
)

func main() {
	if Port == "" {
		Port = "8080"
	}
	log.Printf("starting server on port %s...\n", Port)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", Port))
	if err != nil {
		log.Fatal(err)
	}

	impl := internal.New()

	server := grpc.NewServer()
	if Reflection == "1" {
		reflection.Register(server)
	}
	v1.RegisterSteggyServiceServer(server, impl)

	log.Printf("server started on :%s\n", Port)
	if err = server.Serve(lis); err != nil {
		log.Fatalf("server stopped: %v", err)
	}
}
