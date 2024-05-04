package emulator

import "fmt"

func (emulator *Chip8) execute_alu_instruction(instruction [2]byte) {
	fmt.Printf("execute alu instruction %v\n", instruction)
}
