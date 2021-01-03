package main_test

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/EclesioMeloJunior/grpc-stack-machine/client/mock_machine"
	"github.com/EclesioMeloJunior/grpc-stack-machine/machine"
	"github.com/golang/mock/gomock"
)

func clientExecute(t *testing.T, set *machine.InstructionSet, c machine.MachineClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	_, err := c.Execute(ctx, set)

	if err != nil {
		log.Fatalf("MachineClient Execute got an error: %v\n", err)
	}
}

func TestExecute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockMachineClient := mock_machine.NewMockMachineClient(ctrl)

	instructions := []*machine.Instruction{}
	instructions = append(instructions, &machine.Instruction{Operand: 10, Operator: "PUSH"})
	instructions = append(instructions, &machine.Instruction{Operand: 20, Operator: "PUSH"})
	instructions = append(instructions, &machine.Instruction{Operator: "MUL"})

	set := &machine.InstructionSet{Instructions: instructions}

	mockMachineClient.EXPECT().Execute(
		gomock.Any(),
		set,
	).Return(&machine.Result{Output: 200}, nil)

	clientExecute(t, set, mockMachineClient)
}
