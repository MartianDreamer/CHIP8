package emulator

import "fmt"

func (emulator *Chip8) execute_flow_control_instruction(opcode byte, instruction [2]byte) {
	switch opcode {
	default:
		fmt.Println("executed")
	}
	fmt.Printf("execute flow control instruction %v\n", instruction)
}
