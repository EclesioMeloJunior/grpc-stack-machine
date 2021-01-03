package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/EclesioMeloJunior/grpc-stack-machine/machine"
	"github.com/EclesioMeloJunior/grpc-stack-machine/server"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 9111, "Port where grcp server will listen TCP conn.")
)

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))

	if err != nil {
		log.Fatalf("Could not start server: %v", err)
	}

	grpcServer := grpc.NewServer()
	machine.RegisterMachineServer(grpcServer, &server.MachineServer{})

	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("Could not start grpc server: %v", err)
	}
}
