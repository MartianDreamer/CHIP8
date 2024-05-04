package emulator

import "fmt"

func (emulator *Chip8) execute_flow_control_instruction(instruction [2]byte) {
	fmt.Printf("execute flow control instruction %v\n", instruction)
}
