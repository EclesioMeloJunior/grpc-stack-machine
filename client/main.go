package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/EclesioMeloJunior/grpc-stack-machine/machine"
	"google.golang.org/grpc"
)

var (
	serverAddr = flag.String("serverAddr", "localhost:9111", "The grpc server address with format host:port")
)

func execute(c machine.MachineClient, set *machine.InstructionSet) {
	log.Println("gRPC: calling Execute")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	result, err := c.Execute(ctx, set)

	if err != nil {
		log.Fatalf("Problems to run Execute: %v", err)
	}

	log.Printf("Instruction result: %v\n", result.Output)
}

func main() {
	flag.Parse()

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())
	opts = append(opts, grpc.WithTimeout(time.Second*10))

	conn, err := grpc.Dial(*serverAddr, opts...)

	if err != nil {
		log.Fatalf("Could not connect to %s: %v", *serverAddr, err)
	}

	defer conn.Close()

	client := machine.NewMachineClient(conn)

	instructions := []*machine.Instruction{}
	instructions = append(instructions, &machine.Instruction{
		Operand:  5,
		Operator: "PUSH",
	})
	instructions = append(instructions, &machine.Instruction{
		Operand:  6,
		Operator: "PUSH",
	})
	instructions = append(instructions, &machine.Instruction{
		Operator: "MUL",
	})

	execute(client, &machine.InstructionSet{Instructions: instructions})
}
