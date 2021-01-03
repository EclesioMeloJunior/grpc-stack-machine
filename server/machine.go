package server

import (
	"context"
	"fmt"
	"log"

	"github.com/EclesioMeloJunior/grpc-stack-machine/internal/stack"
	"github.com/EclesioMeloJunior/grpc-stack-machine/machine"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type OperationType string

const (
	PUSH OperationType = "PUSH"
	POP                = "POP"
	ADD                = "ADD"
	SUB                = "SUB"
	MUL                = "MUL"
	DIV                = "DIV"
)

type MachineServer struct{}

// Execute will execute an stack of operations and returns the output
func (m *MachineServer) Execute(ctx context.Context, instructions *machine.InstructionSet) (*machine.Result, error) {
	if len(instructions.GetInstructions()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "No valid instructions received")
	}

	var s stack.Stack
	log.Println("Stack initialized")

	for _, instruction := range instructions.GetInstructions() {
		operand := instruction.GetOperand()
		operator := instruction.GetOperator()

		opType := OperationType(operator)

		switch opType {
		case PUSH:
			s.Push(float32(operand))
			log.Printf("Pushed %v to stack\n", float32(operand))
		case POP:
			if _, ok := s.Pop(); !ok {
				return nil, status.Error(codes.FailedPrecondition, "Stack empty could not pop element")
			}
			log.Printf("Popped %v from the stack\n", float32(operand))
		case ADD, SUB, MUL, DIV:
			item2, ok := s.Pop()
			item1, ok := s.Pop()

			if !ok {
				return nil, status.Error(codes.FailedPrecondition, "Stack empty could not pop element")
			}

			if opType == ADD {
				s.Push(item1 + item2)
				log.Printf("%v + %v \n", item1, item2)
			} else if opType == SUB {
				s.Push(item1 - item2)
				log.Printf("%v - %v \n", item1, item2)
			} else if opType == MUL {
				s.Push(item1 * item2)
				log.Printf("%v * %v \n", item1, item2)
			} else if opType == DIV {
				if item2 == 0 {
					return nil, status.Error(codes.Aborted, fmt.Sprintf("Could not make division of %v by %v", item1, item2))
				}

				s.Push((item1 / item2))
				log.Printf("%v / %v \n", item1, item2)
			}
		default:
			return nil, status.Error(codes.Unimplemented, fmt.Sprintf("Operation %s is not supported", operator))
		}
	}

	var ok bool
	var final float32

	if final, ok = s.Pop(); !ok {
		return nil, status.Error(codes.Aborted, "Stack empty could not pop element")
	}

	log.Println("Stack cleaned")

	return &machine.Result{Output: final}, nil
}
