package server_test

import (
	"context"
	"testing"

	"github.com/EclesioMeloJunior/grpc-stack-machine/machine"
	"github.com/EclesioMeloJunior/grpc-stack-machine/server"
)

func TestMachineServerExecute(t *testing.T) {
	var instruction_set_1 []*machine.Instruction
	instruction_set_1 = append(instruction_set_1, &machine.Instruction{
		Operator: "PUSH",
		Operand:  4,
	})
	instruction_set_1 = append(instruction_set_1, &machine.Instruction{
		Operator: "PUSH",
		Operand:  5,
	})
	instruction_set_1 = append(instruction_set_1, &machine.Instruction{
		Operator: "MUL",
	})

	var instruction_set_2 []*machine.Instruction
	instruction_set_2 = append(instruction_set_2, &machine.Instruction{
		Operator: "PUSH",
		Operand:  4,
	})
	instruction_set_2 = append(instruction_set_2, &machine.Instruction{
		Operator: "PUSH",
		Operand:  5,
	})
	instruction_set_2 = append(instruction_set_2, &machine.Instruction{
		Operator: "ADD",
	})

	tests := []struct {
		set      *machine.InstructionSet
		expected float32
	}{
		{
			set:      &machine.InstructionSet{Instructions: instruction_set_1},
			expected: 20,
		},
		{
			set:      &machine.InstructionSet{Instructions: instruction_set_2},
			expected: 9,
		},
	}

	s := &server.MachineServer{}

	for _, test := range tests {
		resp, err := s.Execute(context.Background(), test.set)

		if err != nil {
			t.Errorf("MachineServer.Execute() got an error: %v", err)
		}

		if resp.Output != test.expected {
			t.Errorf("Got: %v\nExpected: %v", resp.Output, test.expected)
		}
	}
}
