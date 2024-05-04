package emulator

import "fmt"

func (emulator *Chip8) execute_display_instruction(instruction [2]byte) {
	fmt.Printf("execute display instruction %v\n", instruction)
}